// Code generated by MockGen. DO NOT EDIT.
// Source: gin-study/internal/service/user (interfaces: Server)
//
// Generated by this command:
//
//	mockgen -self_package=gin-study/internal/service/user -destination mock_service.go -package user gin-study/internal/service/user Server
//

// Package user is a generated GoMock package.
package user

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockServer is a mock of Server interface.
type MockServer struct {
	ctrl     *gomock.Controller
	recorder *MockServerMockRecorder
	isgomock struct{}
}

// MockServerMockRecorder is the mock recorder for MockServer.
type MockServerMockRecorder struct {
	mock *MockServer
}

// NewMockServer creates a new mock instance.
func NewMockServer(ctrl *gomock.Controller) *MockServer {
	mock := &MockServer{ctrl: ctrl}
	mock.recorder = &MockServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServer) EXPECT() *MockServerMockRecorder {
	return m.recorder
}

// SendCode mocks base method.
func (m *MockServer) SendCode(arg0 context.Context, arg1 *SendCodeRequest) (*SendCodeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCode", arg0, arg1)
	ret0, _ := ret[0].(*SendCodeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendCode indicates an expected call of SendCode.
func (mr *MockServerMockRecorder) SendCode(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCode", reflect.TypeOf((*MockServer)(nil).SendCode), arg0, arg1)
}
