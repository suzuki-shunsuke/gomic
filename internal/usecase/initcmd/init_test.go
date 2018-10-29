package initcmd

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/suzuki-shunsuke/gomic/gomic"
	"github.com/suzuki-shunsuke/gomic/internal/domain"
	"github.com/suzuki-shunsuke/gomic/internal/test"
)

func TestMain(t *testing.T) {
	fsys := test.NewFileSystemMock(nil)
	fsys.CallbackNotImplemented = gomic.DoNothing
	d := []byte{}
	fsys.Impl.Write = func(dst string, data []byte) error {
		d = data
		return nil
	}
	assert.Nil(t, Main(fsys, "/tmp/.gomic.yml"))
	exp := []byte(strings.Trim(domain.ConfigTpl, "\n"))
	assert.Equal(t, exp, d)
	d = []byte{}
	fsys.Impl.Exist = func(dst string) bool {
		return true
	}
	assert.Nil(t, Main(fsys, "/tmp/.gomic.yml"))
	assert.Equal(t, []byte{}, d)
	fsys.Impl.Exist = nil
	fsys.Impl.MkdirAll = func(dir string) error {
		return fmt.Errorf("failed to create a directory")
	}
	assert.NotNil(t, Main(fsys, "/tmp/.gomic.yml"))
}
