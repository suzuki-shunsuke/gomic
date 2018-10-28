package gencmd

import (
	"go/parser"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/suzuki-shunsuke/gomic/internal/domain"
)

func Test_getImportsInExpr(t *testing.T) {
	t.Run("ArrayType", func(t *testing.T) {
		arr, err := parser.ParseExpr("[]string")
		assert.Nil(t, err)
		_, _, err = getImportsInExpr(
			arr, map[string]domain.ImportSpec{},
			NewImportSpecs(), importSpec{}, true)
		assert.Nil(t, err)
	})

	t.Run("ChanType", func(t *testing.T) {
		ch, err := parser.ParseExpr("chan string")
		assert.Nil(t, err)
		_, _, err = getImportsInExpr(
			ch, map[string]domain.ImportSpec{}, NewImportSpecs(), importSpec{}, true)
		assert.Nil(t, err)
	})

	t.Run("FuncType", func(t *testing.T) {
		f, err := parser.ParseExpr("func() string")
		assert.Nil(t, err)
		_, _, err = getImportsInExpr(
			f, map[string]domain.ImportSpec{}, NewImportSpecs(),
			importSpec{}, true)
		assert.Nil(t, err)
	})

	t.Run("Ident", func(t *testing.T) {
		ident, err := parser.ParseExpr("string")
		assert.Nil(t, err)
		_, _, err = getImportsInExpr(
			ident, map[string]domain.ImportSpec{},
			NewImportSpecs(), importSpec{}, true)
		assert.Nil(t, err)
	})

	t.Run("MapType", func(t *testing.T) {
		m, err := parser.ParseExpr("map[string]string")
		assert.Nil(t, err)
		_, _, err = getImportsInExpr(
			m, map[string]domain.ImportSpec{}, NewImportSpecs(), importSpec{}, true)
		assert.Nil(t, err)
	})

	t.Run("SelectorExpr", func(t *testing.T) {
		se, err := parser.ParseExpr("os.FileInfo")
		assert.Nil(t, err)
		fileImports := map[string]domain.ImportSpec{
			"os": importSpec{name: "os", path: "os"}}
		_, _, err = getImportsInExpr(
			se, fileImports, NewImportSpecs(), importSpec{}, true)
		assert.Nil(t, err)
	})

	t.Run("StarExpr", func(t *testing.T) {
		star, err := parser.ParseExpr("*string")
		assert.Nil(t, err)
		_, _, err = getImportsInExpr(
			star, map[string]domain.ImportSpec{}, NewImportSpecs(), importSpec{}, true)
		assert.Nil(t, err)
	})

	t.Run("StructType", func(t *testing.T) {
		st, err := parser.ParseExpr("struct{Foo string}")
		assert.Nil(t, err)
		_, _, err = getImportsInExpr(
			st, map[string]domain.ImportSpec{}, NewImportSpecs(), importSpec{}, true)
		assert.Nil(t, err)
	})

	t.Run("InterfaceType", func(t *testing.T) {
		intf, err := parser.ParseExpr("interface{Foo(string) string}")
		assert.Nil(t, err)
		_, _, err = getImportsInExpr(
			intf, map[string]domain.ImportSpec{}, NewImportSpecs(), importSpec{}, true)
		assert.Nil(t, err)
	})
}
