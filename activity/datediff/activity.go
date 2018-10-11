// Package addtodate adds a specified number of units to a date.
package addtodate

import (
	"time"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// Constants used by the code to represent the input and outputs of the JSON structure
const (
	number = "number"
	units  = "units"
	date   = "date"
	format = "format"
	result = "result"
)

// log is the default package logger
var log = logger.GetLogger("activity-addtodate")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	// Get the inputs
	ivNumber := context.GetInput(number).(int)
	ivUnits := context.GetInput(units).(string)
	ivDate := context.GetInput(date).(string)
	ivFormat := context.GetInput(format).(string)

	date := time.Now()

	if ivFormat == "" {
		ivFormat = "2006-01-02"
	}

	if ivDate != "" {
		date, _ = time.Parse(ivFormat, ivDate)
	}
	log.Infof("Date before: %s", date.Format(ivFormat))
	switch ivUnits {
	case "days":
		date = date.AddDate(0, 0, 1*ivNumber)
	case "months":
		date = date.AddDate(0, 1*ivNumber, 0)
	case "years":
		date = date.AddDate(1*ivNumber, 0, 0)
	case "hours":
		date = date.Add(time.Duration(ivNumber) * time.Hour)
	case "minutes":
		date = date.Add(time.Duration(ivNumber) * time.Minute)
	case "seconds":
		date = date.Add(time.Duration(ivNumber) * time.Second)
	}

	// Set the output value in the context
	context.SetOutput(result, date.Format(ivFormat))

	log.Infof("Date after: %s", date.Format(ivFormat))
	return true, nil
}
