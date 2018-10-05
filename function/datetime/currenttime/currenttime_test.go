package currenttime

import (
	"fmt"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression"
	"github.com/stretchr/testify/assert"
)

func TestNow_Eval(t *testing.T) {
	n := CurrentTime{}
	date := n.Eval()
	assert.NotNil(t, date)
	fmt.Println(date)
}

func TestExpression(t *testing.T) {
	fun, err := expression.ParseExpression(`datetime.currentTime()`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}
