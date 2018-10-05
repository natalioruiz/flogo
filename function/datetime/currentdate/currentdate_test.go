package currentdate

import (
	"testing"

	"os"

	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression"
	"git.tibco.com/git/product/ipaas/wi-contrib.git/function/datetime"
	"github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestNow_Eval(t *testing.T) {
	n := CurrentDate{}
	date := n.Eval()
	assert.NotNil(t, date)
	logrus.Info(date)
}

func TestNow_CDT(t *testing.T) {
	os.Setenv(datetime.WI_DATETIME_LOCATION, "US/Central")
	n := CurrentDate{}
	date := n.Eval()
	assert.NotNil(t, date)
	logrus.Info(date)
}

func TestExpression(t *testing.T) {
	fun, err := expression.ParseExpression(`datetime.currentDate()`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}
