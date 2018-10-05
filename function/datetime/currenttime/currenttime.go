package currenttime

import (
	"time"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression/function"
	"git.tibco.com/git/product/ipaas/wi-contrib.git/function/datetime"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("currentTime-function")

type CurrentTime struct {
}

func init() {
	function.Registry(&CurrentTime{})
}

func (s *CurrentTime) GetName() string {
	return "currentTime"
}

func (s *CurrentTime) GetCategory() string {
	return "datetime"
}

func (s *CurrentTime) Eval() string {
	log.Debugf("Returns the current time with timezone")
	var currentTime time.Time
	location, err := time.LoadLocation(datetime.GetLocation())
	if err != nil {
		log.Errorf("Load location %s error %s", datetime.GetLocation(), err.Error())
		location = time.UTC
		currentTime = time.Now().UTC()
	} else {
		currentTime = time.Now().In(location)
	}
	return currentTime.Format(datetime.GetTimeFormat())
}
