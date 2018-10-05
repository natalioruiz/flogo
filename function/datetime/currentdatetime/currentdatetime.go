package currentdatetime

import (
	"time"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"git.tibco.com/git/product/ipaas/wi-contrib.git/function/datetime"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("currentDatetime-function")

type CurrentDatetime struct {
}

func init() {
	function.Registry(&CurrentDatetime{})
}

func (s *CurrentDatetime) GetName() string {
	return "currentDatetime"
}

func (s *CurrentDatetime) GetCategory() string {
	return "datetime"
}

func (s *CurrentDatetime) Eval() string {
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
