package gencmd

import (
	"fmt"
	"go/ast"

	"github.com/suzuki-shunsuke/gomic/internal/domain"
)

func getImportsInExpr(
	expr ast.Expr, fileImports map[string]domain.ImportSpec,
	imports domain.ImportSpecs, srcPkg domain.ImportSpec, isSamePkg bool,
) (ast.Expr, domain.ImportSpecs, error) {
	var err error
	switch val := expr.(type) {
	case *ast.ArrayType:
		val.Elt, imports, err = getImportsInExpr(val.Elt, fileImports, imports, srcPkg, isSamePkg)
	case *ast.BasicLit:
		return nil, nil, fmt.Errorf("ast.BasicLit is invalid type")
	case *ast.ChanType:
		val.Value, imports, err = getImportsInExpr(val.Value, fileImports, imports, srcPkg, isSamePkg)
	case *ast.Ellipsis:
		val.Elt, imports, err = getImportsInExpr(val.Elt, fileImports, imports, srcPkg, isSamePkg)
	case *ast.FuncType:
		for _, field := range val.Params.List {
			field.Type, imports, err = getImportsInExpr(field.Type, fileImports, imports, srcPkg, isSamePkg)
			if err != nil {
				return nil, nil, err
			}
		}
		for _, field := range val.Results.List {
			field.Type, imports, err = getImportsInExpr(field.Type, fileImports, imports, srcPkg, isSamePkg)
			if err != nil {
				return nil, nil, err
			}
		}
	case *ast.Ident:
		if !isSamePkg && isPublicIdent(val.Name) {
			if _, err := imports.Add(srcPkg); err != nil {
				return nil, nil, err
			}
			return &ast.SelectorExpr{X: &ast.Ident{Name: srcPkg.Name()}, Sel: val}, imports, nil
		}
	case *ast.MapType:
		val.Key, imports, err = getImportsInExpr(
			val.Key, fileImports, imports, srcPkg, isSamePkg)
		if err != nil {
			return nil, nil, err
		}
		val.Value, imports, err = getImportsInExpr(
			val.Value, fileImports, imports, srcPkg, isSamePkg)
	case *ast.SelectorExpr:
		x := val.X.(*ast.Ident)
		pkgName := x.Name
		spec, ok := fileImports[pkgName]
		if !ok {
			return nil, nil, fmt.Errorf("%s is undefined package", pkgName)
		}
		s, err := imports.Add(spec)
		if err != nil {
			return nil, nil, err
		}
		if pkgName != s.Name() {
			x.Name = s.Name()
		}
	case *ast.SliceExpr:
		val.X, imports, err = getImportsInExpr(val.X, fileImports, imports, srcPkg, isSamePkg)
	case *ast.StarExpr:
		val.X, imports, err = getImportsInExpr(val.X, fileImports, imports, srcPkg, isSamePkg)
	case *ast.StructType:
		for _, field := range val.Fields.List {
			field.Type, imports, err = getImportsInExpr(field.Type, fileImports, imports, srcPkg, isSamePkg)
			if err != nil {
				return nil, nil, err
			}
		}
	case *ast.InterfaceType:
		for _, field := range val.Methods.List {
			field.Type, imports, err = getImportsInExpr(field.Type, fileImports, imports, srcPkg, isSamePkg)
			if err != nil {
				return nil, nil, err
			}
		}
	default:
		// ex. *ast.FuncLit
		return nil, nil, fmt.Errorf("unsupported type")
	}
	return expr, imports, err
}
