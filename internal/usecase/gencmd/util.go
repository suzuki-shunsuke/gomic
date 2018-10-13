package gencmd

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/token"
	"strings"
)

func toString(node interface{}) (string, error) {
	fset := token.NewFileSet()
	var buf bytes.Buffer
	err := format.Node(&buf, fset, node)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func getNameFromField(field *ast.Field) string {
	if len(field.Names) == 0 {
		return ""
	}
	return field.Names[0].Name
}

func isPublicIdent(name string) bool {
	if len(name) == 0 {
		return false
	}
	n := name[:1]
	return strings.ToUpper(n) == n
}
