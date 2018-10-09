package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/TIBCOSoftware/flogo-contrib/trigger/rest/cors"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/julienschmidt/httprouter"
	"github.com/xeipuuv/gojsonschema"
)

const (
	REST_CORS_PREFIX = "REST_TRIGGER"
)

// log is the default package logger
var log = logger.GetLogger("trigger-flogo-rest")

var validMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}

// RestTrigger REST trigger struct
type RestTrigger struct {
	metadata *trigger.Metadata
	//runner   action.Runner
	server *Server
	config *trigger.Config
	//handlers []*handler.Handler
}

//NewFactory create a new Trigger factory
func NewFactory(md *trigger.Metadata) trigger.Factory {
	return &RestFactory{metadata: md}
}

// RestFactory REST Trigger factory
type RestFactory struct {
	metadata *trigger.Metadata
}

//New Creates a new trigger instance for a given id
func (t *RestFactory) New(config *trigger.Config) trigger.Trigger {
	return &RestTrigger{metadata: t.metadata, config: config}
}

// Metadata implements trigger.Trigger.Metadata
func (t *RestTrigger) Metadata() *trigger.Metadata {
	return t.metadata
}

func (t *RestTrigger) Initialize(ctx trigger.InitContext) error {
	router := httprouter.New()

	if t.config.Settings == nil {
		return fmt.Errorf("no Settings found for trigger '%s'", t.config.Id)
	}

	if _, ok := t.config.Settings["port"]; !ok {
		return fmt.Errorf("no Port found for trigger '%s' in settings", t.config.Id)
	}

	addr := ":" + t.config.GetSetting("port")

	pathMap := make(map[string]string)

	// Init handlers
	for _, handler := range ctx.GetHandlers() {

		err := validateHandler(t.config.Id, handler)
		if err != nil {
			return err
		}

		method := strings.ToUpper(handler.GetStringSetting("method"))
		path := handler.GetStringSetting("path")
		schema := handler.GetStringSetting("schema")
		requiredHeaders := handler.GetStringSetting("requiredHeaders")

		log.Debugf("Registering handler [%s: %s]", method, path)

		if _, ok := pathMap[path]; !ok {
			pathMap[path] = path
			router.OPTIONS(path, handleCorsPreflight) // for CORS
		}

		//router.OPTIONS(path, handleCorsPreflight) // for CORS
		router.Handle(method, path, newActionHandler(t, handler, schema, requiredHeaders))
	}

	log.Debugf("Configured on port %s", t.config.Settings["port"])
	t.server = NewServer(addr, router)

	return nil
}

func (t *RestTrigger) Start() error {
	return t.server.Start()
}

// Stop implements util.Managed.Stop
func (t *RestTrigger) Stop() error {
	return t.server.Stop()
}

// Handles the cors preflight request
func handleCorsPreflight(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	log.Infof("Received [OPTIONS] request to CorsPreFlight: %+v", r)

	c := cors.New(REST_CORS_PREFIX, log)
	c.HandlePreflight(w, r)
}

// IDResponse id response object
type IDResponse struct {
	ID string `json:"id"`
}

