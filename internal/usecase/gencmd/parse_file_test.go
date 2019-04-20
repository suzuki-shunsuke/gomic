package gencmd

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/suzuki-shunsuke/gomic/internal/domain"
	"github.com/suzuki-shunsuke/gomic/internal/infra"
)

func Test_parseFile(t *testing.T) {
	importer := infra.Importer{}
	bPkg, err := importer.GetBuildPkgByPkgPath("os", "", 0)
	require.Nil(t, err)
	pkgs, err := importer.GetPkgsInDir(bPkg.Dir, nil, 0)
	require.Nil(t, err)
	pkg, ok := pkgs[bPkg.Name]
	require.True(t, ok)
	item := domain.Item{
		Src: domain.Src{
			Interface: "FileInfo",
			Name:      "FileInfoMock",
		},
		Dest: domain.Dest{
			Package: "examples",
			File:    "/tmp/fileinfo_mock.go",
		},
	}
	for f, file := range pkg.Files {
		intf, err := getInterfaceInFile(file, item.Src.Interface)
		require.Nil(t, err)
		if intf != nil {
			item.Src.File = f
			break
		}
	}
	require.NotEqual(t, "", item.Src.File)
	_, err = parseFile(importer, item, item.Src.File)
	require.Nil(t, err)
	item.Src.Interface = "foo"
	_, err = parseFile(importer, item, item.Src.File)
	require.NotNil(t, err)
}
