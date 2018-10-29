package examples_test

import (
	"fmt"

	"github.com/suzuki-shunsuke/gomic/examples"
	"github.com/suzuki-shunsuke/gomic/gomic"
)

func ExampleOSMock() {
	mock := examples.NewOSMock(nil, gomic.DoNothing)
	// return default values
	s, err := mock.Getwd()
	fmt.Println(s == "" && err == nil)
	// implement mock function
	mock.Impl.Getwd = func() (string, error) {
		return "/tmp", fmt.Errorf("")
	}
	s, err = mock.Getwd()
	fmt.Println(s == "/tmp" && err != nil)
	// Output:
	// true
	// true
}
