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
	method.definition = s[4:]
	if len(results) != 0 {
		arr := make([]*ast.Field, len(funcType.Results.List))
		noNamedResultArr := make([]*ast.Field, len(funcType.Results.List))
		for i, result := range funcType.Results.List {
			names := result.Names
			if len(result.Names) == 0 {
				names = []*ast.Ident{ast.NewIdent(fmt.Sprintf("r%d", i))}
			}
			arr[i] = &ast.Field{Names: names, Type: result.Type}
			noNamedResultArr[i] = &ast.Field{Type: result.Type}
		}
		noNamedParamArr := make([]*ast.Field, len(funcType.Params.List))
		for i, param := range funcType.Params.List {
			noNamedParamArr[i] = &ast.Field{Type: param.Type}
		}
		setReturnStr, err := toString(&ast.FuncType{Params: &ast.FieldList{List: arr}})
		if err != nil {
			return method, nil, err
		}
		method.setReturnDefinition = setReturnStr[4:]
		s, err := toString(&ast.FuncType{
			Params:  &ast.FieldList{List: noNamedParamArr},
			Results: &ast.FieldList{List: noNamedResultArr}})
		if err != nil {
			return method, nil, err
		}
		method.setReturnInternalDefinition = s[4:]
	}
	return method, imports, nil
}
