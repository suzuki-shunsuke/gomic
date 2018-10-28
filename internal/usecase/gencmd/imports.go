package gencmd

import (
	"go/ast"
	"sort"
	"strings"

	"github.com/scylladb/go-set/strset"

	"github.com/suzuki-shunsuke/gomic/internal/domain"
)

func getImportsInFile(
	importer domain.Importer, vendorDir string, file *ast.File,
) (map[string]domain.ImportSpec, error) {
	imports := map[string]domain.ImportSpec{}
	for _, spec := range file.Imports {
		k := ""
		v := ""
		if spec.Name != nil {
			k = strings.Trim(spec.Name.Name, "\"")
		}
		if spec.Path != nil {
			v = strings.Trim(spec.Path.Value, "\"")
		}
		if k == "" && v != "" {
			pkg, err := importer.GetBuildPkgByPkgPath(v, vendorDir, 0)
			if err != nil {
				return nil, err
			}
			k = pkg.Name
		}
		s, err := toString(spec)
		if err != nil {
			return nil, err
		}
		imports[k] = importSpec{str: s, path: v, name: k}
	}
	return imports, nil
}

func getNestedImports(imports map[string]domain.ImportSpec) [][]string {
	mapImports := map[string][]string{}
	keys := strset.New()
	for _, spec := range imports {
		arr, ok := mapImports[spec.Host()]
		if !ok {
			arr = []string{}
		}
		arr = append(arr, spec.String())
		mapImports[spec.Host()] = arr
		keys.Add(spec.Host())
	}
	arr := sort.StringSlice(keys.List())
	sort.Sort(arr)
	nestedImports := make([][]string, len(arr))
	for i, t := range arr {
		imps := sort.StringSlice(mapImports[t])
		sort.Sort(imps)
		nestedImports[i] = imps
	}
	return nestedImports
}
