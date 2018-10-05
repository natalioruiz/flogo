package formatdatetime

import (
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"git.tibco.com/git/product/ipaas/wi-contrib.git/function/datetime"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/tkuchiki/parsetime"
)

var log = logger.GetLogger("formatDatetime-function")

type FormatDatetime struct {
}

func init() {
	function.Registry(&FormatDatetime{})
}

func (s *FormatDatetime) GetName() string {
	return "formatDatetime"
}

func (s *FormatDatetime) GetCategory() string {
	return "datetime"
}

func (s *FormatDatetime) Eval(date string, format string) (string, error) {
	log.Debugf("Format datetime %s to format %s", date, format)
	p, err := parsetime.NewParseTime(datetime.GetLocation())
	if err != nil {
		log.Errorf("New datetime parser %s error %s", date, err.Error())
		return date, err
	}
	t, err := p.Parse(date)
	if err != nil {
		log.Errorf("Parsing datetime %s error %s", date, err.Error())
		return date, err
	}
	return t.Format(format), nil
}
