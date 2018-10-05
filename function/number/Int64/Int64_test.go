package Int64

import (
	"fmt"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression"
	"github.com/stretchr/testify/assert"
)

var s = &Int{}

func TestSample(t *testing.T) {
	final, err := s.Eval("123")
	assert.Nil(t, err)
	assert.Equal(t, int(123), final)
}

func TestExpression(t *testing.T) {
	fun, err := expression.ParseExpression(`number.int64("123")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.Equal(t, int(123), v)
	fmt.Println(v)
}
