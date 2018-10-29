package gencmd

import (
	"go/parser"
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getInterfaceInFile(t *testing.T) {
	src := `
package main

type Foo interface {
	Hello()
}
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", src, 0)
	assert.Nil(t, err)
	intf, err := getInterfaceInFile(file, "Foo")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(intf.Methods.List))
	assert.Equal(t, "Hello", intf.Methods.List[0].Names[0].Name)
}
