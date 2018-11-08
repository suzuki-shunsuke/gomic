package gencmd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/suzuki-shunsuke/gomic/gomic"
	"github.com/suzuki-shunsuke/gomic/internal/domain"
	"github.com/suzuki-shunsuke/gomic/internal/infra"
	"github.com/suzuki-shunsuke/gomic/internal/test"
)

func Test_findCfg(t *testing.T) {
	fsys := test.NewFileSystemMock(t, nil).
		SetReturnGetwd("/", nil).
		SetReturnExist(false)
	fsys2 := test.NewFileSystemMock(t, nil).
		SetReturnGetwd("/", nil).
		SetReturnExist(true)
	fsys3 := test.NewFileSystemMock(t, nil).
		SetReturnGetwd("/foo/bar", nil).
		SetFuncExist(func(p string) bool {
			return p == "/foo/.gomic.yml"
		})
	fsys4 := test.NewFileSystemMock(t, nil).
		SetReturnGetwd("/foo/bar", nil).
		SetFuncExist(func(p string) bool {
			return p == "/foo/zoo/.gomic.yml"
		})
	data := []struct {
		testcase string
		fsys     domain.FileSystem
		cfgPath  string
		success  bool
	}{
		{"", fsys, "", false},
		{"", fsys2, "/.gomic.yml", true},
		{"", fsys3, "/foo/.gomic.yml", true},
		{"", fsys4, "", false},
	}
	for _, tt := range data {
		t.Run(fmt.Sprintf("%s %s %t", tt.testcase, tt.cfgPath, tt.success), func(t *testing.T) {
			p, err := findCfg(tt.fsys)
			if tt.success {
				assert.Nil(t, err)
				assert.Equal(t, tt.cfgPath, p)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

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
