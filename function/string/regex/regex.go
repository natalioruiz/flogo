package regex

import (
	"regexp"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("regex-function")

type Regex struct {
}

func init() {
	function.Registry(&Regex{})
}

func (s *Regex) GetName() string {
	return "regex"
}

func (s *Regex) GetCategory() string {
	return "string"
}

func (s *Regex) Eval(pattern, str string) bool {
	log.Debugf("Checks whether a textual regular expression \"%s\" matches a string \"%s\"", pattern, str)

	matched, err := regexp.MatchString(pattern, str)

	if err != nil {
		log.Debugf("Error occurred during regular expression matching: %s", err)
	}

	return matched
}
