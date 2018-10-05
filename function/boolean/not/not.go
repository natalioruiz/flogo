package not

import (
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("not-function")

type Not struct {
}

func init() {
	function.Registry(&Not{})
}

func (s *Not) GetName() string {
	return "not"
}

func (s *Not) GetCategory() string {
	return "boolean"
}

func (s *Not) Eval(value bool) bool {
	log.Debugf("Not %v", value)
	return !value
}
