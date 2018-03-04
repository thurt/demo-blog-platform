// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/thurt/demo-blog-platform/cms/proto (interfaces: EmailerServer)

// Package mock_proto is a generated GoMock package.
package mock_proto

import (
	gomock "github.com/golang/mock/gomock"
	proto "github.com/thurt/demo-blog-platform/cms/proto"
	empty "github.com/thurt/demo-blog-platform/cms/vendor/github.com/golang/protobuf/ptypes/empty"
	context "github.com/thurt/demo-blog-platform/cms/vendor/golang.org/x/net/context"
	reflect "reflect"
)

// MockEmailerServer is a mock of EmailerServer interface
type MockEmailerServer struct {
	ctrl     *gomock.Controller
	recorder *MockEmailerServerMockRecorder
}

// MockEmailerServerMockRecorder is the mock recorder for MockEmailerServer
type MockEmailerServerMockRecorder struct {
	mock *MockEmailerServer
}

// NewMockEmailerServer creates a new mock instance
func NewMockEmailerServer(ctrl *gomock.Controller) *MockEmailerServer {
	mock := &MockEmailerServer{ctrl: ctrl}
	mock.recorder = &MockEmailerServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEmailerServer) EXPECT() *MockEmailerServerMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MockEmailerServer) Send(arg0 context.Context, arg1 *proto.Email) (*empty.Empty, error) {
	ret := m.ctrl.Call(m, "Send", arg0, arg1)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Send indicates an expected call of Send
func (mr *MockEmailerServerMockRecorder) Send(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockEmailerServer)(nil).Send), arg0, arg1)
}
