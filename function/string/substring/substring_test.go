package substring

import (
	"fmt"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression"
	_ "git.tibco.com/git/product/ipaas/wi-contrib.git/function/datetime/currentdate"
	"github.com/stretchr/testify/assert"
)

func TestStaticFunc_Substring(t *testing.T) {
	str := "TIBCO software Inc"
	s := &Substring{}
	subStr := s.Eval(str, 0, 5)
	fmt.Println(subStr)
	assert.Equal(t, subStr, "TIBCO")
}

func TestSample(t *testing.T) {
	sub := &Substring{}
	result := sub.Eval("12345", 2, 3)
	assert.Equal(t, "345", result)
}

func TestExpression(t *testing.T) {
	fun, err := expression.ParseExpression(`string.substring("1999/04/01",2,3)`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}

func TestExpression2(t *testing.T) {
	fun, err := expression.ParseExpression(`string.substring("TIBCO",2,3)`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}

func TestExpression3(t *testing.T) {
	fun, err := expression.ParseExpression(`string.substring(datetime.currentDate(), 0, 10)`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}
