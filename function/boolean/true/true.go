package true

import (
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("true-function")

type True struct {
}

func init() {
	function.Registry(&True{})
}

func (s *True) GetName() string {
	return "true"
}

func (s *True) GetCategory() string {
	return "boolean"
}

func (s *True) Eval() bool {
	log.Debugf("Always returns true")
	return true
}
