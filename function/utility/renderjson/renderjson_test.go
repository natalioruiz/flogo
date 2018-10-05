package renderjson

import (
	"fmt"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression"
	_ "git.tibco.com/git/product/ipaas/wi-contrib.git/function/boolean/false"
	"github.com/stretchr/testify/assert"
)

var s = &RenderJson{}

func TestStaticFunc_RenderJsonStruct(t *testing.T) {
	type Person struct {
		FirstName, LastName string
	}
	person := &Person{FirstName: "Jon", LastName: "Snow"}
	final, err := s.Eval(person, false)
	assert.Nil(t, err)
	fmt.Println(final)
	assert.Equal(t, final, `{"FirstName":"Jon","LastName":"Snow"}`)
}

func TestStaticFunc_RenderJsonMap(t *testing.T) {
	person := make(map[string]string)
	person["FirstName"] = "Jon"
	person["LastName"] = "Snow"
	final, err := s.Eval(person, false)
	assert.Nil(t, err)
	fmt.Println(final)
	assert.Equal(t, final, `{"FirstName":"Jon","LastName":"Snow"}`)
}

func TestExpression(t *testing.T) {
	fun, err := expression.ParseExpression(`utility.renderJSON("{}", boolean.false())`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}
