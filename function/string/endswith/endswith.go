package endswith

import (
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("endsWith-function")

type EndsWith struct {
}

func init() {
	function.Registry(&EndsWith{})
}

func (s *EndsWith) GetName() string {
	return "endsWith"
}

func (s *EndsWith) GetCategory() string {
	return "string"
}
func (s *EndsWith) Eval(str, substr string) bool {
	log.Debugf("Reports whether \"%s\" ends with \"%s\"", str, substr)

	return strings.HasSuffix(str, substr)
}
