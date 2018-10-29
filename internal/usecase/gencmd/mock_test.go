package gencmd

import (
	"go/parser"
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/suzuki-shunsuke/gomic/internal/domain"
	"github.com/suzuki-shunsuke/gomic/internal/infra"
)

func Test_getMockFromInterface(t *testing.T) {
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
	item := domain.Item{
		Src: domain.Src{
			Name: "FooMock",
		},
		Dest: domain.Dest{
			Package: "test",
			File:    "test.go",
		},
	}
	importer := infra.Importer{}
	srcPkg := importSpec{name: "rterror", path: "github.com/suzuki-shunsuke/rterror"}
	mock, err := getMockFromInterface(intf, nil, file, item, importer, srcPkg, false)
	assert.Nil(t, err)
	assert.Equal(t, "FooMock", mock.MockName())
}
