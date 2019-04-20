package gencmd

import (
	"go/parser"
	"go/token"
	"testing"

	"github.com/stretchr/testify/require"
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
	require.Nil(t, err)
	intf, err := getInterfaceInFile(file, "Foo")
	require.Nil(t, err)
	require.Equal(t, 1, len(intf.Methods.List))
	require.Equal(t, "Hello", intf.Methods.List[0].Names[0].Name)
}
