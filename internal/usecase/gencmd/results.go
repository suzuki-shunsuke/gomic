package gencmd

import (
	"go/ast"

	"github.com/suzuki-shunsuke/gomic/internal/domain"
)

func getResults(
	results *ast.FieldList, srcPkg domain.ImportSpec, isSamePkg bool,
	fileImports map[string]domain.ImportSpec, imports domain.ImportSpecs,
	idents *Idents,
) ([]domain.Var, domain.ImportSpecs, bool, error) {
	if results == nil || results.NumFields() == 0 {
		return []domain.Var{}, imports, false, nil
	}
	vars := make([]domain.Var, results.NumFields())
	var err error
	i := 0
	for _, p := range results.List {
		p.Type, imports, err = getImportsInExpr(p.Type, fileImports, imports, srcPkg, isSamePkg)
		if err != nil {
			return nil, nil, false, err
		}
		t, err := toString(p.Type)
		if err != nil {
			return nil, nil, false, err
		}
		if len(p.Names) == 0 {
			name := idents.AddNoName("r")
			vars[i] = Var{name: name, t: t}
			p.Names = []*ast.Ident{ast.NewIdent(name)}
			i++
			continue
		}
		for _, ident := range p.Names {
			name := ident.Name
			if name == "" {
				name = idents.AddNoName("r")
			} else {
				name = idents.Add(ident.Name)
			}
			ident.Name = name
			vars[i] = Var{name: name, t: t}
			i++
		}
	}
	return vars, imports, true, nil
}
