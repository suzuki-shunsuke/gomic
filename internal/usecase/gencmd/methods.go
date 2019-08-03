package gencmd

import (
	"fmt"
	"go/ast"
	"path/filepath"

	"github.com/suzuki-shunsuke/gomic/internal/domain"
)

func getMethodsFromInterface(
	intf *ast.InterfaceType, pkg *ast.Package, file *ast.File,
	imports domain.ImportSpecs, item domain.Item,
	importer domain.Importer, srcPkg domain.ImportSpec, isSamePkg bool,
) ([]domain.Method, domain.ImportSpecs, error) {
	// srcPkg is a package which the interface is defined
	methods := []domain.Method{}
	if imports == nil {
		imports = NewImportSpecs()
	}
	fileImports, err := getImportsInFile(importer, item.Src.VendorDir, file)
	if err != nil {
		return methods, imports, err
	}
	for _, field := range intf.Methods.List {
		ms, imps, err := getMethodsInField(
			field, file, pkg, importer, item, fileImports, imports,
			srcPkg, isSamePkg)
		if err != nil {
			return nil, nil, err
		}
		methods = append(methods, ms...)
		imports = imps
	}
	return methods, imports, nil
}

func getMethodsInField(
	field *ast.Field, file *ast.File, pkg *ast.Package,
	importer domain.Importer, item domain.Item,
	fileImports map[string]domain.ImportSpec, imports domain.ImportSpecs,
	srcPkg domain.ImportSpec, isSamePkg bool,
) ([]domain.Method, domain.ImportSpecs, error) {
	// srcPkg is a package which the interface is defined
	if funcType, ok := field.Type.(*ast.FuncType); ok {
		method, imports, err := getMethodFromFuncType(
			srcPkg, field, funcType, isSamePkg, fileImports, imports)
		if err != nil {
			return nil, nil, err
		}
		return []domain.Method{method}, imports, err
	}
	if ident, ok := field.Type.(*ast.Ident); ok {
		return getMethodsInIdent(
			ident.Name, pkg, file, importer, item, fileImports, imports,
			srcPkg, isSamePkg)
	}
	if se, ok := field.Type.(*ast.SelectorExpr); ok {
		return getMethodsInSelectorExpr(se, importer, item, fileImports, imports)
	}
	return nil, nil, fmt.Errorf("unknown type field")
}

func getMethodsInSelectorExpr(
	se *ast.SelectorExpr, importer domain.Importer, item domain.Item,
	fileImports map[string]domain.ImportSpec, imports domain.ImportSpecs,
) ([]domain.Method, domain.ImportSpecs, error) {
	pkgName := se.X.(*ast.Ident).Name
	spec, ok := fileImports[pkgName]
	if !ok {
		return nil, nil, fmt.Errorf("%s is undefined package", pkgName)
	}
	bPkg, err := importer.GetBuildPkgByPkgPath(spec.Path(), item.Src.VendorDir)
	if err != nil {
		return nil, nil, err
	}
	dir := filepath.Dir(bPkg.GoFiles[0])
	pkgs, err := importer.GetPkgsInDir(dir, nil, 0)
	if err != nil {
		return nil, nil, err
	}
	pkg, ok := pkgs[bPkg.Name]
	if !ok {
		return nil, nil, fmt.Errorf(`source package "%s" is not found`, bPkg.Name)
	}
	file, intf, err := getFileAndIntfFromPkg(pkg, se.Sel.Name)
	if err != nil {
		return nil, nil, err
	}

	a, err := filepath.Rel(dir, filepath.Dir(item.Dest.File))
	if err != nil {
		return nil, nil, err
	}
	isSamePkg := a == "." && item.Dest.Package == pkg.Name

	return getMethodsFromInterface(
		intf, pkg, file, imports, item, importer, spec, isSamePkg)
}

func getMethodsInIdent(
	identName string, pkg *ast.Package, file *ast.File,
	importer domain.Importer, item domain.Item,
	fileImports map[string]domain.ImportSpec, imports domain.ImportSpecs,
	srcPkg domain.ImportSpec, isSamePkg bool,
) ([]domain.Method, domain.ImportSpecs, error) {
	var (
		intf *ast.InterfaceType
		err  error
	)
	if pkg == nil {
		intf, err = getInterfaceInFile(file, identName)
	} else {
		file, intf, err = getFileAndIntfFromPkg(pkg, identName)
	}
	if err != nil {
		return nil, nil, err
	}
	if intf == nil {
		return nil, nil, fmt.Errorf("%s is not found", identName)
	}
	return getMethodsFromInterface(
		intf, pkg, file, imports, item, importer, srcPkg, isSamePkg)
}
