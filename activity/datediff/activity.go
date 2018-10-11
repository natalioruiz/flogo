// Package datediff adds a specified number of units to a date.
package datediff

import (
	"time"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// Constants used by the code to represent the input and outputs of the JSON structure
const (
	date1  = "date1"
	date2  = "date2"
	unit   = "unit"
	result = "result"
)

// log is the default package logger
var log = logger.GetLogger("activity-datediff")

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
	ivDate1 := context.GetInput("date1").(string)
	ivFormat1 := context.GetInput("format1").(string)
	ivDate2 := context.GetInput("date2").(string)
	ivFormat2 := context.GetInput("format2").(string)
	ivUnits := context.GetInput("units").(string)

	if ivFormat1 == "" {
		ivFormat1 = "2006-01-02 15:04:05.000"
	}

	if ivFormat2 == "" {
		ivFormat1 = "2006-01-02 15:04:05.000"
	}

	date1, _ := time.Parse(ivFormat1, ivDate1)
	date2 := time.Now()
	if ivDate2 != "" {
		date2, _ = time.Parse(ivFormat2, ivDate2)
	}

	log.Infof("Date 1: %s", date1.Format(ivFormat1))
	log.Infof("Date 2: %s", date2.Format(ivFormat2))

	d := date2.Sub(date1)
	var output float64
	switch ivUnits {
	case "hours":
		output = d.Hours()
	case "minutes":
		output = d.Minutes()
	case "seconds":
		output = d.Seconds()
	case "millis":
		output = d.Seconds() * 1e3
	}

	// Set the output value in the context
	context.SetOutput(result, int(output))

	log.Infof("Difference: %s %s", output, ivUnits)
	return true, nil
}
