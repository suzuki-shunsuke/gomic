package examples_test

import (
	"fmt"

	"github.com/suzuki-shunsuke/gomic/examples"
	"github.com/suzuki-shunsuke/gomic/gomic"
)

func ExampleReadCloserMock() {
	mock := examples.NewReadCloserMock(nil, gomic.DoNothing)
	fmt.Println(mock.Close() == nil)
	n, err := mock.Read(nil)
	fmt.Println(n == 0 && err == nil)
	mock.Impl.Read = func(p []byte) (int, error) {
		return 10, nil
	}
	n, err = mock.Read(nil)
	fmt.Println(n == 10 && err == nil)
	mock.Impl.Close = func() error {
		return fmt.Errorf("failed to close")
	}
	fmt.Println(mock.Close() != nil)
	// Output:
	// true
	// true
	// true
	// true
}