func newActionHandler(rt *RestTrigger, handler *trigger.Handler, schema string, requiredHeaders string) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		log.Infof("Received request for id '%s'", rt.config.Id)

		c := cors.New(REST_CORS_PREFIX, log)
		c.WriteCorsActualRequestHeaders(w)

		pathParams := make(map[string]string)
		for _, param := range ps {
			pathParams[param.Key] = param.Value
		}

		queryValues := r.URL.Query()
		queryParams := make(map[string]string, len(queryValues))
		header := make(map[string]string, len(r.Header))

		for key, value := range r.Header {
			header[key] = strings.Join(value, ",")
		}

		for key, value := range queryValues {
			queryParams[key] = strings.Join(value, ",")
		}

		triggerData := map[string]interface{}{
			"params":      pathParams,
			"pathParams":  pathParams,
			"queryParams": queryParams,
			"header":      header,
		}

		// Check the HTTP Header Content-Type
		contentType := r.Header.Get("Content-Type")
		switch contentType {
		case "application/x-www-form-urlencoded":
			buf := new(bytes.Buffer)
			buf.ReadFrom(r.Body)
			s := buf.String()
			m, err := url.ParseQuery(s)
			content := make(map[string]interface{}, 0)
			if err != nil {
				log.Errorf("Error while parsing query string: %s", err.Error())
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			for key, val := range m {
				if len(val) == 1 {
					content[key] = val[0]
				} else {
					content[key] = val[0]
				}
			}
			triggerData["content"] = content
		default:
			var content interface{}
			err := json.NewDecoder(r.Body).Decode(&content)
			if err != nil {
				switch {
				case err == io.EOF:
					// empty body
					//todo should handler say if content is expected?
				case err != nil:
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
			}
			triggerData["content"] = content
		}

		validHeaders := true
		errorMessage := ""
		for _, rh := range strings.Split(requiredHeaders, ",") {
			if _, ok := header[strings.TrimSpace(rh)]; !ok {
				validHeaders = false
				errorMessage = fmt.Sprintf("Missing required header: '%s'", rh)
				break
			}
		}
		results := make(map[string]*data.Attribute)
		var err error
		if validHeaders {
			validRequest := true
			if schema != "" {

				jsonData, _ := ioutil.ReadAll(r.Body) //<--- here!
				requestDump, _ := httputil.DumpRequest(r, true)
				doc := gojsonschema.NewStringLoader(string(jsonData))
				schema := gojsonschema.NewStringLoader(schema)

				log.Infof("Schema: %s", schema)
				log.Infof("Body: %s", string(requestDump))
				result, err := gojsonschema.Validate(schema, doc)
				if err != nil {
					validRequest = false
				} else {
					if !result.Valid() {
						validRequest = false
						var errorsList []string
						for _, e := range result.Errors() {
							errorsList = append(errorsList, fmt.Sprintf("%s", e))
						}
						err = errors.New(strings.Join(errorsList, "\n"))
					}
				}
			}
			if validRequest {
				results, err = handler.Handle(context.Background(), triggerData)
			}
		} else {
			err = errors.New(errorMessage)
		}

		var replyData interface{}
		var replyCode int
		var replyHeaders map[string]interface{}

		if len(results) != 0 {
			dataAttr, ok := results["data"]
			if ok {
				replyData = dataAttr.Value()
			}
			codeAttr, ok := results["code"]
			if ok {
				replyCode, _ = data.CoerceToInteger(codeAttr.Value())
			}
			headersAttr, ok := results["headers"]
			if ok {
				replyHeaders = headersAttr.Value().(map[string]interface{})
			}
		}

		if err != nil {
			log.Debugf("REST Trigger Error: %s", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if replyHeaders != nil {
			for k, v := range replyHeaders {
				w.Header().Set(k, v.(string))
			}
		}

		if replyData != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			if replyCode == 0 {
				replyCode = 200
			}
			w.WriteHeader(replyCode)
			if err := json.NewEncoder(w).Encode(replyData); err != nil {
				log.Error(err)
			}
			return
		}

		if replyCode > 0 {
			w.WriteHeader(replyCode)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}

////////////////////////////////////////////////////////////////////////////////////////
// Utils

func validateHandler(triggerId string, handler *trigger.Handler) error {

	method := handler.GetStringSetting("method")
	if method == "" {
		return fmt.Errorf("no method specified in handler for trigger '%s'", triggerId)
	}

	if !stringInList(strings.ToUpper(method), validMethods) {
		return fmt.Errorf("invalid method '%s' specified in handler for trigger '%s'", method, triggerId)
	}

	//validate path

	return nil
}

func stringInList(str string, list []string) bool {
	for _, value := range list {
		if value == str {
			return true
		}
	}
	return false
}
