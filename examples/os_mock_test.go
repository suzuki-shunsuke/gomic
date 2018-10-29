package examples_test

import (
	"fmt"

	"github.com/suzuki-shunsuke/gomic/examples"
	"github.com/suzuki-shunsuke/gomic/gomic"
)

func ExampleOSMock() {
	hello := examples.NewOSMock(nil)
	hello.CallbackNotImplemented = gomic.DoNothing
	// return default values
	s, err := hello.Getwd()
	fmt.Println(s == "" && err == nil)
	// implement mock function
	hello.Impl.Getwd = func() (string, error) {
		return "/tmp", fmt.Errorf("")
	}
	s, err = hello.Getwd()
	fmt.Println(s == "/tmp" && err != nil)
	// Output:
	// true
	// true
}
