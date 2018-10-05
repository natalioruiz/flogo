package formatdate

import (
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"git.tibco.com/git/product/ipaas/wi-contrib.git/function/datetime"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/tkuchiki/parsetime"
)

var log = logger.GetLogger("formatDate-function")

type FormatDate struct {
}

func init() {
	function.Registry(&FormatDate{})
}

func (s *FormatDate) GetName() string {
	return "formatDate"
}

func (s *FormatDate) GetCategory() string {
	return "datetime"
}

func (s *FormatDate) Eval(date string, format string) (string, error) {
	log.Debugf("Format date %s to format %s", date, format)
	p, err := parsetime.NewParseTime(datetime.GetLocation())
	if err != nil {
		log.Errorf("New date parser %s error %s", date, err.Error())
		return date, err
	}
	t, err := p.US(date)
	if err != nil {
		//Try toresolved it with just parse
		t, err = p.Parse(date)
		if err != nil {
			log.Errorf("Parsing date %s error %s", date, err.Error())
			return date, err
		}
	}
	return t.Format(format), nil
}
