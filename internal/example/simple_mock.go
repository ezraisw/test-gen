// Code generated by MockGen. DO NOT EDIT.
// Source: simple.go
//
// Generated by this command:
//
//	mockgen -typed -source=simple.go -destination=simple_mock.go -package=example
//

// Package example is a generated GoMock package.
package example

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockSimple is a mock of Simple interface.
type MockSimple struct {
	ctrl     *gomock.Controller
	recorder *MockSimpleMockRecorder
	isgomock struct{}
}

// MockSimpleMockRecorder is the mock recorder for MockSimple.
type MockSimpleMockRecorder struct {
	mock *MockSimple
}

// NewMockSimple creates a new mock instance.
func NewMockSimple(ctrl *gomock.Controller) *MockSimple {
	mock := &MockSimple{ctrl: ctrl}
	mock.recorder = &MockSimpleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSimple) EXPECT() *MockSimpleMockRecorder {
	return m.recorder
}

// A mocks base method.
func (m *MockSimple) A() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "A")
	ret0, _ := ret[0].(string)
	return ret0
}

// A indicates an expected call of A.
func (mr *MockSimpleMockRecorder) A() *MockSimpleACall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "A", reflect.TypeOf((*MockSimple)(nil).A))
	return &MockSimpleACall{Call: call}
}

// MockSimpleACall wrap *gomock.Call
type MockSimpleACall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSimpleACall) Return(arg0 string) *MockSimpleACall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSimpleACall) Do(f func() string) *MockSimpleACall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSimpleACall) DoAndReturn(f func() string) *MockSimpleACall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// B mocks base method.
func (m *MockSimple) B(a int) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "B", a)
	ret0, _ := ret[0].(string)
	return ret0
}

// B indicates an expected call of B.
func (mr *MockSimpleMockRecorder) B(a any) *MockSimpleBCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "B", reflect.TypeOf((*MockSimple)(nil).B), a)
	return &MockSimpleBCall{Call: call}
}

// MockSimpleBCall wrap *gomock.Call
type MockSimpleBCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSimpleBCall) Return(arg0 string) *MockSimpleBCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSimpleBCall) Do(f func(int) string) *MockSimpleBCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSimpleBCall) DoAndReturn(f func(int) string) *MockSimpleBCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// C mocks base method.
func (m *MockSimple) C(b, c int) (string, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "C", b, c)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int)
	return ret0, ret1
}

// C indicates an expected call of C.
func (mr *MockSimpleMockRecorder) C(b, c any) *MockSimpleCCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "C", reflect.TypeOf((*MockSimple)(nil).C), b, c)
	return &MockSimpleCCall{Call: call}
}

// MockSimpleCCall wrap *gomock.Call
type MockSimpleCCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c_2 *MockSimpleCCall) Return(arg0 string, arg1 int) *MockSimpleCCall {
	c_2.Call = c_2.Call.Return(arg0, arg1)
	return c_2
}

// Do rewrite *gomock.Call.Do
func (c_2 *MockSimpleCCall) Do(f func(int, int) (string, int)) *MockSimpleCCall {
	c_2.Call = c_2.Call.Do(f)
	return c_2
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c_2 *MockSimpleCCall) DoAndReturn(f func(int, int) (string, int)) *MockSimpleCCall {
	c_2.Call = c_2.Call.DoAndReturn(f)
	return c_2
}
