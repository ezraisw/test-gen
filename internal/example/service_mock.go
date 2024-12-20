// Code generated by MockGen. DO NOT EDIT.
// Source: service.go
//
// Generated by this command:
//
//	mockgen -typed -source=service.go -destination=service_mock.go -package=example
//

// Package example is a generated GoMock package.
package example

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
	isgomock struct{}
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// A mocks base method.
func (m *MockService) A(ctx context.Context, in int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "A", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// A indicates an expected call of A.
func (mr *MockServiceMockRecorder) A(ctx, in any) *MockServiceACall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "A", reflect.TypeOf((*MockService)(nil).A), ctx, in)
	return &MockServiceACall{Call: call}
}

// MockServiceACall wrap *gomock.Call
type MockServiceACall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockServiceACall) Return(arg0 error) *MockServiceACall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockServiceACall) Do(f func(context.Context, int) error) *MockServiceACall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockServiceACall) DoAndReturn(f func(context.Context, int) error) *MockServiceACall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// B mocks base method.
func (m *MockService) B(ctx context.Context, in int) (*Out, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "B", ctx, in)
	ret0, _ := ret[0].(*Out)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// B indicates an expected call of B.
func (mr *MockServiceMockRecorder) B(ctx, in any) *MockServiceBCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "B", reflect.TypeOf((*MockService)(nil).B), ctx, in)
	return &MockServiceBCall{Call: call}
}

// MockServiceBCall wrap *gomock.Call
type MockServiceBCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockServiceBCall) Return(arg0 *Out, arg1 error) *MockServiceBCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockServiceBCall) Do(f func(context.Context, int) (*Out, error)) *MockServiceBCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockServiceBCall) DoAndReturn(f func(context.Context, int) (*Out, error)) *MockServiceBCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// C mocks base method.
func (m *MockService) C() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "C")
	ret0, _ := ret[0].(int)
	return ret0
}

// C indicates an expected call of C.
func (mr *MockServiceMockRecorder) C() *MockServiceCCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "C", reflect.TypeOf((*MockService)(nil).C))
	return &MockServiceCCall{Call: call}
}

// MockServiceCCall wrap *gomock.Call
type MockServiceCCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockServiceCCall) Return(arg0 int) *MockServiceCCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockServiceCCall) Do(f func() int) *MockServiceCCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockServiceCCall) DoAndReturn(f func() int) *MockServiceCCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
