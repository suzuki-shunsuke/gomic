package gencmd

import (
	"fmt"
	"path/filepath"

	"github.com/suzuki-shunsuke/gomic/internal/domain"
)

func parseDir(
	importer domain.Importer, item domain.Item, srcDirPath string,
) (domain.MockTemplateArg, error) {
	pkgs, err := importer.GetPkgsInDir(srcDirPath, nil, 0)
	if err != nil {
		return nil, err
	}
	pkg, err := importer.GetBuildPkgInDir(srcDirPath)
	if err != nil {
		return nil, err
	}
	tPkg, ok := pkgs[pkg.Name]
	if !ok {
		return nil, fmt.Errorf(`source package "%s" is not found`, pkg.Name)
	}
	a, err := filepath.Rel(item.Src.Dir, filepath.Dir(item.Dest.File))
	if err != nil {
		return nil, err
	}
	isSamePkg := a == "." && item.Dest.Package == pkg.Name

	file, intf, err := getFileAndIntfFromPkg(tPkg, item.Src.Interface)
	if err != nil {
		return nil, err
	}

	return getMockFromInterface(
		intf, tPkg, file, item, importer,
		importSpec{name: pkg.Name, path: pkg.PkgPath},
		isSamePkg)
}
