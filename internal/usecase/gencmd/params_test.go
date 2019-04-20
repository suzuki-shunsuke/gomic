package gencmd

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/suzuki-shunsuke/gomic/internal/domain"
)

func Test_getParams(t *testing.T) {
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
	require.Nil(t, err)
	intf, err := getInterfaceInFile(file, "Foo")
	require.Nil(t, err)
	srcPkg := importSpec{name: "rterror", path: "github.com/suzuki-shunsuke/rterror"}
	prms := intf.Methods.List[0].Type.(*ast.FuncType).Params
	fileImports := map[string]domain.ImportSpec{
		"os": importSpec{name: "os", path: "os"},
	}
	imports := NewImportSpecs()
	params, specs, isEllipsis, err := getParams(
		prms, srcPkg, false, fileImports, imports)
	require.Nil(t, err)
	require.False(t, isEllipsis)
	require.Equal(t, 1, len(params))
	require.Equal(t, 1, len(specs.Names()))
}
