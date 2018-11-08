package test

// Don't edit this file.
// This file is generated by gomic 0.5.0.
// https://github.com/suzuki-shunsuke/gomic

import (
	"io"
	testing "testing"

	gomic "github.com/suzuki-shunsuke/gomic/gomic"
)

type (
	// FileSystemMock is a mock.
	FileSystemMock struct {
		t                      *testing.T
		name                   string
		callbackNotImplemented gomic.CallbackNotImplemented
		impl                   struct {
			Exist          func(p0 string) bool
			MkdirAll       func(p0 string) error
			Write          func(p0 string, p1 []byte) error
			GetWriteCloser func(p0 string) (io.WriteCloser, error)
			Getwd          func() (string, error)
		}
	}
)

// NewFileSystemMock returns FileSystemMock .
func NewFileSystemMock(t *testing.T, cb gomic.CallbackNotImplemented) *FileSystemMock {
	return &FileSystemMock{t: t, callbackNotImplemented: cb}
}

// Exist is a mock method.
func (mock FileSystemMock) Exist(p0 string) bool {
	methodName := "Exist" // nolint: goconst
	if mock.impl.Exist != nil {
		return mock.impl.Exist(p0)
	}
	if mock.callbackNotImplemented != nil {
		mock.callbackNotImplemented(mock.t, mock.name, methodName)
	} else {
		gomic.DefaultCallbackNotImplemented(mock.t, mock.name, methodName)
	}
	return mock.fakeZeroExist(p0)
}

// SetFuncExist sets a method and returns the mock.
func (mock *FileSystemMock) SetFuncExist(impl func(p0 string) bool) *FileSystemMock {
	mock.impl.Exist = impl
	return mock
}

// SetReturnExist sets a fake method.
func (mock *FileSystemMock) SetReturnExist(r0 bool) *FileSystemMock {
	mock.impl.Exist = func(string) bool {
		return r0
	}
	return mock
}

// fakeZeroExist is a fake method which returns zero values.
func (mock FileSystemMock) fakeZeroExist(p0 string) bool {
	var (
		r0 bool
	)
	return r0
}

// MkdirAll is a mock method.
func (mock FileSystemMock) MkdirAll(p0 string) error {
	methodName := "MkdirAll" // nolint: goconst
	if mock.impl.MkdirAll != nil {
		return mock.impl.MkdirAll(p0)
	}
	if mock.callbackNotImplemented != nil {
		mock.callbackNotImplemented(mock.t, mock.name, methodName)
	} else {
		gomic.DefaultCallbackNotImplemented(mock.t, mock.name, methodName)
	}
	return mock.fakeZeroMkdirAll(p0)
}

// SetFuncMkdirAll sets a method and returns the mock.
func (mock *FileSystemMock) SetFuncMkdirAll(impl func(p0 string) error) *FileSystemMock {
	mock.impl.MkdirAll = impl
	return mock
}

// SetReturnMkdirAll sets a fake method.
func (mock *FileSystemMock) SetReturnMkdirAll(r0 error) *FileSystemMock {
	mock.impl.MkdirAll = func(string) error {
		return r0
	}
	return mock
}

// fakeZeroMkdirAll is a fake method which returns zero values.
func (mock FileSystemMock) fakeZeroMkdirAll(p0 string) error {
	var (
		r0 error
	)
	return r0
}

// Write is a mock method.
func (mock FileSystemMock) Write(p0 string, p1 []byte) error {
	methodName := "Write" // nolint: goconst
	if mock.impl.Write != nil {
		return mock.impl.Write(p0, p1)
	}
	if mock.callbackNotImplemented != nil {
		mock.callbackNotImplemented(mock.t, mock.name, methodName)
	} else {
		gomic.DefaultCallbackNotImplemented(mock.t, mock.name, methodName)
	}
	return mock.fakeZeroWrite(p0, p1)
}

// SetFuncWrite sets a method and returns the mock.
func (mock *FileSystemMock) SetFuncWrite(impl func(p0 string, p1 []byte) error) *FileSystemMock {
	mock.impl.Write = impl
	return mock
}

// SetReturnWrite sets a fake method.
func (mock *FileSystemMock) SetReturnWrite(r0 error) *FileSystemMock {
	mock.impl.Write = func(string, []byte) error {
		return r0
	}
	return mock
}

// fakeZeroWrite is a fake method which returns zero values.
func (mock FileSystemMock) fakeZeroWrite(p0 string, p1 []byte) error {
	var (
		r0 error
	)
	return r0
}

// GetWriteCloser is a mock method.
func (mock FileSystemMock) GetWriteCloser(p0 string) (io.WriteCloser, error) {
	methodName := "GetWriteCloser" // nolint: goconst
	if mock.impl.GetWriteCloser != nil {
		return mock.impl.GetWriteCloser(p0)
	}
	if mock.callbackNotImplemented != nil {
		mock.callbackNotImplemented(mock.t, mock.name, methodName)
	} else {
		gomic.DefaultCallbackNotImplemented(mock.t, mock.name, methodName)
	}
	return mock.fakeZeroGetWriteCloser(p0)
}

// SetFuncGetWriteCloser sets a method and returns the mock.
func (mock *FileSystemMock) SetFuncGetWriteCloser(impl func(p0 string) (io.WriteCloser, error)) *FileSystemMock {
	mock.impl.GetWriteCloser = impl
	return mock
}

// SetReturnGetWriteCloser sets a fake method.
func (mock *FileSystemMock) SetReturnGetWriteCloser(r0 io.WriteCloser, r1 error) *FileSystemMock {
	mock.impl.GetWriteCloser = func(string) (io.WriteCloser, error) {
		return r0, r1
	}
	return mock
}

// fakeZeroGetWriteCloser is a fake method which returns zero values.
func (mock FileSystemMock) fakeZeroGetWriteCloser(p0 string) (io.WriteCloser, error) {
	var (
		r0 io.WriteCloser
		r1 error
	)
	return r0, r1
}

// Getwd is a mock method.
func (mock FileSystemMock) Getwd() (string, error) {
	methodName := "Getwd" // nolint: goconst
	if mock.impl.Getwd != nil {
		return mock.impl.Getwd()
	}
	if mock.callbackNotImplemented != nil {
		mock.callbackNotImplemented(mock.t, mock.name, methodName)
	} else {
		gomic.DefaultCallbackNotImplemented(mock.t, mock.name, methodName)
	}
	return mock.fakeZeroGetwd()
}

// SetFuncGetwd sets a method and returns the mock.
func (mock *FileSystemMock) SetFuncGetwd(impl func() (string, error)) *FileSystemMock {
	mock.impl.Getwd = impl
	return mock
}

// SetReturnGetwd sets a fake method.
func (mock *FileSystemMock) SetReturnGetwd(r0 string, r1 error) *FileSystemMock {
	mock.impl.Getwd = func() (string, error) {
		return r0, r1
	}
	return mock
}

// fakeZeroGetwd is a fake method which returns zero values.
func (mock FileSystemMock) fakeZeroGetwd() (string, error) {
	var (
		r0 string
		r1 error
	)
	return r0, r1
}
