package gencmd

import (
	"go/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_toString(t *testing.T) {
	a := "[]string"
	arr, err := parser.ParseExpr(a)
	if err != nil {
		t.Fatal(err)
	}
	s, err := toString(arr)
	assert.Nil(t, err)
	assert.Equal(t, a, s)
}

func Test_isPublicIdent(t *testing.T) {
	assert.False(t, isPublicIdent(""))
	assert.False(t, isPublicIdent("foo"))
	assert.True(t, isPublicIdent("Foo"))
}
