package substringafter

import (
	"fmt"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression"
	"github.com/stretchr/testify/assert"
)

var s = &Substringafter{}

func TestStaticFunc_SubstringAfter(t *testing.T) {
	str := "TIBCO software Inc"
	final := s.Eval(str, " ")
	fmt.Println(final)
	assert.Equal(t, final, "software Inc")
}

func TestSample(t *testing.T) {
	final := s.Eval("1999/04/01", "/")
	fmt.Println(final)
	assert.Equal(t, final, "04/01")
}

func TestExpression(t *testing.T) {
	fun, err := expression.ParseExpression(`string.substringAfter("1999/04/01","/")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}

func TestExpressionIPAS4962(t *testing.T) {
	fun, err := expression.ParseExpression(`string.substringAfter("This is a sample server Petstore server ","sample")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}

func TestExpressionIPAS5021(t *testing.T) {
	fun, err := expression.ParseExpression(`string.substringAfter("This is a sample server Petstore server ", "This")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}
