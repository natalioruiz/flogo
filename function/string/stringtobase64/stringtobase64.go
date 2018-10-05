package stringtobase64

import (
	"encoding/base64"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("stringToBase64-function")

type StringToBase64 struct {
}

func init() {
	function.Registry(&StringToBase64{})
}

func (s *StringToBase64) GetName() string {
	return "stringToBase64"
}

func (s *StringToBase64) GetCategory() string {
	return "string"
}

func (s *StringToBase64) Eval(str string) string {
	log.Debugf("Returns base64 encoding of a given string \"%s\"", str)

	encoded := base64.StdEncoding.EncodeToString([]byte(str))

	return encoded
}
