package index

import (
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("index-function")

type Index struct {
}

func init() {
	function.Registry(&Index{})
}

func (s *Index) GetName() string {
	return "index"
}

func (s *Index) GetCategory() string {
	return "string"
}

func (s *Index) Eval(str, sep string) int {
	log.Debugf("Returns the index of the first instance of \"%s\" in \"%s\", or -1 if not present", sep, str)
	var l int
	l = strings.Index(str, sep)
	log.Debugf("Done calculating the index: %n", l)
	return l
}
