package test

// Don't edit this file.
// This file is generated by gomic 0.5.6.
// https://github.com/suzuki-shunsuke/gomic

import (
	testing "testing"

	gomic "github.com/suzuki-shunsuke/gomic/gomic"
	domain "github.com/suzuki-shunsuke/gomic/internal/domain"
)

type (
	// CfgReaderMock is a mock.
	CfgReaderMock struct {
		t                      *testing.T
		name                   string
		callbackNotImplemented gomic.CallbackNotImplemented
		impl                   struct {
			Read func(p0 string) (r0 domain.Config, r1 error)
		}
	}
)

// NewCfgReaderMock returns CfgReaderMock .
func NewCfgReaderMock(t *testing.T, cb gomic.CallbackNotImplemented) *CfgReaderMock {
	return &CfgReaderMock{
		t: t, name: "CfgReaderMock", callbackNotImplemented: cb}
}

// Read is a mock method.
func (mock CfgReaderMock) Read(p0 string) (r0 domain.Config, r1 error) {
	methodName := "Read" // nolint: goconst
	if mock.impl.Read != nil {
		return mock.impl.Read(p0)
	}
	if mock.callbackNotImplemented != nil {
		mock.callbackNotImplemented(mock.t, mock.name, methodName)
	} else {
		gomic.DefaultCallbackNotImplemented(mock.t, mock.name, methodName)
	}
	return mock.fakeZeroRead(p0)
}

// SetFuncRead sets a method and returns the mock.
func (mock *CfgReaderMock) SetFuncRead(impl func(p0 string) (r0 domain.Config, r1 error)) *CfgReaderMock {
	mock.impl.Read = impl
	return mock
}

// SetReturnRead sets a fake method.
func (mock *CfgReaderMock) SetReturnRead(r0 domain.Config, r1 error) *CfgReaderMock {
	mock.impl.Read = func(string) (domain.Config, error) {
		return r0, r1
	}
	return mock
}

// fakeZeroRead is a fake method which returns zero values.
func (mock CfgReaderMock) fakeZeroRead(p0 string) (r0 domain.Config, r1 error) {
	return r0, r1
}
