package substring

import (
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("substring-function")

type Substring struct {
}

func init() {
	function.Registry(&Substring{})
}

func (s *Substring) GetName() string {
	return "substring"
}

func (s *Substring) GetCategory() string {
	return "string"
}

func (s *Substring) Eval(str string, index, length int) string {
	log.Debugf("Start substring function with parameters string %s index %d and lenght %d", str, index, length)
	return str[index : index+length]
}
