package len

import (
	"fmt"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression"
	"github.com/stretchr/testify/assert"
)

var s = &Len{}

func TestSample(t *testing.T) {
	final := s.Eval("123")
	assert.Equal(t, int(3), final)
}

func TestExpression(t *testing.T) {
	fun, err := expression.ParseExpression(`number.len("TIBCO NAME")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}
