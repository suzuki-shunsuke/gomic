package gencmd

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/suzuki-shunsuke/gomic/internal/domain"
	"github.com/suzuki-shunsuke/gomic/internal/infra"
)

func Test_parsePkg(t *testing.T) {
	importer := infra.Importer{}
	item := domain.Item{
		Src: domain.Src{
			Package:   "os",
			Interface: "FileInfo",
			Name:      "FileInfoMock",
		},
		Dest: domain.Dest{
			Package: "examples",
			File:    "/tmp/fileinfo_mock.go",
		},
	}
	_, err := parsePkg(importer, item)
	assert.Nil(t, err)
}
