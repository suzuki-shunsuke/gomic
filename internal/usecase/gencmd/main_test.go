package gencmd

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/suzuki-shunsuke/gomic/gomic"
	"github.com/suzuki-shunsuke/gomic/internal/domain"
	"github.com/suzuki-shunsuke/gomic/internal/infra"
	"github.com/suzuki-shunsuke/gomic/internal/test"
)

func TestMain(t *testing.T) {
	fsys := test.NewFileSystemMock(t, gomic.DoNothing).
		SetReturnGetWriteCloser(test.NewWriteCloserMock(t, gomic.DoNothing), nil)
	cfgReader := test.NewCfgReaderMock(t, gomic.DoNothing).
		SetReturnRead(domain.Config{
			Items: []domain.Item{
				{
					Src: domain.Src{
						Package:   "os",
						Interface: "FileInfo",
						Name:      "FileInfoMock",
					},
					Dest: domain.Dest{
						Package: "examples",
						File:    "/tmp/fileinfo_mock.go",
					},
				},
			},
		}, nil)
	importer := infra.Importer{}
	assert.Nil(t, Main(fsys, importer, cfgReader, "/tmp/.gomic.yml"))
	bPkg, err := importer.GetBuildPkgByPkgPath("os", "", 0)
	assert.Nil(t, err)
	cfgReader.SetReturnRead(domain.Config{
		Items: []domain.Item{
			{
				Src: domain.Src{
					Dir:       bPkg.Dir,
					Interface: "FileInfo",
					Name:      "FileInfoMock",
				},
				Dest: domain.Dest{
					Package: "examples",
					File:    "/tmp/fileinfo_mock.go",
				},
			},
		},
	}, nil)
	assert.Nil(t, Main(fsys, importer, cfgReader, "/tmp/.gomic.yml"))
}

func Test_initCfg(t *testing.T) {
	cfg := domain.Config{
		Default: domain.DefaultConfiguration{
			SrcDefaultConfiguration: domain.SrcDefaultConfiguration{
				InterfacePrefix: "Prefix",
				InterfaceSuffix: "Suffix",
			},
		},
		Items: []domain.Item{
			{
				Src: domain.Src{
					Package:   "os",
					Interface: "FileInfo",
				},
				Dest: domain.Dest{
					Package: "examples",
					File:    "/tmp/fileinfo_mock.go",
				},
			},
		},
	}
	c, err := initCfg(cfg, "/tmp/.gomic.yml")
	assert.Nil(t, err)
	assert.Equal(t, "PrefixFileInfoSuffix", c.Items[0].Src.Name)
}
