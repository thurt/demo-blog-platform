// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/thurt/demo-blog-platform/cms/cacher/proto (interfaces: CacherServer)

// Package mock_proto is a generated GoMock package.
package mock_proto

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	proto "github.com/thurt/demo-blog-platform/cms/cacher/proto"
	context "golang.org/x/net/context"
)

// MockCacherServer is a mock of CacherServer interface
type MockCacherServer struct {
	ctrl     *gomock.Controller
	recorder *MockCacherServerMockRecorder
}

// MockCacherServerMockRecorder is the mock recorder for MockCacherServer
type MockCacherServerMockRecorder struct {
	mock *MockCacherServer
}

// NewMockCacherServer creates a new mock instance
func NewMockCacherServer(ctrl *gomock.Controller) *MockCacherServer {
	mock := &MockCacherServer{ctrl: ctrl}
	mock.recorder = &MockCacherServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCacherServer) EXPECT() *MockCacherServerMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockCacherServer) Get(arg0 context.Context, arg1 *proto.GetRequest) (*wrappers.StringValue, error) {
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*wrappers.StringValue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockCacherServerMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCacherServer)(nil).Get), arg0, arg1)
}

// Set mocks base method
func (m *MockCacherServer) Set(arg0 context.Context, arg1 *proto.SetRequest) (*empty.Empty, error) {
	ret := m.ctrl.Call(m, "Set", arg0, arg1)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Set indicates an expected call of Set
func (mr *MockCacherServerMockRecorder) Set(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockCacherServer)(nil).Set), arg0, arg1)
}