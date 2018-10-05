package Int64

import (
	"time"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"git.tibco.com/git/product/ipaas/wi-contrib.git/function/datetime"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("now-function")

type Now struct {
}

func init() {
	function.Registry(&Now{})
}

func (s *Now) GetName() string {
	return "now"
}

func (s *Now) GetCategory() string {
	return "datetime"
}

func (s *Now) Eval() string {
	log.Debugf("Returns the current datetime with timezone")
	var currentTime time.Time
	location, err := time.LoadLocation(datetime.GetLocation())
	if err != nil {
		log.Errorf("Load location %s error %s", datetime.GetLocation(), err.Error())
		location = time.UTC
		currentTime = time.Now().UTC()
	} else {
		currentTime = time.Now().In(location)
	}
	return currentTime.Format(datetime.GetDatetimeFormat())
}
