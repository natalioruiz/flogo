package dateformat

import (
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"git.tibco.com/git/product/ipaas/wi-contrib.git/function/datetime"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("dateFormat-function")

type DateFormat struct {
}

func init() {
	function.Registry(&DateFormat{})
}

func (s *DateFormat) GetName() string {
	return "dateFormat"
}

func (s *DateFormat) GetCategory() string {
	return "string"
}

func (s *DateFormat) Eval() string {
	return datetime.GetDateFormat()
}
