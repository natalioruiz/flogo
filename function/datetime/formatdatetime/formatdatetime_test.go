package formatdatetime

import (
	"fmt"
	"testing"

	"git.tibco.com/git/product/ipaas/wi-contrib.git/function/string/datetimeformat"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression"
	"github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestFormatDatetime_Eval(t *testing.T) {
	n := &FormatDatetime{}

	date, err := n.Eval("2017-04-12T22:15:09", "2006-01-02 15:04:05")
	assert.Nil(t, err)
	assert.NotNil(t, date)
	assert.Equal(t, "2017-04-12 22:15:09", date)
	fmt.Println(date)
}

func TestFormatDatetime_Eval2(t *testing.T) {
	n := &FormatDatetime{}
	date, err := n.Eval("2017-04-10T22:17:32.000+0000", "2006-01-02 15:04:05")
	assert.Nil(t, err)
	assert.NotNil(t, date)
	assert.Equal(t, "2017-04-10 22:17:32", date)
	fmt.Println(date)
}

func TestFormatDatetime_Default(t *testing.T) {
	n := &FormatDatetime{}
	datetimeForamt := datetimeformat.DatetimeFormat{}

	date, err := n.Eval("2017-04-10T22:17:32.000+0000", datetimeForamt.Eval())
	assert.Nil(t, err)
	assert.NotNil(t, date)
	logrus.Info(date)
	assert.Equal(t, "2017-04-10T22:17:32+00:00", date)
	fmt.Println(date)
}

func TestExpression(t *testing.T) {
	fun, err := expression.ParseExpression(`datetime.formatDatetime("2017-04-10T22:17:32.000+0000", string.datetimeFormat())`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}
