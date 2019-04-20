package initcmd

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/suzuki-shunsuke/gomic/gomic"
	"github.com/suzuki-shunsuke/gomic/internal/domain"
	"github.com/suzuki-shunsuke/gomic/internal/test"
)

func TestMain(t *testing.T) {
	d := []byte{}
	fsys := test.NewFileSystemMock(nil, gomic.DoNothing).
		SetFuncWrite(func(dst string, data []byte) error {
			d = data
			return nil
		})
	require.Nil(t, Main(fsys, "/tmp/.gomic.yml"))
	exp := []byte(strings.Trim(domain.ConfigTpl, "\n"))
	require.Equal(t, exp, d)
	d = []byte{}
	fsys.SetReturnExist(true)
	require.Nil(t, Main(fsys, "/tmp/.gomic.yml"))
	require.Equal(t, []byte{}, d)
	fsys.
		SetFuncExist(nil).
		SetReturnMkdirAll(fmt.Errorf("failed to create a directory"))
	require.NotNil(t, Main(fsys, "/tmp/.gomic.yml"))
}
