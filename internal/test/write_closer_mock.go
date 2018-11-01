package test

// Don't edit this file.
// This file is generated by gomic 0.2.1.
// https://github.com/suzuki-shunsuke/gomic

import (
	testing "testing"

	gomic "github.com/suzuki-shunsuke/gomic/gomic"
)

type (
	// WriteCloserMock is a mock.
	WriteCloserMock struct {
		t                      *testing.T
		name                   string
		CallbackNotImplemented gomic.CallbackNotImplemented
		Impl                   WriteCloserMockImpl
	}

	// WriteCloserMockImpl holds functions which implement interface's methods.
	WriteCloserMockImpl struct {
		Write func(p []byte) (n int, err error)
		Close func() error
	}
)

// NewWriteCloserMock returns WriteCloserMock .
func NewWriteCloserMock(t *testing.T, cb gomic.CallbackNotImplemented) *WriteCloserMock {
	return &WriteCloserMock{t: t, CallbackNotImplemented: cb}
}

// Write is a mock method.
func (mock WriteCloserMock) Write(p []byte) (n int, err error) {
	methodName := "Write" // nolint: goconst
	if mock.Impl.Write != nil {
		return mock.Impl.Write(p)
	}
	if mock.CallbackNotImplemented != nil {
		mock.CallbackNotImplemented(mock.t, mock.name, methodName)
	} else {
		gomic.DefaultCallbackNotImplemented(mock.t, mock.name, methodName)
	}
	return mock.FakeWrite(p)
}

// FakeWrite is a fake method.
func (mock WriteCloserMock) FakeWrite(p []byte) (n int, err error) {
	return n, err
}

// Close is a mock method.
func (mock WriteCloserMock) Close() error {
	methodName := "Close" // nolint: goconst
	if mock.Impl.Close != nil {
		return mock.Impl.Close()
	}
	if mock.CallbackNotImplemented != nil {
		mock.CallbackNotImplemented(mock.t, mock.name, methodName)
	} else {
		gomic.DefaultCallbackNotImplemented(mock.t, mock.name, methodName)
	}
	return mock.FakeClose()
}

// FakeClose is a fake method.
func (mock WriteCloserMock) FakeClose() error {
	var (
		r0 error
	)
	return r0
}
