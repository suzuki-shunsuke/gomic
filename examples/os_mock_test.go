package examples_test

import (
	"fmt"

	"github.com/suzuki-shunsuke/gomic/examples"
	"github.com/suzuki-shunsuke/gomic/gomic"
)

func ExampleOSMock() {
	mock := examples.NewOSMock(nil, gomic.DoNothing)
	// return zero values
	s, err := mock.Getwd()
	fmt.Println(s == "" && err == nil)
	// set return values
	mock.SetFakeGetwd("foo", nil)
	s, err = mock.Getwd()
	fmt.Println(s == "foo" && err == nil)
	// implement mock method
	mock.Impl.Getwd = func() (string, error) {
		return "/tmp", fmt.Errorf("")
	}
	s, err = mock.Getwd()
	fmt.Println(s == "/tmp" && err != nil)
	// Output:
	// true
	// true
	// true
}
