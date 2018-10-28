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
	if arr, ok := expr.(*ast.ArrayType); ok {
		arr.Elt, imports, err = getImportsInExpr(arr.Elt, fileImports, imports, srcPkg, isSamePkg)
		return expr, imports, err
	}
	if _, ok := expr.(*ast.BasicLit); ok {
		return nil, nil, fmt.Errorf("ast.BasicLit is invalid type")
	}
	if ch, ok := expr.(*ast.ChanType); ok {
		ch.Value, imports, err = getImportsInExpr(ch.Value, fileImports, imports, srcPkg, isSamePkg)
		return expr, imports, err
	}
	if e, ok := expr.(*ast.Ellipsis); ok {
		e.Elt, imports, err = getImportsInExpr(e.Elt, fileImports, imports, srcPkg, isSamePkg)
		return expr, imports, err
	}
	// if _, ok := expr.(*ast.FuncLit); ok {
	// 	fmt.Println("funcLit")
	// }
	if f, ok := expr.(*ast.FuncType); ok {
		for _, field := range f.Params.List {
			field.Type, imports, err = getImportsInExpr(field.Type, fileImports, imports, srcPkg, isSamePkg)
			if err != nil {
				return nil, nil, err
			}
		}
		for _, field := range f.Results.List {
			field.Type, imports, err = getImportsInExpr(field.Type, fileImports, imports, srcPkg, isSamePkg)
			if err != nil {
				return nil, nil, err
			}
		}
		return expr, imports, nil
	}
	if ident, ok := expr.(*ast.Ident); ok {
		if !isSamePkg && isPublicIdent(ident.Name) {
			if _, err := imports.Add(srcPkg); err != nil {
				return nil, nil, err
			}
			return &ast.SelectorExpr{X: &ast.Ident{Name: srcPkg.Name()}, Sel: ident}, imports, nil
		}
		return expr, imports, nil
	}
	if f, ok := expr.(*ast.MapType); ok {
		f.Key, imports, err = getImportsInExpr(
			f.Key, fileImports, imports, srcPkg, isSamePkg)
		if err != nil {
			return nil, nil, err
		}
		f.Value, imports, err = getImportsInExpr(
			f.Value, fileImports, imports, srcPkg, isSamePkg)
		return f, imports, err
	}
	if se, ok := expr.(*ast.SelectorExpr); ok {
		x := se.X.(*ast.Ident)
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
		return expr, imports, nil
	}
	if slice, ok := expr.(*ast.SliceExpr); ok {
		slice.X, imports, err = getImportsInExpr(slice.X, fileImports, imports, srcPkg, isSamePkg)
		return expr, imports, err
	}
	if star, ok := expr.(*ast.StarExpr); ok {
		star.X, imports, err = getImportsInExpr(star.X, fileImports, imports, srcPkg, isSamePkg)
		return expr, imports, err
	}
	if st, ok := expr.(*ast.StructType); ok {
		for _, field := range st.Fields.List {
			field.Type, imports, err = getImportsInExpr(field.Type, fileImports, imports, srcPkg, isSamePkg)
			if err != nil {
				return nil, nil, err
			}
		}
		return expr, imports, nil
	}
	if intf, ok := expr.(*ast.InterfaceType); ok {
		for _, field := range intf.Methods.List {
			field.Type, imports, err = getImportsInExpr(field.Type, fileImports, imports, srcPkg, isSamePkg)
			if err != nil {
				return nil, nil, err
			}
		}
		return expr, imports, nil
	}
	return nil, nil, fmt.Errorf("unsupported type")
}
