// Code generated by MockGen. DO NOT EDIT.
// Source: future.go

// Package future is a generated GoMock package.
package future

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFuture is a mock of Future interface.
type MockFuture struct {
	ctrl     *gomock.Controller
	recorder *MockFutureMockRecorder
}

// MockFutureMockRecorder is the mock recorder for MockFuture.
type MockFutureMockRecorder struct {
	mock *MockFuture
}

// NewMockFuture creates a new mock instance.
func NewMockFuture(ctrl *gomock.Controller) *MockFuture {
	mock := &MockFuture{ctrl: ctrl}
	mock.recorder = &MockFutureMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFuture) EXPECT() *MockFutureMockRecorder {
	return m.recorder
}

// Await mocks base method.
func (m *MockFuture) Await(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Await", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Await indicates an expected call of Await.
func (mr *MockFutureMockRecorder) Await(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Await", reflect.TypeOf((*MockFuture)(nil).Await), arg0)
}