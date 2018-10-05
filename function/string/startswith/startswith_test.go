package startswith

import (
	"fmt"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression"
	"github.com/stretchr/testify/assert"
)

var s = &StartsWith{}

func TestStaticFunc_Starts_with(t *testing.T) {
	final1 := s.Eval("TIBCO Web Integrator", "TIBCO")
	fmt.Println(final1)
	assert.Equal(t, true, final1)

	final2 := s.Eval("网路 Integrator", "网路")
	fmt.Println(final2)
	assert.Equal(t, true, final2)

	final3 := s.Eval("TIBCO 网路 Integrator", "网路")
	fmt.Println(final3)
	assert.Equal(t, false, final3)
}

func TestExpression(t *testing.T) {
	fun, err := expression.ParseExpression(`string.startsWith("seafood,name", "sea")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}
