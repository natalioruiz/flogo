package Int64

import (
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("Int64-function")

type Int struct {
}

func init() {
	function.Registry(&Int{})
}

func (s *Int) GetName() string {
	return "int64"
}

func (s *Int) GetCategory() string {
	return "number"
}

func (s *Int) Eval(in interface{}) (int, error) {
	log.Debugf("Start Int64 function with parameters %s", in)
	return data.CoerceToInteger(in)
}
