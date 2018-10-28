package gencmd

import (
	"fmt"
	"go/ast"

	"github.com/suzuki-shunsuke/gomic/internal/domain"
)

func getParams(
	prms *ast.FieldList, srcPkg domain.ImportSpec, isSamePkg bool,
	fileImports map[string]domain.ImportSpec, imports domain.ImportSpecs,
) ([]domain.Var, domain.ImportSpecs, bool, error) {
	if prms == nil || len(prms.List) == 0 {
		return []domain.Var{}, imports, false, nil
	}
	size := len(prms.List)
	params := make([]domain.Var, size)
	var err error
	for i, p := range prms.List {
		if p.Type, imports, err = getImportsInExpr(p.Type, fileImports, imports, srcPkg, isSamePkg); err != nil {
			return nil, nil, false, err
		}
		name := getNameFromField(p)
		if name == "" {
			name = fmt.Sprintf("p%d", i)
		}
		param := Var{name: name}
		params[i] = param
		if len(p.Names) == 0 {
			p.Names = []*ast.Ident{ast.NewIdent(fmt.Sprintf("p%d", i))}
		}
	}
	_, isEllipsis := prms.List[size-1].Type.(*ast.Ellipsis)
	return params, imports, isEllipsis, nil
}
