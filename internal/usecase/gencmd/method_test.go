package gencmd

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/suzuki-shunsuke/gomic/internal/domain"
)

func Test_getMethodFromFuncType(t *testing.T) {
	src := `
package main

import (
	"os"
)

type Foo interface {
	Hello(os.FileInfo)
}
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", src, 0)
	assert.Nil(t, err)
	intf, err := getInterfaceInFile(file, "Foo")
	assert.Nil(t, err)
	field := intf.Methods.List[0]
	funcType := field.Type.(*ast.FuncType)
	srcPkg := importSpec{
		name: "rterror", path: "github.com/suzuki-shunsuke/rterror"}
	fileImports := map[string]domain.ImportSpec{
		"os": importSpec{name: "os", path: "os"},
	}
	method, specs, err := getMethodFromFuncType(srcPkg, field, funcType, false, fileImports, NewImportSpecs())
	assert.Nil(t, err)
	assert.Equal(t, "Hello", method.Name())
	exp := ImportSpecs{
		names: map[string]domain.ImportSpec{
			"os": importSpec{name: "os", path: "os"},
		},
		paths: map[string]domain.ImportSpec{
			"os": importSpec{name: "os", path: "os"},
		},
	}
	assert.Equal(t, &exp, specs)
}
