// Code generated by MockGen. DO NOT EDIT.
// Source: user.trpc.go
//
// Generated by this command:
//
//	mockgen -source=user.trpc.go -destination=mock/user.trpc.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	user "trpc-go-example/proto/user"

	gomock "go.uber.org/mock/gomock"
	client "trpc.group/trpc-go/trpc-go/client"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// GetAccountByUserName mocks base method.
func (m *MockUserService) GetAccountByUserName(ctx context.Context, req *user.GetAccountByUserNameRequest) (*user.GetAccountByUserNameResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountByUserName", ctx, req)
	ret0, _ := ret[0].(*user.GetAccountByUserNameResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountByUserName indicates an expected call of GetAccountByUserName.
func (mr *MockUserServiceMockRecorder) GetAccountByUserName(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountByUserName", reflect.TypeOf((*MockUserService)(nil).GetAccountByUserName), ctx, req)
}

// MockUserClientProxy is a mock of UserClientProxy interface.
type MockUserClientProxy struct {
	ctrl     *gomock.Controller
	recorder *MockUserClientProxyMockRecorder
}

// MockUserClientProxyMockRecorder is the mock recorder for MockUserClientProxy.
type MockUserClientProxyMockRecorder struct {
	mock *MockUserClientProxy
}

// NewMockUserClientProxy creates a new mock instance.
func NewMockUserClientProxy(ctrl *gomock.Controller) *MockUserClientProxy {
	mock := &MockUserClientProxy{ctrl: ctrl}
	mock.recorder = &MockUserClientProxyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserClientProxy) EXPECT() *MockUserClientProxyMockRecorder {
	return m.recorder
}

// GetAccountByUserName mocks base method.
func (m *MockUserClientProxy) GetAccountByUserName(ctx context.Context, req *user.GetAccountByUserNameRequest, opts ...client.Option) (*user.GetAccountByUserNameResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAccountByUserName", varargs...)
	ret0, _ := ret[0].(*user.GetAccountByUserNameResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountByUserName indicates an expected call of GetAccountByUserName.
func (mr *MockUserClientProxyMockRecorder) GetAccountByUserName(ctx, req any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountByUserName", reflect.TypeOf((*MockUserClientProxy)(nil).GetAccountByUserName), varargs...)
}
