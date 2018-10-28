package gencmd

import (
	"fmt"
	"go/ast"
)

func getFileAndIntfFromPkg(pkg *ast.Package, intfName string) (*ast.File, *ast.InterfaceType, error) {
	for _, file := range pkg.Files {
		intf, err := getInterfaceInFile(file, intfName)
		if err != nil {
			return nil, nil, err
		}
		if intf != nil {
			return file, intf, nil
		}
	}
	return nil, nil, fmt.Errorf("%s is not found", intfName)
}

func getInterfaceInFile(file *ast.File, intfName string) (*ast.InterfaceType, error) {
	for _, decl := range file.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			if typeSpec.Name == nil || typeSpec.Name.Name != intfName {
				continue
			}
			intf, ok := typeSpec.Type.(*ast.InterfaceType)
			if !ok {
				return nil, fmt.Errorf("%s is not an interface", intfName)
			}
			return intf, nil
		}
	}
	return nil, nil
}
