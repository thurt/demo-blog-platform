// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/thurt/demo-blog-platform/cms/proto (interfaces: CmsInternalServer)

// Package mock_proto is a generated GoMock package.
package mock_proto

import (
	gomock "github.com/golang/mock/gomock"
	proto "github.com/thurt/demo-blog-platform/cms/proto"
	context "golang.org/x/net/context"
	reflect "reflect"
)

// MockCmsInternalServer is a mock of CmsInternalServer interface
type MockCmsInternalServer struct {
	ctrl     *gomock.Controller
	recorder *MockCmsInternalServerMockRecorder
}

// MockCmsInternalServerMockRecorder is the mock recorder for MockCmsInternalServer
type MockCmsInternalServerMockRecorder struct {
	mock *MockCmsInternalServer
}

// NewMockCmsInternalServer creates a new mock instance
func NewMockCmsInternalServer(ctrl *gomock.Controller) *MockCmsInternalServer {
	mock := &MockCmsInternalServer{ctrl: ctrl}
	mock.recorder = &MockCmsInternalServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCmsInternalServer) EXPECT() *MockCmsInternalServerMockRecorder {
	return m.recorder
}

// GetUserPassword mocks base method
func (m *MockCmsInternalServer) GetUserPassword(arg0 context.Context, arg1 *proto.UserRequest) (*proto.UserPassword, error) {
	ret := m.ctrl.Call(m, "GetUserPassword", arg0, arg1)
	ret0, _ := ret[0].(*proto.UserPassword)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserPassword indicates an expected call of GetUserPassword
func (mr *MockCmsInternalServerMockRecorder) GetUserPassword(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserPassword", reflect.TypeOf((*MockCmsInternalServer)(nil).GetUserPassword), arg0, arg1)
}