package gencmd

import (
	"fmt"
	"path/filepath"

	"github.com/suzuki-shunsuke/gomic/internal/domain"
)

func parseFile(
	importer domain.Importer, item domain.Item, filePath string,
) (domain.MockTemplateArg, error) {
	// get source pkg name and pkg path
	pkg, err := importer.GetBuildPkgInDir(filepath.Dir(filePath), 0)
	if err != nil {
		return nil, err
	}

	a, err := filepath.Rel(filepath.Dir(item.Src.File), filepath.Dir(item.Dest.File))
	if err != nil {
		return nil, err
	}

	isSamePkg := a == "." && item.Dest.Package == pkg.Name

	file, err := importer.GetFileByFilePath(filePath, 0)
	if err != nil {
		return nil, err
	}

	intf, err := getInterfaceInFile(file, item.Src.Interface)
	if err != nil {
		return nil, err
	}
	if intf == nil {
		return nil, fmt.Errorf("%s is not found", item.Src.Interface)
	}

	return getMockFromInterface(
		intf, nil, file, item, importer,
		importSpec{name: pkg.Name, path: pkg.ImportPath}, isSamePkg)
}
