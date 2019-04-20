package gencmd

import (
	"go/parser"
	"go/token"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/suzuki-shunsuke/gomic/internal/domain"
	"github.com/suzuki-shunsuke/gomic/internal/infra"
)

func Test_getImportsInFile(t *testing.T) {
	src := `
package main

import (
	"os"
)`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", src, 0)
	require.Nil(t, err)
	importer := infra.Importer{}
	imports, err := getImportsInFile(importer, "", file)
	require.Nil(t, err)
	spec := importSpec{name: "os", path: "os", str: `"os"`}
	require.Equal(t, map[string]domain.ImportSpec{"os": spec}, imports)
}

func Test_getNestedImports(t *testing.T) {
	imports := map[string]domain.ImportSpec{
		"os":      importSpec{name: "os", path: "os"},
		"io":      importSpec{name: "io", path: "io"},
		"rterror": importSpec{name: "rterror", path: "github.com/suzuki-shunsuke/rterror"},
	}
	arr := getNestedImports(imports)
	exp := [][]string{
		{`io "io"`, `os "os"`},
		{`rterror "github.com/suzuki-shunsuke/rterror"`}}
	require.Equal(t, exp, arr)
}
