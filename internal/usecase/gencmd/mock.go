package gencmd

import (
	"go/ast"

	"github.com/suzuki-shunsuke/gomic/internal/domain"
)

func getMockFromInterface(
	intf *ast.InterfaceType, pkg *ast.Package, file *ast.File, item domain.Item,
	importer domain.Importer, srcPkg domain.ImportSpec, isSamePkg bool,
) (domain.MockTemplateArg, error) {
	imports := NewImportSpecs()
	if _, err := imports.Add(importSpec{name: "testing", path: "testing"}); err != nil {
		return nil, err
	}
	if _, err := imports.Add(importSpec{name: "gomic", path: "github.com/suzuki-shunsuke/gomic/gomic"}); err != nil {
		return nil, err
	}
	methods, imps, err := getMethodsFromInterface(
		intf, pkg, file, imports, item, importer, srcPkg, isSamePkg)
	if err != nil {
		return nil, err
	}

	return MockTemplateArg{
		methods: methods,
		Item:    item,
		imports: getNestedImports(imps.Names()),
	}, nil
}
