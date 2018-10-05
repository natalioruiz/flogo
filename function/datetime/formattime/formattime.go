package formattime

import (
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"git.tibco.com/git/product/ipaas/wi-contrib.git/function/datetime"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/tkuchiki/parsetime"
)

var log = logger.GetLogger("formatTime-function")

type FormatTime struct {
}

func init() {
	function.Registry(&FormatTime{})
}

func (s *FormatTime) GetName() string {
	return "formatTime"
}

func (s *FormatTime) GetCategory() string {
	return "datetime"
}

func (s *FormatTime) Eval(date string, format string) (string, error) {
	log.Debugf("Format time %s to format %s", date, format)
	p, err := parsetime.NewParseTime(datetime.GetLocation())
	if err != nil {
		log.Errorf("New time parser %s error %s", date, err.Error())
		return date, err
	}
	t, err := p.Parse(date)
	if err != nil {
		log.Errorf("Parsing time %s error %s", date, err.Error())
		return date, err
	}
	return t.Format(format), nil
}
