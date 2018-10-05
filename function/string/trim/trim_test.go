package trim

import (
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStaticFunc_Trim(t *testing.T) {
	str := " \t\n TIBCO software Inc \n\t\r\n"
	s := &Trim{}
	subStr := s.Eval(str)
	fmt.Println(subStr)
	assert.Equal(t, "TIBCO software Inc", subStr)
}

func TestExpression(t *testing.T) {
	fun, err := expression.ParseExpression(`string.trim("    TIBCO")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}
