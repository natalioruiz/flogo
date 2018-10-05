package true

import (
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression"
	"github.com/stretchr/testify/assert"
	"testing"
)

var s = &True{}

func TestStaticFunc_Contains(t *testing.T) {
	final1 := s.Eval()
	fmt.Println(final1)
	assert.Equal(t, true, final1)
}

func TestExpression(t *testing.T) {
	fun, err := expression.ParseExpression(`boolean.true()`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.Equal(t, true, v)
}
