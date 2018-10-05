package equals

import (
	"fmt"
	"testing"

	_ "git.tibco.com/git/product/ipaas/wi-contrib.git/function/boolean/false"
	_ "git.tibco.com/git/product/ipaas/wi-contrib.git/function/boolean/true"
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression"
	"github.com/stretchr/testify/assert"
)

var s = &Equals{}

func TestStaticFunc_Starts_with(t *testing.T) {
	final1 := s.Eval("TIBCO Web Integrator", "TIBCO")
	fmt.Println(final1)
	assert.Equal(t, false, final1)

	final2 := s.Eval("123T", "123t")
	fmt.Println(final2)
	assert.Equal(t, false, final2)

}

func TestExpression(t *testing.T) {
	fun, err := expression.ParseExpression(`string.equals("TIBCO NAME", "TIBCO NAME")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}

func TestExpressionIgnoreCase(t *testing.T) {
	fun, err := expression.ParseExpression(`string.equals("TIBCO name", "TIBCO NAME")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}
