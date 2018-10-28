package gencmd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/suzuki-shunsuke/gomic/internal/domain"
	"github.com/suzuki-shunsuke/gomic/internal/test"
)

func Test_findCfg(t *testing.T) {
	fsys := test.NewFileSystemMock(t)
	fsys.Impl.Getwd = func() (string, error) {
		return "/", nil
	}
	fsys.Impl.Exist = func(p string) bool {
		return false
	}
	fsys2 := test.NewFileSystemMock(t)
	fsys2.Impl.Getwd = func() (string, error) {
		return "/", nil
	}
	fsys2.Impl.Exist = func(p string) bool {
		return true
	}
	fsys3 := test.NewFileSystemMock(t)
	fsys3.Impl.Getwd = func() (string, error) {
		return "/foo/bar", nil
	}
	fsys3.Impl.Exist = func(p string) bool {
		return p == "/foo/.gomic.yml"
	}
	fsys4 := test.NewFileSystemMock(t)
	fsys4.Impl.Getwd = func() (string, error) {
		return "/foo/bar", nil
	}
	fsys4.Impl.Exist = func(p string) bool {
		return p == "/foo/zoo/.gomic.yml"
	}
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
