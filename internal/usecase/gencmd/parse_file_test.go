package gencmd

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/suzuki-shunsuke/gomic/internal/domain"
	"github.com/suzuki-shunsuke/gomic/internal/infra"
)

func Test_parseFile(t *testing.T) {
	importer := infra.Importer{}
	bPkg, err := importer.GetBuildPkgByPkgPath("os", "", 0)
	assert.Nil(t, err)
	pkgs, err := importer.GetPkgsInDir(bPkg.Dir, nil, 0)
	assert.Nil(t, err)
	pkg, ok := pkgs[bPkg.Name]
	assert.True(t, ok)
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
		assert.Nil(t, err)
		if intf != nil {
			item.Src.File = f
			break
		}
	}
	assert.NotEqual(t, "", item.Src.File)
	_, err = parseFile(importer, item, item.Src.File)
	assert.Nil(t, err)
	item.Src.Interface = "foo"
	_, err = parseFile(importer, item, item.Src.File)
	assert.NotNil(t, err)
}
