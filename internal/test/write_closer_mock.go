package test

// Don't edit this file.
// This file is generated by gomic 0.4.0.
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
		impl                   WriteCloserMockImpl
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
	if mock.impl.Write != nil {
		return mock.impl.Write(p)
	}
	if mock.CallbackNotImplemented != nil {
		mock.CallbackNotImplemented(mock.t, mock.name, methodName)
	} else {
		gomic.DefaultCallbackNotImplemented(mock.t, mock.name, methodName)
	}
	return mock.fakeZeroWrite(p)
}

// SetWrite sets a method and returns the mock.
func (mock *WriteCloserMock) SetWrite(impl func(p []byte) (n int, err error)) *WriteCloserMock {
	mock.impl.Write = impl
	return mock
}

// SetFakeWrite sets a fake method.
func (mock *WriteCloserMock) SetFakeWrite(n int, err error) *WriteCloserMock {
	mock.impl.Write = func([]byte) (int, error) {
		return n, err
	}
	return mock
}

// fakeZeroWrite is a fake method which returns zero values.
func (mock WriteCloserMock) fakeZeroWrite(p []byte) (n int, err error) {
	return n, err
}

// Close is a mock method.
func (mock WriteCloserMock) Close() error {
	methodName := "Close" // nolint: goconst
	if mock.impl.Close != nil {
		return mock.impl.Close()
	}
	if mock.CallbackNotImplemented != nil {
		mock.CallbackNotImplemented(mock.t, mock.name, methodName)
	} else {
		gomic.DefaultCallbackNotImplemented(mock.t, mock.name, methodName)
	}
	return mock.fakeZeroClose()
}

// SetClose sets a method and returns the mock.
func (mock *WriteCloserMock) SetClose(impl func() error) *WriteCloserMock {
	mock.impl.Close = impl
	return mock
}

// SetFakeClose sets a fake method.
func (mock *WriteCloserMock) SetFakeClose(r0 error) *WriteCloserMock {
	mock.impl.Close = func() error {
		return r0
	}
	return mock
}

// fakeZeroClose is a fake method which returns zero values.
func (mock WriteCloserMock) fakeZeroClose() error {
	var (
		r0 error
	)
	return r0
}
