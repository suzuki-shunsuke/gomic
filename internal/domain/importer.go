package domain

import (
	"go/ast"
	"go/build"
	"go/parser"
	"os"
)

type (
	// Importer imports go's file and packages.
	Importer interface {
		GetFileByFilePath(filePath string, mode parser.Mode) (*ast.File, error)
		GetPkgsInDir(dirPath string, filter func(os.FileInfo) bool, mode parser.Mode) (map[string]*ast.Package, error)
		GetBuildPkgInDir(dir string, mode build.ImportMode) (*build.Package, error)
		GetBuildPkgByPkgPath(pkgPath, srcDir string, mode build.ImportMode) (*build.Package, error)
	}
)
