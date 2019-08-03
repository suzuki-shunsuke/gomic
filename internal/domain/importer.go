package domain

import (
	"go/ast"
	"go/parser"
	"os"

	"golang.org/x/tools/go/packages"
)

type (
	// Importer imports go's file and packages.
	Importer interface {
		GetFileByFilePath(filePath string, mode parser.Mode) (*ast.File, error)
		GetPkgsInDir(dirPath string, filter func(os.FileInfo) bool, mode parser.Mode) (map[string]*ast.Package, error)
		GetBuildPkgInDir(dir string) (*packages.Package, error)
		GetBuildPkgByPkgPath(pkgPath, srcDir string) (*packages.Package, error)
	}
)
