package Int64

import (
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNow_Eval(t *testing.T) {
	n := Now{}
	now := n.Eval()
	assert.NotNil(t, now)
	fmt.Println(now)
}

func TestExpression(t *testing.T) {
	fun, err := expression.ParseExpression(`datetime.now()`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}
