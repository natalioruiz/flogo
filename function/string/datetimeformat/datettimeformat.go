package datetimeformat

import (
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"git.tibco.com/git/product/ipaas/wi-contrib.git/function/datetime"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("datetimeFormat-function")

type DatetimeFormat struct {
}

func init() {
	function.Registry(&DatetimeFormat{})
}

func (s *DatetimeFormat) GetName() string {
	return "datetimeFormat"
}

func (s *DatetimeFormat) GetCategory() string {
	return "string"
}

func (s *DatetimeFormat) Eval() string {
	return datetime.GetDatetimeFormat()
}
