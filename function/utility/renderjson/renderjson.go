package renderjson

import (
	"bytes"
	"encoding/json"
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("renderJSON-function")

type RenderJson struct {
}

func init() {
	function.Registry(&RenderJson{})
}

func (s *RenderJson) GetName() string {
	return "renderJSON"
}

func (s *RenderJson) GetCategory() string {
	return "utility"
}

func (s *RenderJson) Eval(object interface{}, printPretty bool) (string, error) {
	log.Debugf("Start render-json function with parameters %s", object)
	var result string
	if printPretty {
		b, err := json.MarshalIndent(object, "", "    ")
		if err != nil {
			return "", err
		}
		result = string(b)
	} else {
		buffer := new(bytes.Buffer)
		json_bytes, err := json.Marshal(object)
		if err != nil {
			return "", err
		}
		err = json.Compact(buffer, json_bytes)
		if err != nil {
			return "", err
		}
		result = buffer.String()
	}

	log.Debugf("Done render-json function with result %s", result)
	return result, nil
}
