package gencmd

import (
	"fmt"
	"go/ast"

	"github.com/scylladb/go-set/strset"

	"github.com/suzuki-shunsuke/gomic/internal/domain"
)

type (
	Idents struct {
		idents *strset.Set
	}
)

func newIdents() *Idents {
	return &Idents{
		idents: strset.New("mock"),
	}
}

func (idents *Idents) AddNoName(ident string) string {
	for i := 0; i < 100; i++ {
		s := fmt.Sprintf("%s%d", ident, i)
		if !idents.idents.Has(s) {
			idents.idents.Add(s)
			return s
		}
	}
	return ""
}

func (idents *Idents) Add(ident string) string {
	if !idents.idents.Has(ident) {
		idents.idents.Add(ident)
		return ident
	}
	return idents.AddNoName(ident)
}

func getParams(
	prms *ast.FieldList, srcPkg domain.ImportSpec, isSamePkg bool,
	fileImports map[string]domain.ImportSpec, imports domain.ImportSpecs,
	idents *Idents,
) ([]domain.Var, domain.ImportSpecs, bool, error) {
	// srcPkg is a package which the interface is defined
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
			name := idents.AddNoName("p")
			params[i] = Var{name: name}
			p.Names = []*ast.Ident{ast.NewIdent(name)}
			i++
			continue
		}
		for _, ident := range p.Names {
			if ident.Name == "" {
				ident.Name = idents.AddNoName("p")
			} else {
				ident.Name = idents.Add(ident.Name)
			}
			params[i] = Var{name: ident.Name}
			i++
		}
	}
	_, isEllipsis := prms.List[len(prms.List)-1].Type.(*ast.Ellipsis)
	return params, imports, isEllipsis, nil
}
