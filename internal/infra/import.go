package infra

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"

	"golang.org/x/tools/go/packages"
)

type (
	// Importer implements domain.Importer .
	Importer struct{}
)

// file path

// GetFileByFilePath implements domain.Importer#GetFileByFilePath .
func (imp Importer) GetFileByFilePath(
	filePath string, mode parser.Mode,
) (*ast.File, error) {
	fset := token.NewFileSet()
	return parser.ParseFile(fset, filePath, nil, mode)
}

// dir path

// GetBuildPkgByPkgPath implements domain.Importer#GetBuildPkgByPkgPath .
func (imp Importer) GetBuildPkgByPkgPath(
	pkgPath, srcDir string,
) (*packages.Package, error) {
	pkgs, err := packages.Load(nil, pkgPath)
	if err != nil {
		return nil, err
	}
	return pkgs[0], nil
}

// pkg path

// GetPkgsInDir implements domain.Importer#GetPkgsInDir .
func (imp Importer) GetPkgsInDir(
	dirPath string, filter func(os.FileInfo) bool, mode parser.Mode,
) (map[string]*ast.Package, error) {
	fset := token.NewFileSet()
	return parser.ParseDir(fset, dirPath, filter, mode)
}

// GetBuildPkgInDir implements domain.Importer#GetBuildPkgInDir .
func (imp Importer) GetBuildPkgInDir(dir string) (*packages.Package, error) {
	pkgs, err := packages.Load(nil, dir)
	if err != nil {
		return nil, err
	}
	return pkgs[0], nil
}
