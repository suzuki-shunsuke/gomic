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
	if prms == nil || prms.NumFields() == 0 {
		return []domain.Var{}, imports, false, nil
	}
	params := make([]domain.Var, prms.NumFields())
	var err error
	i := 0
	for _, p := range prms.List {
		if p.Type, imports, err = getImportsInExpr(p.Type, fileImports, imports, srcPkg, isSamePkg); err != nil {
			return nil, nil, false, err
		}
		if len(p.Names) == 0 {
			params[i] = Var{name: fmt.Sprintf("p%d", i)}
			i++
			continue
		}
		for _, ident := range p.Names {
			if ident.Name == "" {
				ident.Name = fmt.Sprintf("p%d", i)
			}
			params[i] = Var{name: ident.Name}
			i++
		}
	}
	_, isEllipsis := prms.List[len(prms.List)-1].Type.(*ast.Ellipsis)
	return params, imports, isEllipsis, nil
}
