package gencmd

import (
	"fmt"
	"path/filepath"

	"github.com/suzuki-shunsuke/gomic/internal/domain"
)

func parsePkg(
	importer domain.Importer, item domain.Item,
) (domain.MockTemplateArg, error) {
	bPkg, err := importer.GetBuildPkgByPkgPath(item.Src.Package, item.Src.VendorDir)
	if err != nil {
		return nil, err
	}

	dir := filepath.Dir(bPkg.GoFiles[0])

	a, err := filepath.Rel(dir, filepath.Dir(item.Dest.File))
	if err != nil {
		return nil, err
	}
	isSamePkg := a == "." && item.Dest.Package == bPkg.Name

	pkgs, err := importer.GetPkgsInDir(dir, nil, 0)
	if err != nil {
		return nil, err
	}
	pkg, ok := pkgs[bPkg.Name]
	if !ok {
		return nil, fmt.Errorf(`source package "%s" is not found`, bPkg.Name)
	}

	file, intf, err := getFileAndIntfFromPkg(pkg, item.Src.Interface)
	if err != nil {
		return nil, err
	}

	return getMockFromInterface(
		intf, pkg, file, item, importer,
		importSpec{name: pkg.Name, path: item.Src.Package}, isSamePkg)
}
