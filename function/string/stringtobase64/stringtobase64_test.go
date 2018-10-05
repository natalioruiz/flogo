package stringtobase64

import (
	"fmt"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression"
	"github.com/stretchr/testify/assert"
)

var s = &StringToBase64{}

func TestStaticFunc_String_to_base64(t *testing.T) {
	final1 := s.Eval("Hello, 世界")
	fmt.Println(final1)
	assert.Equal(t, "SGVsbG8sIOS4lueVjA==", final1)

}

func TestExpression(t *testing.T) {
	fun, err := expression.ParseExpression(`string.stringToBase64("Hello, 世界")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}
