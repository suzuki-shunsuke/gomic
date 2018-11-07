package examples_test

import (
	"fmt"

	"github.com/suzuki-shunsuke/gomic/examples"
	"github.com/suzuki-shunsuke/gomic/gomic"
)

func ExampleReadCloserMock() {
	mock := examples.NewReadCloserMock(nil, gomic.DoNothing).
		SetRead(func(p []byte) (int, error) {
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
