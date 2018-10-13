package examples_test

import (
	"fmt"
	"testing"

	"github.com/suzuki-shunsuke/gomic/examples"
)

func ExampleHelloMock4() {
	hello := examples.NewHelloMock4(nil)
	hello.CallbackNotImplemented = func(t *testing.T, intf, method string) {}
	fmt.Println(hello.ExistFile("foo"))
	hello.Impl.ExistFile = func(a string) bool {
		return true
	}
	fmt.Println(hello.ExistFile("foo"))
	// Output:
	// false
	// true
}
