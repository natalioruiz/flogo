package trim

import (
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("trim-function")

type Trim struct {
}

func init() {
	function.Registry(&Trim{})
}

func (s *Trim) GetName() string {
	return "trim"
}

func (s *Trim) GetCategory() string {
	return "string"
}

func (s *Trim) Eval(str string) string {
	log.Debugf("Trim a string \"%s\" with all leading and trailing white space removed", str)
	return strings.TrimSpace(str)
}
