package contains

import (
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression"
	"github.com/stretchr/testify/assert"
	"testing"
)

var s = &Contains{}

func TestStaticFunc_Contains(t *testing.T) {
	final1 := s.Eval("TIBCO Web Integrator", "Web")
	fmt.Println(final1)
	assert.Equal(t, true, final1)

	final2 := s.Eval("TIBCO 网路 Integrator", "网路")
	fmt.Println(final2)
	assert.Equal(t, true, final2)
}

func TestExpression(t *testing.T) {
	fun, err := expression.ParseExpression(`string.contains("TIBCO Web Integrator","Web")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.Equal(t, true, v)
	fmt.Println(v)
}
