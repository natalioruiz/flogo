package lastindex

import (
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("lastIndex-function")

type LastIndex struct {
}

func init() {
	function.Registry(&LastIndex{})
}

func (s *LastIndex) GetName() string {
	return "lastIndex"
}

func (s *LastIndex) GetCategory() string {
	return "string"
}

func (s *LastIndex) Eval(str, sep string) int {
	log.Debugf("Returns the index of the last instance of \"%s\" in \"%s\", or -1 if not present", sep, str)
	var l int
	l = strings.LastIndex(str, sep)
	log.Debugf("Done calculating the last index: %n", l)
	return l
}
