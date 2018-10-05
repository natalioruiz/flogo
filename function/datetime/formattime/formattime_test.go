package formattime

import (
	"fmt"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression"
	"git.tibco.com/git/product/ipaas/wi-contrib.git/function/datetime"
	_ "git.tibco.com/git/product/ipaas/wi-contrib.git/function/string/timeformat"

	"github.com/stretchr/testify/assert"
)

func TestFormatTime_Eval(t *testing.T) {
	n := FormatTime{}
	date, err := n.Eval("10:11:05.00000 ", datetime.GetTimeFormat())
	assert.Nil(t, err)
	assert.NotNil(t, date)
	assert.Equal(t, "10:11:05+00:00", date)
	fmt.Println(date)
}

func TestExpression(t *testing.T) {
	fun, err := expression.ParseExpression(`datetime.formatTime("10:11:05.00000", string.timeFormat())`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}
