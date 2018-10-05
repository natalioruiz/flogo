package split

import (
	"fmt"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression"
	"github.com/stretchr/testify/assert"
)

var s = &Split{}

func TestStaticFunc_Split(t *testing.T) {
	final1 := s.Eval("TIBCO Web Integrator", " ")
	fmt.Println(final1)
	assert.Equal(t, "TIBCO", final1[0])
	assert.Equal(t, "Web", final1[1])
	assert.Equal(t, "Integrator", final1[2])

	final2 := s.Eval("TIBCO。网路。Integrator", "。")
	fmt.Println(final2)
	assert.Equal(t, "TIBCO", final2[0])
	assert.Equal(t, "网路", final2[1])
	assert.Equal(t, "Integrator", final2[2])

}

func TestExpression(t *testing.T) {
	fun, err := expression.ParseExpression(`string.split("seafood,name", ",")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}

func TestExpression2(t *testing.T) {
	fun, err := expression.ParseExpression(`string.split("seafood namefood", " ")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}
