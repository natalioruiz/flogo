package currentdatetime

import (
	"fmt"
	"os"
	"testing"

	"git.tibco.com/git/product/ipaas/wi-contrib.git/function/datetime"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression"
	"github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestNow_Eval(t *testing.T) {
	n := CurrentDatetime{}
	datetime := n.Eval()
	assert.NotNil(t, datetime)
	logrus.Info(datetime)
	fmt.Println(datetime)
}

func TestDatetime_CDT(t *testing.T) {
	os.Setenv(datetime.WI_DATETIME_LOCATION, "US/Central")
	n := CurrentDatetime{}
	date := n.Eval()
	assert.NotNil(t, date)
	logrus.Info(date)
}

func TestExpression(t *testing.T) {
	fun, err := expression.ParseExpression(`datetime.currentDatetime()`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}
