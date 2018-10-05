package false

import (
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression"
	"github.com/stretchr/testify/assert"
	"testing"
)

var s = &False{}

func TestStaticFunc_Contains(t *testing.T) {
	final1 := s.Eval()
	fmt.Println(final1)
	assert.Equal(t, false, final1)
}

func TestExpression(t *testing.T) {
	fun, err := expression.ParseExpression(`boolean.false()`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.Equal(t, false, v)
}
