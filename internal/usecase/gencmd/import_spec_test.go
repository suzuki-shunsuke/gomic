package gencmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_importSpecHost(t *testing.T) {
	spec := importSpec{}
	require.Equal(t, "", spec.Host())
	spec.path = "io"
	require.Equal(t, "", spec.Host())
	spec.path = "github.com/suzuki-shunsuke/gomic"
	require.Equal(t, "github.com", spec.Host())
}

func Test_importSpecString(t *testing.T) {
	spec := importSpec{name: "io", path: "io"}
	require.Equal(t, `io "io"`, spec.String())
	spec.str = `"io"`
	require.Equal(t, `"io"`, spec.String())
}
