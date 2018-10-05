package len

import (
	"git.tibco.com/git/product/ipaas/wi-contrib.git/engine/conversion"
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("len-function")

type Len struct {
}

func init() {
	function.Registry(&Len{})
}

func (s *Len) GetName() string {
	return "len"
}

func (s *Len) GetCategory() string {
	return "number"
}

func (s *Len) Eval(str interface{}) int {
	log.Debugf("Start len function with parameters %s", str)
	switch t := str.(type) {
	case string:
		return len(t)
	case []interface{}:
		return len(t)
	default:
		s, err := conversion.ConvertToString(str)
		if err != nil {
			log.Errorf("Convert %+v to string error %s", str, err.Error())
		} else {
			return len(s)
		}
	}
	return 0
}
