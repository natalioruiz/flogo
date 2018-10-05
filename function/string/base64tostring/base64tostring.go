package base64tostring

import (
	"encoding/base64"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("base64ToString-function")

type Base64ToString struct {
}

func init() {
	function.Registry(&Base64ToString{})
}

func (s *Base64ToString) GetName() string {
	return "base64ToString"
}

func (s *Base64ToString) GetCategory() string {
	return "string"
}

func (s *Base64ToString) Eval(base64str string) (string, error) {
	log.Debugf("Decode base64 to string \"%s\"", base64str)

	decoded, err := base64.StdEncoding.DecodeString(base64str)
	if err != nil {
		log.Debugf("decode error:", err)
		return "", err
	}

	return string(decoded), err
}
