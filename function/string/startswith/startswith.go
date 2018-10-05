package startswith

import (
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("startsWith-function")

type StartsWith struct {
}

func init() {
	function.Registry(&StartsWith{})
}

func (s *StartsWith) GetName() string {
	return "startsWith"
}

func (s *StartsWith) GetCategory() string {
	return "string"
}

func (s *StartsWith) Eval(str, substr string) bool {
	log.Infof("Reports whether \"%s\" begins with \"%s\"", str, substr)

	return strings.HasPrefix(str, substr)
}
