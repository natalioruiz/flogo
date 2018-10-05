package contains

import (
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("contains-function")

type Contains struct {
}

func init() {
	function.Registry(&Contains{})
}

func (s *Contains) GetName() string {
	return "contains"
}

func (s *Contains) GetCategory() string {
	return "string"
}

func (s *Contains) Eval(str, substr string) bool {
	log.Debugf("Reports whether \"%s\" is within \"%s\"", substr, str)

	return strings.Contains(str, substr)
}
