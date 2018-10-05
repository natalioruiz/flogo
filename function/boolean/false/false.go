package false

import (
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("false-function")

type False struct {
}

func init() {
	function.Registry(&False{})
}

func (s *False) GetName() string {
	return "false"
}

func (s *False) GetCategory() string {
	return "boolean"
}

func (s *False) Eval() bool {
	log.Debugf("Always returns false")
	return false
}
