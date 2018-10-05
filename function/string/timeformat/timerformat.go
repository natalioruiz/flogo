package timeformat

import (
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"git.tibco.com/git/product/ipaas/wi-contrib.git/function/datetime"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("timeFormat-function")

type TimeFormat struct {
}

func init() {
	function.Registry(&TimeFormat{})
}

func (s *TimeFormat) GetName() string {
	return "timeFormat"
}

func (s *TimeFormat) GetCategory() string {
	return "string"
}

func (s *TimeFormat) Eval() string {
	return datetime.GetTimeFormat()
}
