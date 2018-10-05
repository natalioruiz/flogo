package index

import (
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression"
	"github.com/stretchr/testify/assert"
	"testing"
)

var s = &Index{}

func TestStaticFunc_Index(t *testing.T) {
	final1 := s.Eval("TIBCO Web Integrator", "Web")
	fmt.Println(final1)
	assert.Equal(t, 6, final1)

	final2 := s.Eval("TIBCO Web Integrator", "Internet")
	fmt.Println(final2)
	assert.Equal(t, -1, final2)

	final3 := s.Eval("TIBCO 网络 Integrator", "Integrator")
	fmt.Println(final3)
	assert.Equal(t, 13, final3)
}

func TestExpression(t *testing.T) {
	fun, err := expression.ParseExpression(`string.index("TIBCO NAME", "NAME")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}
