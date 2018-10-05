package substringbefore

import (
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("substringBefore-function")

type Substringbefore struct {
}

func init() {
	function.Registry(&Substringbefore{})
}

func (s *Substringbefore) GetName() string {
	return "substringBefore"
}

func (s *Substringbefore) GetCategory() string {
	return "string"
}

func (s *Substringbefore) Eval(str, beforeStr string) string {
	log.Debugf("Start substringbefore function with parameters string %s before string %s", str, beforeStr)
	if strings.Index(str, beforeStr) >= 0 {
		return str[:strings.Index(str, beforeStr)]
	} else {
		return str
	}
}
