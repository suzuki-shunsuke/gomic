package examples_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/suzuki-shunsuke/gomic/examples"
	"github.com/suzuki-shunsuke/gomic/gomic"
)

func ExampleOSMock() {
	mock := examples.NewOSMock(nil, gomic.DoNothing)
	// return zero values
	s, err := mock.Getwd()
	fmt.Println(s == "" && err == nil)
	// set return values
	mock.SetReturnGetwd("foo", nil)
	s, err = mock.Getwd()
	fmt.Println(s == "foo" && err == nil)
	// implement mock method
	mock.SetFuncGetwd(func() (string, error) {
		return "/tmp", fmt.Errorf("")
	})
	s, err = mock.Getwd()
	fmt.Println(s == "/tmp" && err != nil)
	// Output:
	// true
	// true
	// true
}

func TestOSMockMkdir(t *testing.T) {
	mock := examples.NewOSMock(t, gomic.DoNothing)
	assert.Nil(t, mock.Mkdir("", 0))
	mock.SetFuncMkdir(func(name string, perm os.FileMode) error {
		return nil
	})
	assert.Nil(t, mock.Mkdir("", 0))
	mock.SetReturnMkdir(fmt.Errorf(""))
	assert.NotNil(t, mock.Mkdir("", 0))
}
