package split

import (
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("split-function")

type Split struct {
}

func init() {
	function.Registry(&Split{})
}

func (s *Split) GetName() string {
	return "split"
}

func (s *Split) GetCategory() string {
	return "string"
}

func (s *Split) Eval(str, sep string) []string {
	log.Debugf("Slices \"%s\" into all substrings separated by \"%s\" and returns a slice of the substrings between those separators", str, sep)
	return strings.Split(str, sep)
}
