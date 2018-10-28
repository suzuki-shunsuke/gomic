package gencmd

import (
	"fmt"
	"go/ast"

	"github.com/suzuki-shunsuke/gomic/internal/domain"
)

func getResults(
	results *ast.FieldList, srcPkg domain.ImportSpec, isSamePkg bool,
	fileImports map[string]domain.ImportSpec, imports domain.ImportSpecs,
) ([]domain.Var, domain.ImportSpecs, bool, error) {
	if results == nil || len(results.List) == 0 {
		return []domain.Var{}, imports, false, nil
	}
	vars := make([]domain.Var, len(results.List))
	hasResultNames := false
	var err error
	for i, p := range results.List {
		p.Type, imports, err = getImportsInExpr(p.Type, fileImports, imports, srcPkg, isSamePkg)
		if err != nil {
			return nil, nil, false, err
		}
		t, err := toString(p.Type)
		if err != nil {
			return nil, nil, false, err
		}
		name := getNameFromField(p)
		if name == "" {
			name = fmt.Sprintf("r%d", i)
		} else {
			hasResultNames = true
		}
		result := Var{name: name, t: t}
		vars[i] = result
	}
	return vars, imports, hasResultNames, nil
}
