// Code generated by MockGen. DO NOT EDIT.
// Source: example/foobar.go

// Package example is a generated GoMock package.
package example

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFooBarAble is a mock of FooBarAble interface.
type MockFooBarAble struct {
	ctrl     *gomock.Controller
	recorder *MockFooBarAbleMockRecorder
}

// MockFooBarAbleMockRecorder is the mock recorder for MockFooBarAble.
type MockFooBarAbleMockRecorder struct {
	mock *MockFooBarAble
}

// NewMockFooBarAble creates a new mock instance.
func NewMockFooBarAble(ctrl *gomock.Controller) *MockFooBarAble {
	mock := &MockFooBarAble{ctrl: ctrl}
	mock.recorder = &MockFooBarAbleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFooBarAble) EXPECT() *MockFooBarAbleMockRecorder {
	return m.recorder
}

// WithMap mocks base method.
func (m *MockFooBarAble) WithMap(arg1 map[string]int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WithMap", arg1)
}

// WithMap indicates an expected call of WithMap.
func (mr *MockFooBarAbleMockRecorder) WithMap(arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithMap", reflect.TypeOf((*MockFooBarAble)(nil).WithMap), arg1)
}

// WithSlice mocks base method.
func (m *MockFooBarAble) WithSlice(arg1 []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WithSlice", arg1)
}

// WithSlice indicates an expected call of WithSlice.
func (mr *MockFooBarAbleMockRecorder) WithSlice(arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithSlice", reflect.TypeOf((*MockFooBarAble)(nil).WithSlice), arg1)
}

// WithString mocks base method.
func (m *MockFooBarAble) WithString(arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WithString", arg1)
}

// WithString indicates an expected call of WithString.
func (mr *MockFooBarAbleMockRecorder) WithString(arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithString", reflect.TypeOf((*MockFooBarAble)(nil).WithString), arg1)
}
