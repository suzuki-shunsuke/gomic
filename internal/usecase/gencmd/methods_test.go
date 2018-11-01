package gencmd

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/suzuki-shunsuke/gomic/internal/domain"
	"github.com/suzuki-shunsuke/gomic/internal/infra"
)

func Test_getMethodsInField(t *testing.T) {
	src := `
package main

type bar interface {}

type Foo interface {
	bar
}
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", src, 0)
	assert.Nil(t, err)
	srcPkg := importSpec{
		name: "rterror", path: "github.com/suzuki-shunsuke/rterror"}
	importer := infra.Importer{}
	fileImports := map[string]domain.ImportSpec{}
	item := domain.Item{
		Src: domain.Src{
			Name: "FooMock",
		},
		Dest: domain.Dest{
			Package: "test",
			File:    "test.go",
		},
	}
	specs := NewImportSpecs()
	field := &ast.Field{Type: ast.NewIdent("bar")}
	_, _, err = getMethodsInField(
		field, file, nil, importer, item, fileImports, specs, srcPkg, false)
	assert.Nil(t, err)
}

func Test_getMethodsInIdent(t *testing.T) {
	src := `
package main

type bar interface {}

type Foo interface {
	bar
}
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", src, 0)
	assert.Nil(t, err)
	srcPkg := importSpec{
		name: "rterror", path: "github.com/suzuki-shunsuke/rterror"}
	importer := infra.Importer{}
	fileImports := map[string]domain.ImportSpec{}
	item := domain.Item{
		Src: domain.Src{
			Name: "FooMock",
		},
		Dest: domain.Dest{
			Package: "test",
			File:    "test.go",
		},
	}
	specs := NewImportSpecs()
	_, _, err = getMethodsInIdent(
		"bar", nil, file, importer, item, fileImports, specs, srcPkg, false)
	assert.Nil(t, err)
	_, _, err = getMethodsInIdent(
		"FileInfo", nil, file, importer, item, fileImports, specs, srcPkg, false)
	assert.NotNil(t, err)
}

func Test_getMethodsInSelectorExpr(t *testing.T) {
	importer := infra.Importer{}
	fileImports := map[string]domain.ImportSpec{
		"os": importSpec{name: "os", path: "os"},
	}
	item := domain.Item{
		Src: domain.Src{
			Name: "FooMock",
		},
		Dest: domain.Dest{
			Package: "test",
			File:    "/tmp/file/test.go",
		},
	}
	specs := NewImportSpecs()
	se := &ast.SelectorExpr{
		X:   ast.NewIdent("os"),
		Sel: ast.NewIdent("FileInfo"),
	}
	_, _, err := getMethodsInSelectorExpr(
		se, importer, item, fileImports, specs)
	assert.Nil(t, err)
}
