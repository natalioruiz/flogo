package substringafter

import (
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("substringAfter-function")

type Substringafter struct {
}

func init() {
	function.Registry(&Substringafter{})
}

func (s *Substringafter) GetName() string {
	return "substringAfter"
}

func (s *Substringafter) GetCategory() string {
	return "string"
}

func (s *Substringafter) Eval(str, afterStr string) string {
	log.Debugf("Start Substringafter function with parameters string %s after string %s", str, afterStr)
	if strings.Index(str, afterStr) >= 0 {
		return str[strings.Index(str, afterStr)+len(afterStr):]
	} else {
		return str
	}
}
