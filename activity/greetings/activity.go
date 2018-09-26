package greetings

import (
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

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
func (a *MyActivity) Eval(context activity.Context) (done bool, err error)  {

	firstName := context.GetInput("firstName").(string)
	lastName := context.GetInput("lastName").(string)

	message := fmt.Sprintf("Welcome back %s %s!!!", firstName, lastName)

	context.SetOutput("message", message)
	return true, nil
}
