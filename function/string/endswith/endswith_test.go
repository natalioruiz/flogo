package endswith

import (
	"fmt"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression"
	"github.com/stretchr/testify/assert"
)

var s = &EndsWith{}

func TestStaticFunc_Starts_with(t *testing.T) {
	final1 := s.Eval("TIBCO Web Integrator", "Integrator")
	fmt.Println(final1)
	assert.Equal(t, true, final1)

	final2 := s.Eval("TIBCO Web 集成器", "集成器")
	fmt.Println(final2)
	assert.Equal(t, true, final2)

	final3 := s.Eval("TIBCO 网路 Integrator", "网路")
	fmt.Println(final3)
	assert.Equal(t, false, final3)
}

func TestExpression(t *testing.T) {
	fun, err := expression.ParseExpression(`string.endsWith("TIBCO NAME", "NAME")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}
