package not

import (
	"fmt"

	_ "git.tibco.com/git/product/ipaas/wi-contrib.git/function/boolean/false"
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression"

	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var s = &Not{}

func TestStaticFunc_Contains(t *testing.T) {
	final1 := s.Eval(true)
	fmt.Println(final1)
	assert.Equal(t, false, final1)
}

func TestExpression(t *testing.T) {
	fun, err := expression.ParseExpression(`boolean.not(boolean.false())`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v1, _ := json.Marshal(fun)
	fmt.Println(string(v1))
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.Equal(t, true, v)
}
