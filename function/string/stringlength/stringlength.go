package stringlength

import (
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("length-function")

type StringLength struct {
}

func init() {
	function.Registry(&StringLength{})
}

func (s *StringLength) GetName() string {
	return "length"
}

func (s *StringLength) GetCategory() string {
	return "string"
}

func (s *StringLength) Eval(str string) int {
	log.Debugf("Return the length of a string \"%s\"", str)
	var l int
	//l = len([]rune(str))
	l = len(str)
	log.Debugf("Done calculating the length %d", l)
	return l
}
