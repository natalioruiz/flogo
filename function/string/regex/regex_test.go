package regex

import (
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression"
	"github.com/stretchr/testify/assert"
	"testing"
)

var s = &Regex{}

func TestStaticFunc_Concat(t *testing.T) {
	final := s.Eval("foo.*", "Tseafood")
	fmt.Println(final)
	assert.Equal(t, true, final)

	final2 := s.Eval("bar.*", "seafood")
	fmt.Println(final2)
	assert.Equal(t, false, final2)

	final3 := s.Eval("a(b", "seafood")
	fmt.Println(final3)
	assert.Equal(t, false, final3)
}

func TestExpression(t *testing.T) {
	fun, err := expression.ParseExpression(`string.regex("foo.*","seafood")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}
