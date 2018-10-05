package stringlength

import (
	"fmt"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression"
	"github.com/stretchr/testify/assert"
)

var s = &StringLength{}

func TestStaticFunc_String_length(t *testing.T) {
	final11 := s.Eval("TIBCO Web Integrator")
	fmt.Println(final11)
	assert.Equal(t, int(20), final11)

	final2 := s.Eval("TIBCO 网路集成器")
	fmt.Println(final2)
	assert.Equal(t, int(21), final2)
}

func TestExpression(t *testing.T) {
	fun, err := expression.ParseExpression(`string.length("seafood,name")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}
