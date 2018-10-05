package lastindex

import (
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression"
	"github.com/stretchr/testify/assert"
	"testing"
)

var s = &LastIndex{}

func TestStaticFunc_Index(t *testing.T) {
	final1 := s.Eval("Integration with TIBCO Web Integrator", "Integrat")
	fmt.Println(final1)
	assert.Equal(t, 27, final1)

	final2 := s.Eval("TIBCO Web Integrator", "rocks")
	fmt.Println(final2)
	assert.Equal(t, -1, final2)

	final3 := s.Eval("Integration with TIBCO 网络 Integrator", "Integrat")
	fmt.Println(final3)
	assert.Equal(t, 30, final3)
}

func TestExpression(t *testing.T) {
	fun, err := expression.ParseExpression(`string.lastIndex("TIBCO NAME", "NAME")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}
