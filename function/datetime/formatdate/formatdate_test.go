package formatdate

import (
	"fmt"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression"
	"git.tibco.com/git/product/ipaas/wi-contrib.git/function/datetime"
	_ "git.tibco.com/git/product/ipaas/wi-contrib.git/function/datetime/currentdate"
	_ "git.tibco.com/git/product/ipaas/wi-contrib.git/function/string/dateformat"

	"github.com/stretchr/testify/assert"
)

func TestFormatDate_Eval(t *testing.T) {
	n := FormatDate{}

	date, err := n.Eval("02/08/2017", "20060102")
	assert.Nil(t, err)
	assert.NotNil(t, date)
	assert.Equal(t, "20170208", date)
	fmt.Println(date)
}

func TestFormatDate_Eval2(t *testing.T) {
	n := FormatDate{}
	date, err := n.Eval("02/08/2017", "2006-02-01")
	assert.Nil(t, err)
	assert.NotNil(t, date)
	assert.Equal(t, "2017-08-02", date)
	fmt.Println(date)
}

func TestFormatDate_Eval3(t *testing.T) {
	n := FormatDate{}
	date, err := n.Eval("02/08/2017", "01-02-2006")
	assert.Nil(t, err)
	assert.NotNil(t, date)
	assert.Equal(t, "02-08-2017", date)
	fmt.Println(date)
}

func TestFormatDate_Eval4(t *testing.T) {
	n := FormatDate{}
	date, err := n.Eval("02-08-2017", datetime.GetDateFormat())
	assert.Nil(t, err)
	assert.NotNil(t, date)
	assert.Equal(t, "2017-02-08+00:00", date)
	fmt.Println(date)
}

func TestExpression(t *testing.T) {
	fun, err := expression.ParseExpression(`datetime.formatDate("02/08/2017", string.dateFormat())`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}

func TestExpression3(t *testing.T) {
	fun, err := expression.ParseExpression(`datetime.formatDate(datetime.currentDate(), string.dateFormat())`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}
