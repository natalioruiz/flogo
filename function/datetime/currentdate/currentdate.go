package currentdate

import (
	"time"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"git.tibco.com/git/product/ipaas/wi-contrib.git/function/datetime"

	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("currentDate-function")

type CurrentDate struct {
}

func init() {
	function.Registry(&CurrentDate{})
}

func (s *CurrentDate) GetName() string {
	return "currentDate"
}

func (s *CurrentDate) GetCategory() string {
	return "datetime"
}

func (s *CurrentDate) Eval() string {
	log.Debugf("Returns the current date with timezone")
	var currentTime time.Time
	location, err := time.LoadLocation(datetime.GetLocation())
	if err != nil {
		log.Errorf("Load location %s error %s", datetime.GetLocation(), err.Error())
		location = time.UTC
		currentTime = time.Now().UTC()
	} else {
		currentTime = time.Now().In(location)
	}
	return currentTime.Format(datetime.GetDateFormat())
}
