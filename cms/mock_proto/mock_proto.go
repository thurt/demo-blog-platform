// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/thurt/demo-blog-platform/cms/proto (interfaces: CmsServer)

// Package mock_proto is a generated GoMock package.
package mock_proto

import (
	gomock "github.com/golang/mock/gomock"
	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	proto "github.com/thurt/demo-blog-platform/cms/proto"
	context "golang.org/x/net/context"
	reflect "reflect"
)

// MockCmsServer is a mock of CmsServer interface
type MockCmsServer struct {
	ctrl     *gomock.Controller
	recorder *MockCmsServerMockRecorder
}

// MockCmsServerMockRecorder is the mock recorder for MockCmsServer
type MockCmsServerMockRecorder struct {
	mock *MockCmsServer
}

// NewMockCmsServer creates a new mock instance
func NewMockCmsServer(ctrl *gomock.Controller) *MockCmsServer {
	mock := &MockCmsServer{ctrl: ctrl}
	mock.recorder = &MockCmsServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCmsServer) EXPECT() *MockCmsServerMockRecorder {
	return m.recorder
}

// AuthUser mocks base method
func (m *MockCmsServer) AuthUser(arg0 context.Context, arg1 *proto.AuthUserRequest) (*proto.AccessToken, error) {
	ret := m.ctrl.Call(m, "AuthUser", arg0, arg1)
	ret0, _ := ret[0].(*proto.AccessToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthUser indicates an expected call of AuthUser
func (mr *MockCmsServerMockRecorder) AuthUser(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthUser", reflect.TypeOf((*MockCmsServer)(nil).AuthUser), arg0, arg1)
}

// CreateComment mocks base method
func (m *MockCmsServer) CreateComment(arg0 context.Context, arg1 *proto.CreateCommentRequest) (*proto.CommentRequest, error) {
	ret := m.ctrl.Call(m, "CreateComment", arg0, arg1)
	ret0, _ := ret[0].(*proto.CommentRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateComment indicates an expected call of CreateComment
func (mr *MockCmsServerMockRecorder) CreateComment(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComment", reflect.TypeOf((*MockCmsServer)(nil).CreateComment), arg0, arg1)
}

// CreatePost mocks base method
func (m *MockCmsServer) CreatePost(arg0 context.Context, arg1 *proto.CreatePostRequest) (*proto.PostRequest, error) {
	ret := m.ctrl.Call(m, "CreatePost", arg0, arg1)
	ret0, _ := ret[0].(*proto.PostRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePost indicates an expected call of CreatePost
func (mr *MockCmsServerMockRecorder) CreatePost(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockCmsServer)(nil).CreatePost), arg0, arg1)
}

// CreateUser mocks base method
func (m *MockCmsServer) CreateUser(arg0 context.Context, arg1 *proto.CreateUserRequest) (*proto.UserRequest, error) {
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(*proto.UserRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser
func (mr *MockCmsServerMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockCmsServer)(nil).CreateUser), arg0, arg1)
}

// DeleteComment mocks base method
func (m *MockCmsServer) DeleteComment(arg0 context.Context, arg1 *proto.CommentRequest) (*empty.Empty, error) {
	ret := m.ctrl.Call(m, "DeleteComment", arg0, arg1)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteComment indicates an expected call of DeleteComment
func (mr *MockCmsServerMockRecorder) DeleteComment(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteComment", reflect.TypeOf((*MockCmsServer)(nil).DeleteComment), arg0, arg1)
}

// DeletePost mocks base method
func (m *MockCmsServer) DeletePost(arg0 context.Context, arg1 *proto.PostRequest) (*empty.Empty, error) {
	ret := m.ctrl.Call(m, "DeletePost", arg0, arg1)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePost indicates an expected call of DeletePost
func (mr *MockCmsServerMockRecorder) DeletePost(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePost", reflect.TypeOf((*MockCmsServer)(nil).DeletePost), arg0, arg1)
}

// DeleteUser mocks base method
func (m *MockCmsServer) DeleteUser(arg0 context.Context, arg1 *proto.UserRequest) (*empty.Empty, error) {
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUser indicates an expected call of DeleteUser
func (mr *MockCmsServerMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockCmsServer)(nil).DeleteUser), arg0, arg1)
}

// GetComment mocks base method
func (m *MockCmsServer) GetComment(arg0 context.Context, arg1 *proto.CommentRequest) (*proto.Comment, error) {
	ret := m.ctrl.Call(m, "GetComment", arg0, arg1)
	ret0, _ := ret[0].(*proto.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComment indicates an expected call of GetComment
func (mr *MockCmsServerMockRecorder) GetComment(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComment", reflect.TypeOf((*MockCmsServer)(nil).GetComment), arg0, arg1)
}

// GetComments mocks base method
func (m *MockCmsServer) GetComments(arg0 *empty.Empty, arg1 proto.Cms_GetCommentsServer) error {
	ret := m.ctrl.Call(m, "GetComments", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetComments indicates an expected call of GetComments
func (mr *MockCmsServerMockRecorder) GetComments(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComments", reflect.TypeOf((*MockCmsServer)(nil).GetComments), arg0, arg1)
}

// GetPost mocks base method
func (m *MockCmsServer) GetPost(arg0 context.Context, arg1 *proto.PostRequest) (*proto.Post, error) {
	ret := m.ctrl.Call(m, "GetPost", arg0, arg1)
	ret0, _ := ret[0].(*proto.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPost indicates an expected call of GetPost
func (mr *MockCmsServerMockRecorder) GetPost(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPost", reflect.TypeOf((*MockCmsServer)(nil).GetPost), arg0, arg1)
}

// GetPostBySlug mocks base method
func (m *MockCmsServer) GetPostBySlug(arg0 context.Context, arg1 *proto.PostBySlugRequest) (*proto.Post, error) {
	ret := m.ctrl.Call(m, "GetPostBySlug", arg0, arg1)
	ret0, _ := ret[0].(*proto.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostBySlug indicates an expected call of GetPostBySlug
func (mr *MockCmsServerMockRecorder) GetPostBySlug(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostBySlug", reflect.TypeOf((*MockCmsServer)(nil).GetPostBySlug), arg0, arg1)
}

// GetPostComments mocks base method
func (m *MockCmsServer) GetPostComments(arg0 *proto.PostRequest, arg1 proto.Cms_GetPostCommentsServer) error {
	ret := m.ctrl.Call(m, "GetPostComments", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetPostComments indicates an expected call of GetPostComments
func (mr *MockCmsServerMockRecorder) GetPostComments(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostComments", reflect.TypeOf((*MockCmsServer)(nil).GetPostComments), arg0, arg1)
}

// GetPosts mocks base method
func (m *MockCmsServer) GetPosts(arg0 *proto.GetPostsOptions, arg1 proto.Cms_GetPostsServer) error {
	ret := m.ctrl.Call(m, "GetPosts", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetPosts indicates an expected call of GetPosts
func (mr *MockCmsServerMockRecorder) GetPosts(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPosts", reflect.TypeOf((*MockCmsServer)(nil).GetPosts), arg0, arg1)
}

// GetUser mocks base method
func (m *MockCmsServer) GetUser(arg0 context.Context, arg1 *proto.UserRequest) (*proto.User, error) {
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(*proto.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser
func (mr *MockCmsServerMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockCmsServer)(nil).GetUser), arg0, arg1)
}

// GetUserComments mocks base method
func (m *MockCmsServer) GetUserComments(arg0 *proto.UserRequest, arg1 proto.Cms_GetUserCommentsServer) error {
	ret := m.ctrl.Call(m, "GetUserComments", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetUserComments indicates an expected call of GetUserComments
func (mr *MockCmsServerMockRecorder) GetUserComments(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserComments", reflect.TypeOf((*MockCmsServer)(nil).GetUserComments), arg0, arg1)
}

// IsSetup mocks base method
func (m *MockCmsServer) IsSetup(arg0 context.Context, arg1 *empty.Empty) (*wrappers.BoolValue, error) {
	ret := m.ctrl.Call(m, "IsSetup", arg0, arg1)
	ret0, _ := ret[0].(*wrappers.BoolValue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsSetup indicates an expected call of IsSetup
func (mr *MockCmsServerMockRecorder) IsSetup(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSetup", reflect.TypeOf((*MockCmsServer)(nil).IsSetup), arg0, arg1)
}

// Setup mocks base method
func (m *MockCmsServer) Setup(arg0 context.Context, arg1 *proto.CreateUserRequest) (*proto.UserRequest, error) {
	ret := m.ctrl.Call(m, "Setup", arg0, arg1)
	ret0, _ := ret[0].(*proto.UserRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Setup indicates an expected call of Setup
func (mr *MockCmsServerMockRecorder) Setup(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Setup", reflect.TypeOf((*MockCmsServer)(nil).Setup), arg0, arg1)
}

// UpdateComment mocks base method
func (m *MockCmsServer) UpdateComment(arg0 context.Context, arg1 *proto.UpdateCommentRequest) (*empty.Empty, error) {
	ret := m.ctrl.Call(m, "UpdateComment", arg0, arg1)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateComment indicates an expected call of UpdateComment
func (mr *MockCmsServerMockRecorder) UpdateComment(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateComment", reflect.TypeOf((*MockCmsServer)(nil).UpdateComment), arg0, arg1)
}

// UpdatePost mocks base method
func (m *MockCmsServer) UpdatePost(arg0 context.Context, arg1 *proto.UpdatePostRequest) (*empty.Empty, error) {
	ret := m.ctrl.Call(m, "UpdatePost", arg0, arg1)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePost indicates an expected call of UpdatePost
func (mr *MockCmsServerMockRecorder) UpdatePost(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePost", reflect.TypeOf((*MockCmsServer)(nil).UpdatePost), arg0, arg1)
}
