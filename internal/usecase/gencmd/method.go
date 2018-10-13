package gencmd

import (
	"fmt"
	"go/ast"

	"github.com/suzuki-shunsuke/gomic/internal/domain"
)

func getMethodFromFuncType(
	srcPkg domain.ImportSpec, field *ast.Field, funcType *ast.FuncType,
	isSamePkg bool,
	fileImports map[string]domain.ImportSpec, imports domain.ImportSpecs,
) (Method, domain.ImportSpecs, error) {
	method := Method{name: getNameFromField(field)}

	params, imports, isEllipsis, err := getParams(
		funcType.Params, srcPkg, isSamePkg, fileImports, imports)
	if err != nil {
		return method, nil, err
	}
	method.params = params
	method.isEllipsis = isEllipsis

	results, imports, hasResultNames, err := getResults(
		funcType.Results, srcPkg, isSamePkg, fileImports, imports)
	if err != nil {
		return method, nil, err
	}
	method.hasResultNames = hasResultNames
	method.results = results

	s, err := toString(field.Type)
	if err != nil {
		return method, nil, err
	}
	method.decl = fmt.Sprintf("%s %s", method.name, s)
	method.definition = fmt.Sprintf("%s%s", method.name, s[4:])
	return method, imports, nil
}
