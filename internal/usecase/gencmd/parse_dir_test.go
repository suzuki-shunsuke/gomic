package gencmd

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/suzuki-shunsuke/gomic/internal/domain"
	"github.com/suzuki-shunsuke/gomic/internal/infra"
)

func Test_parseDir(t *testing.T) {
	importer := infra.Importer{}
	bPkg, err := importer.GetBuildPkgByPkgPath("os", "")
	require.Nil(t, err)
	item := domain.Item{
		Src: domain.Src{
			Dir:       filepath.Dir(bPkg.GoFiles[0]),
			Interface: "FileInfo",
			Name:      "FileInfoMock",
		},
		Dest: domain.Dest{
			Package: "examples",
			File:    "/tmp/fileinfo_mock.go",
		},
	}
	_, err = parseDir(importer, item, item.Src.Dir)
	require.Nil(t, err)
}
