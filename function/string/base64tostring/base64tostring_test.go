package base64tostring

import (
	"fmt"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression"
	"github.com/stretchr/testify/assert"
)

var s = &Base64ToString{}

func TestStaticFunc_Base64_to_string(t *testing.T) {
	final1, err := s.Eval("SGVsbG8sIOS4lueVjA==")
	fmt.Println(final1)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, "Hello, 世界", final1)

	//Negative test case: invalid input base64 string
	final2, err := s.Eval("SSSSGVsbG8sIOS4lueVjA==")
	fmt.Println(final2)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, "", final2)

}

func TestExpression(t *testing.T) {
	fun, err := expression.ParseExpression(`string.base64ToString("SGVsbG8sIOS4lueVjA==")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.Equal(t, "Hello, 世界", v)
	fmt.Println(v)
}
