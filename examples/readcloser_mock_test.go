package examples_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/suzuki-shunsuke/gomic/examples"
	"github.com/suzuki-shunsuke/gomic/gomic"
)

func ExampleReadCloserMock() {
	mock := examples.NewReadCloserMock(nil, gomic.DoNothing).
		SetFuncRead(func(p []byte) (int, error) {
			if p == nil {
				return 0, fmt.Errorf("")
			}
			return 1, nil
		})
	n, err := mock.Read(nil)
	fmt.Println(n == 0 && err != nil)
	n, err = mock.Read([]byte{})
	fmt.Println(n == 1 && err == nil)
	// Output:
	// true
	// true
}

func TestReadCloserMockClose(t *testing.T) {
	mock := examples.NewReadCloserMock(t, gomic.DoNothing)
	require.Nil(t, mock.Close())
	mock.SetReturnClose(fmt.Errorf(""))
	require.NotNil(t, mock.Close())
}

func TestReadCloserMockRead(t *testing.T) {
	mock := examples.NewReadCloserMock(t, gomic.DoNothing)
	mock.SetReturnRead(1, nil)
	n, err := mock.Read(nil)
	require.Equal(t, 1, n)
	require.Equal(t, nil, err)
}
