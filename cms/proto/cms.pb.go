// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cms.proto

/*
Package cms is a generated protocol buffer package.

CMS

CMS Service API provides access to CMS entities and supports CMS use-cases

It is generated from these files:
	cms.proto

It has these top-level messages:
	Post
	PostRequest
	CreatePostRequest
	UpdatePostRequest
	Comment
	CommentRequest
	CreateCommentRequest
	UpdateCommentRequest
	User
	UserRequest
	CreateUserRequest
*/
package cms

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/mwitkow/go-proto-validators"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Posts
type Post struct {
	Id         uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Title      string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Content    string `protobuf:"bytes,3,opt,name=content" json:"content,omitempty"`
	Created    string `protobuf:"bytes,4,opt,name=created" json:"created,omitempty"`
	LastEdited string `protobuf:"bytes,5,opt,name=last_edited,json=lastEdited" json:"last_edited,omitempty"`
	Published  string `protobuf:"bytes,6,opt,name=published" json:"published,omitempty"`
	Slug       string `protobuf:"bytes,7,opt,name=slug" json:"slug,omitempty"`
}

func (m *Post) Reset()                    { *m = Post{} }
func (m *Post) String() string            { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()               {}
func (*Post) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Post) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Post) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Post) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Post) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *Post) GetLastEdited() string {
	if m != nil {
		return m.LastEdited
	}
	return ""
}

func (m *Post) GetPublished() string {
	if m != nil {
		return m.Published
	}
	return ""
}

func (m *Post) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

type PostRequest struct {
	Id uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *PostRequest) Reset()                    { *m = PostRequest{} }
func (m *PostRequest) String() string            { return proto.CompactTextString(m) }
func (*PostRequest) ProtoMessage()               {}
func (*PostRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PostRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type CreatePostRequest struct {
	Title   string `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content" json:"content,omitempty"`
}

func (m *CreatePostRequest) Reset()                    { *m = CreatePostRequest{} }
func (m *CreatePostRequest) String() string            { return proto.CompactTextString(m) }
func (*CreatePostRequest) ProtoMessage()               {}
func (*CreatePostRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreatePostRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CreatePostRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type UpdatePostRequest struct {
	Id      uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=content" json:"content,omitempty"`
}

func (m *UpdatePostRequest) Reset()                    { *m = UpdatePostRequest{} }
func (m *UpdatePostRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdatePostRequest) ProtoMessage()               {}
func (*UpdatePostRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UpdatePostRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdatePostRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *UpdatePostRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

// Comments
type Comment struct {
	Id         uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Content    string `protobuf:"bytes,2,opt,name=content" json:"content,omitempty"`
	Created    string `protobuf:"bytes,3,opt,name=created" json:"created,omitempty"`
	LastEdited string `protobuf:"bytes,4,opt,name=last_edited,json=lastEdited" json:"last_edited,omitempty"`
	UserId     string `protobuf:"bytes,5,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	PostId     string `protobuf:"bytes,6,opt,name=post_id,json=postId" json:"post_id,omitempty"`
}

func (m *Comment) Reset()                    { *m = Comment{} }
func (m *Comment) String() string            { return proto.CompactTextString(m) }
func (*Comment) ProtoMessage()               {}
func (*Comment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Comment) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Comment) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Comment) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *Comment) GetLastEdited() string {
	if m != nil {
		return m.LastEdited
	}
	return ""
}

func (m *Comment) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Comment) GetPostId() string {
	if m != nil {
		return m.PostId
	}
	return ""
}

type CommentRequest struct {
	Id uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *CommentRequest) Reset()                    { *m = CommentRequest{} }
func (m *CommentRequest) String() string            { return proto.CompactTextString(m) }
func (*CommentRequest) ProtoMessage()               {}
func (*CommentRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CommentRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type CreateCommentRequest struct {
	Content string `protobuf:"bytes,1,opt,name=content" json:"content,omitempty"`
	UserId  string `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	PostId  string `protobuf:"bytes,3,opt,name=post_id,json=postId" json:"post_id,omitempty"`
}

func (m *CreateCommentRequest) Reset()                    { *m = CreateCommentRequest{} }
func (m *CreateCommentRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateCommentRequest) ProtoMessage()               {}
func (*CreateCommentRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CreateCommentRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *CreateCommentRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *CreateCommentRequest) GetPostId() string {
	if m != nil {
		return m.PostId
	}
	return ""
}

type UpdateCommentRequest struct {
	Id      uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content" json:"content,omitempty"`
}

func (m *UpdateCommentRequest) Reset()                    { *m = UpdateCommentRequest{} }
func (m *UpdateCommentRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateCommentRequest) ProtoMessage()               {}
func (*UpdateCommentRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *UpdateCommentRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateCommentRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

// Users
type User struct {
	Id         string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Email      string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	Created    string `protobuf:"bytes,3,opt,name=created" json:"created,omitempty"`
	LastActive string `protobuf:"bytes,4,opt,name=last_active,json=lastActive" json:"last_active,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *User) GetLastActive() string {
	if m != nil {
		return m.LastActive
	}
	return ""
}

type UserRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *UserRequest) Reset()                    { *m = UserRequest{} }
func (m *UserRequest) String() string            { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()               {}
func (*UserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *UserRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CreateUserRequest struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
}

func (m *CreateUserRequest) Reset()                    { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()               {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *CreateUserRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*Post)(nil), "cms.Post")
	proto.RegisterType((*PostRequest)(nil), "cms.PostRequest")
	proto.RegisterType((*CreatePostRequest)(nil), "cms.CreatePostRequest")
	proto.RegisterType((*UpdatePostRequest)(nil), "cms.UpdatePostRequest")
	proto.RegisterType((*Comment)(nil), "cms.Comment")
	proto.RegisterType((*CommentRequest)(nil), "cms.CommentRequest")
	proto.RegisterType((*CreateCommentRequest)(nil), "cms.CreateCommentRequest")
	proto.RegisterType((*UpdateCommentRequest)(nil), "cms.UpdateCommentRequest")
	proto.RegisterType((*User)(nil), "cms.User")
	proto.RegisterType((*UserRequest)(nil), "cms.UserRequest")
	proto.RegisterType((*CreateUserRequest)(nil), "cms.CreateUserRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Cms service

type CmsClient interface {
	// Post CRUD
	CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*PostRequest, error)
	GetPost(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*Post, error)
	UpdatePost(ctx context.Context, in *UpdatePostRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	DeletePost(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	// Post Use-Cases
	GetPostComments(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (Cms_GetPostCommentsClient, error)
	GetPosts(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (Cms_GetPostsClient, error)
	PublishPost(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	UnPublishPost(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	// User CRD
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserRequest, error)
	GetUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*User, error)
	DeleteUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	// User Use-Cases
	GetUserComments(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (Cms_GetUserCommentsClient, error)
	// Comment CRUD
	CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CommentRequest, error)
	GetComment(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*Comment, error)
	UpdateComment(ctx context.Context, in *UpdateCommentRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	DeleteComment(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	// Comment Use-Cases
	GetComments(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (Cms_GetCommentsClient, error)
}

type cmsClient struct {
	cc *grpc.ClientConn
}

func NewCmsClient(cc *grpc.ClientConn) CmsClient {
	return &cmsClient{cc}
}

func (c *cmsClient) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*PostRequest, error) {
	out := new(PostRequest)
	err := grpc.Invoke(ctx, "/cms.Cms/CreatePost", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsClient) GetPost(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := grpc.Invoke(ctx, "/cms.Cms/GetPost", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsClient) UpdatePost(ctx context.Context, in *UpdatePostRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/cms.Cms/UpdatePost", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsClient) DeletePost(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/cms.Cms/DeletePost", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsClient) GetPostComments(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (Cms_GetPostCommentsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Cms_serviceDesc.Streams[0], c.cc, "/cms.Cms/GetPostComments", opts...)
	if err != nil {
		return nil, err
	}
	x := &cmsGetPostCommentsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Cms_GetPostCommentsClient interface {
	Recv() (*Comment, error)
	grpc.ClientStream
}

type cmsGetPostCommentsClient struct {
	grpc.ClientStream
}

func (x *cmsGetPostCommentsClient) Recv() (*Comment, error) {
	m := new(Comment)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cmsClient) GetPosts(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (Cms_GetPostsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Cms_serviceDesc.Streams[1], c.cc, "/cms.Cms/GetPosts", opts...)
	if err != nil {
		return nil, err
	}
	x := &cmsGetPostsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Cms_GetPostsClient interface {
	Recv() (*Post, error)
	grpc.ClientStream
}

type cmsGetPostsClient struct {
	grpc.ClientStream
}

func (x *cmsGetPostsClient) Recv() (*Post, error) {
	m := new(Post)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cmsClient) PublishPost(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/cms.Cms/PublishPost", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsClient) UnPublishPost(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/cms.Cms/UnPublishPost", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserRequest, error) {
	out := new(UserRequest)
	err := grpc.Invoke(ctx, "/cms.Cms/CreateUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsClient) GetUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := grpc.Invoke(ctx, "/cms.Cms/GetUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsClient) DeleteUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/cms.Cms/DeleteUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsClient) GetUserComments(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (Cms_GetUserCommentsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Cms_serviceDesc.Streams[2], c.cc, "/cms.Cms/GetUserComments", opts...)
	if err != nil {
		return nil, err
	}
	x := &cmsGetUserCommentsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Cms_GetUserCommentsClient interface {
	Recv() (*Comment, error)
	grpc.ClientStream
}

type cmsGetUserCommentsClient struct {
	grpc.ClientStream
}

func (x *cmsGetUserCommentsClient) Recv() (*Comment, error) {
	m := new(Comment)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cmsClient) CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CommentRequest, error) {
	out := new(CommentRequest)
	err := grpc.Invoke(ctx, "/cms.Cms/CreateComment", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsClient) GetComment(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*Comment, error) {
	out := new(Comment)
	err := grpc.Invoke(ctx, "/cms.Cms/GetComment", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsClient) UpdateComment(ctx context.Context, in *UpdateCommentRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/cms.Cms/UpdateComment", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsClient) DeleteComment(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/cms.Cms/DeleteComment", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsClient) GetComments(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (Cms_GetCommentsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Cms_serviceDesc.Streams[3], c.cc, "/cms.Cms/GetComments", opts...)
	if err != nil {
		return nil, err
	}
	x := &cmsGetCommentsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Cms_GetCommentsClient interface {
	Recv() (*Comment, error)
	grpc.ClientStream
}

type cmsGetCommentsClient struct {
	grpc.ClientStream
}

func (x *cmsGetCommentsClient) Recv() (*Comment, error) {
	m := new(Comment)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Cms service

type CmsServer interface {
	// Post CRUD
	CreatePost(context.Context, *CreatePostRequest) (*PostRequest, error)
	GetPost(context.Context, *PostRequest) (*Post, error)
	UpdatePost(context.Context, *UpdatePostRequest) (*google_protobuf1.Empty, error)
	DeletePost(context.Context, *PostRequest) (*google_protobuf1.Empty, error)
	// Post Use-Cases
	GetPostComments(*PostRequest, Cms_GetPostCommentsServer) error
	GetPosts(*google_protobuf1.Empty, Cms_GetPostsServer) error
	PublishPost(context.Context, *PostRequest) (*google_protobuf1.Empty, error)
	UnPublishPost(context.Context, *PostRequest) (*google_protobuf1.Empty, error)
	// User CRD
	CreateUser(context.Context, *CreateUserRequest) (*UserRequest, error)
	GetUser(context.Context, *UserRequest) (*User, error)
	DeleteUser(context.Context, *UserRequest) (*google_protobuf1.Empty, error)
	// User Use-Cases
	GetUserComments(*UserRequest, Cms_GetUserCommentsServer) error
	// Comment CRUD
	CreateComment(context.Context, *CreateCommentRequest) (*CommentRequest, error)
	GetComment(context.Context, *CommentRequest) (*Comment, error)
	UpdateComment(context.Context, *UpdateCommentRequest) (*google_protobuf1.Empty, error)
	DeleteComment(context.Context, *CommentRequest) (*google_protobuf1.Empty, error)
	// Comment Use-Cases
	GetComments(*google_protobuf1.Empty, Cms_GetCommentsServer) error
}

func RegisterCmsServer(s *grpc.Server, srv CmsServer) {
	s.RegisterService(&_Cms_serviceDesc, srv)
}

func _Cms_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cms.Cms/CreatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServer).CreatePost(ctx, req.(*CreatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cms_GetPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServer).GetPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cms.Cms/GetPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServer).GetPost(ctx, req.(*PostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cms_UpdatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServer).UpdatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cms.Cms/UpdatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServer).UpdatePost(ctx, req.(*UpdatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cms_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cms.Cms/DeletePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServer).DeletePost(ctx, req.(*PostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cms_GetPostComments_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PostRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CmsServer).GetPostComments(m, &cmsGetPostCommentsServer{stream})
}

type Cms_GetPostCommentsServer interface {
	Send(*Comment) error
	grpc.ServerStream
}

type cmsGetPostCommentsServer struct {
	grpc.ServerStream
}

func (x *cmsGetPostCommentsServer) Send(m *Comment) error {
	return x.ServerStream.SendMsg(m)
}

func _Cms_GetPosts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(google_protobuf1.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CmsServer).GetPosts(m, &cmsGetPostsServer{stream})
}

type Cms_GetPostsServer interface {
	Send(*Post) error
	grpc.ServerStream
}

type cmsGetPostsServer struct {
	grpc.ServerStream
}

func (x *cmsGetPostsServer) Send(m *Post) error {
	return x.ServerStream.SendMsg(m)
}

func _Cms_PublishPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServer).PublishPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cms.Cms/PublishPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServer).PublishPost(ctx, req.(*PostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cms_UnPublishPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServer).UnPublishPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cms.Cms/UnPublishPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServer).UnPublishPost(ctx, req.(*PostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cms_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cms.Cms/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cms_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cms.Cms/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServer).GetUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cms_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cms.Cms/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServer).DeleteUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cms_GetUserComments_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UserRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CmsServer).GetUserComments(m, &cmsGetUserCommentsServer{stream})
}

type Cms_GetUserCommentsServer interface {
	Send(*Comment) error
	grpc.ServerStream
}

type cmsGetUserCommentsServer struct {
	grpc.ServerStream
}

func (x *cmsGetUserCommentsServer) Send(m *Comment) error {
	return x.ServerStream.SendMsg(m)
}

func _Cms_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cms.Cms/CreateComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServer).CreateComment(ctx, req.(*CreateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cms_GetComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServer).GetComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cms.Cms/GetComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServer).GetComment(ctx, req.(*CommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cms_UpdateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServer).UpdateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cms.Cms/UpdateComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServer).UpdateComment(ctx, req.(*UpdateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cms_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cms.Cms/DeleteComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServer).DeleteComment(ctx, req.(*CommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cms_GetComments_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(google_protobuf1.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CmsServer).GetComments(m, &cmsGetCommentsServer{stream})
}

type Cms_GetCommentsServer interface {
	Send(*Comment) error
	grpc.ServerStream
}

type cmsGetCommentsServer struct {
	grpc.ServerStream
}

func (x *cmsGetCommentsServer) Send(m *Comment) error {
	return x.ServerStream.SendMsg(m)
}

var _Cms_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cms.Cms",
	HandlerType: (*CmsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePost",
			Handler:    _Cms_CreatePost_Handler,
		},
		{
			MethodName: "GetPost",
			Handler:    _Cms_GetPost_Handler,
		},
		{
			MethodName: "UpdatePost",
			Handler:    _Cms_UpdatePost_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _Cms_DeletePost_Handler,
		},
		{
			MethodName: "PublishPost",
			Handler:    _Cms_PublishPost_Handler,
		},
		{
			MethodName: "UnPublishPost",
			Handler:    _Cms_UnPublishPost_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _Cms_CreateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Cms_GetUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _Cms_DeleteUser_Handler,
		},
		{
			MethodName: "CreateComment",
			Handler:    _Cms_CreateComment_Handler,
		},
		{
			MethodName: "GetComment",
			Handler:    _Cms_GetComment_Handler,
		},
		{
			MethodName: "UpdateComment",
			Handler:    _Cms_UpdateComment_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _Cms_DeleteComment_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetPostComments",
			Handler:       _Cms_GetPostComments_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetPosts",
			Handler:       _Cms_GetPosts_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetUserComments",
			Handler:       _Cms_GetUserComments_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetComments",
			Handler:       _Cms_GetComments_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cms.proto",
}

func init() { proto.RegisterFile("cms.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 908 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x55, 0xdd, 0x6e, 0xdc, 0x44,
	0x14, 0xc6, 0xbb, 0xdb, 0xdd, 0xec, 0x59, 0x76, 0x49, 0x26, 0xab, 0x65, 0x6b, 0x22, 0xad, 0x65,
	0x81, 0x14, 0x45, 0x64, 0x5d, 0x11, 0x89, 0x0b, 0xb8, 0x40, 0x4d, 0xa8, 0x42, 0x04, 0x42, 0x21,
	0x6a, 0x50, 0xd5, 0xaa, 0x54, 0xb3, 0xf6, 0xb0, 0x31, 0xd8, 0x9e, 0xc5, 0x33, 0x6e, 0x94, 0x56,
	0xbd, 0xe1, 0x15, 0x78, 0x01, 0x6e, 0xb9, 0xe6, 0x09, 0x78, 0x06, 0x1e, 0xa0, 0x52, 0xc5, 0x83,
	0xa0, 0xf9, 0xb1, 0x3d, 0xf6, 0xc6, 0x4d, 0x41, 0x48, 0xdc, 0x79, 0xce, 0xcf, 0x37, 0xdf, 0xf9,
	0xce, 0x99, 0x63, 0xe8, 0xfb, 0x31, 0x9b, 0xaf, 0x52, 0xca, 0x29, 0x6a, 0xfb, 0x31, 0xb3, 0x77,
	0x96, 0x94, 0x2e, 0x23, 0xe2, 0xe1, 0x55, 0xe8, 0xe1, 0x24, 0xa1, 0x1c, 0xf3, 0x90, 0x26, 0x3a,
	0xc4, 0x7e, 0x4f, 0x7b, 0xe5, 0x69, 0x91, 0x7d, 0xef, 0x91, 0x78, 0xc5, 0xaf, 0xb4, 0xf3, 0xe3,
	0x65, 0xc8, 0x2f, 0xb2, 0xc5, 0xdc, 0xa7, 0xb1, 0x17, 0x5f, 0x86, 0xfc, 0x47, 0x7a, 0xe9, 0x2d,
	0xe9, 0xbe, 0x74, 0xee, 0x3f, 0xc5, 0x51, 0x18, 0x60, 0x4e, 0x53, 0xe6, 0x15, 0x9f, 0x2a, 0xcf,
	0xfd, 0xdd, 0x82, 0xce, 0x29, 0x65, 0x1c, 0x8d, 0xa0, 0x15, 0x06, 0x53, 0xcb, 0xb1, 0x76, 0x87,
	0x67, 0xad, 0x30, 0x40, 0x63, 0xb8, 0xc5, 0x43, 0x1e, 0x91, 0x69, 0xcb, 0xb1, 0x76, 0xfb, 0x67,
	0xea, 0x80, 0xa6, 0xd0, 0xf3, 0x69, 0xc2, 0x49, 0xc2, 0xa7, 0x6d, 0x69, 0xcf, 0x8f, 0xd2, 0x93,
	0x12, 0xcc, 0x49, 0x30, 0xed, 0x68, 0x8f, 0x3a, 0xa2, 0x19, 0x0c, 0x22, 0xcc, 0xf8, 0x13, 0x12,
	0x84, 0xc2, 0x7b, 0x4b, 0x7a, 0x41, 0x98, 0xee, 0x49, 0x0b, 0xda, 0x81, 0xfe, 0x2a, 0x5b, 0x44,
	0x21, 0xbb, 0x20, 0xc1, 0xb4, 0x2b, 0xdd, 0xa5, 0x01, 0x21, 0xe8, 0xb0, 0x28, 0x5b, 0x4e, 0x7b,
	0xd2, 0x21, 0xbf, 0xdd, 0x0f, 0x60, 0x20, 0x48, 0x9f, 0x91, 0x9f, 0x32, 0xc2, 0x38, 0x9a, 0x94,
	0xdc, 0x0f, 0xbb, 0xaf, 0x5e, 0xce, 0x5a, 0x9b, 0x6f, 0x89, 0x1a, 0xdc, 0x23, 0xd8, 0x3a, 0x92,
	0x24, 0xcc, 0xe0, 0xa2, 0x30, 0xab, 0xa1, 0xb0, 0x56, 0xa5, 0x30, 0xf7, 0x11, 0x6c, 0x9d, 0xaf,
	0x82, 0x1a, 0x48, 0xc3, 0x8d, 0xff, 0x54, 0x35, 0xf7, 0x57, 0x0b, 0x7a, 0x47, 0x34, 0x8e, 0x85,
	0x82, 0xf5, 0x0e, 0x34, 0x52, 0x32, 0xb5, 0x6e, 0xbf, 0x56, 0xeb, 0xce, 0x9a, 0xd6, 0xef, 0x42,
	0x2f, 0x63, 0x24, 0x7d, 0x12, 0xe6, 0x8d, 0xe8, 0x8a, 0xe3, 0x89, 0x74, 0xac, 0x28, 0xe3, 0xc2,
	0xa1, 0x5a, 0xd0, 0x15, 0xc7, 0x93, 0xc0, 0xdd, 0x85, 0x91, 0x66, 0x78, 0x93, 0xdc, 0xcf, 0x60,
	0xac, 0xe4, 0xae, 0xc5, 0x3b, 0x65, 0x21, 0x52, 0x73, 0x95, 0xf4, 0xc0, 0x2a, 0x0b, 0x9a, 0x95,
	0xac, 0x5a, 0x95, 0x88, 0x9c, 0xdd, 0xac, 0x64, 0xd7, 0xae, 0x06, 0x68, 0x96, 0xa7, 0x30, 0x56,
	0x5d, 0x7a, 0x33, 0xae, 0x26, 0xa7, 0xd6, 0xb5, 0x9c, 0xdc, 0x25, 0x74, 0xce, 0x19, 0x49, 0x8d,
	0xb6, 0xf4, 0xf3, 0x16, 0x93, 0x18, 0x87, 0x51, 0xde, 0x62, 0x79, 0x78, 0x83, 0x96, 0x60, 0x9f,
	0x87, 0x4f, 0x89, 0xd9, 0x92, 0xbb, 0xd2, 0xe2, 0x5e, 0xc1, 0x40, 0x5c, 0x94, 0x33, 0xfe, 0xa1,
	0xbc, 0xef, 0xf0, 0xe1, 0xab, 0x97, 0xb3, 0x6f, 0x61, 0xf4, 0xdd, 0x23, 0xbc, 0xff, 0xec, 0xee,
	0xfe, 0xc3, 0xc7, 0xcf, 0x0f, 0x3e, 0x7c, 0xf1, 0xfe, 0xde, 0xa1, 0x08, 0x77, 0x4e, 0x02, 0x27,
	0xce, 0x18, 0x77, 0x16, 0xc4, 0xc1, 0xdc, 0x89, 0x08, 0x66, 0xdc, 0x39, 0x70, 0xfc, 0x0b, 0x9c,
	0x62, 0x9f, 0x93, 0x94, 0x39, 0x38, 0x09, 0x1c, 0x1f, 0x27, 0x0e, 0x4d, 0xa2, 0x2b, 0x27, 0x4c,
	0xfc, 0x28, 0x0b, 0x88, 0x13, 0x11, 0x2e, 0x9c, 0x0f, 0x2c, 0xd9, 0xb1, 0x3f, 0xac, 0xfc, 0x85,
	0xfc, 0x4f, 0x0c, 0xd0, 0x4e, 0x45, 0xcd, 0xa2, 0x0b, 0x5a, 0x55, 0x17, 0x36, 0x56, 0x98, 0xb1,
	0x4b, 0x9a, 0xd6, 0xfb, 0x5e, 0xd8, 0x3f, 0xfa, 0x0d, 0xa0, 0x7d, 0x14, 0x33, 0x74, 0x02, 0x50,
	0x3e, 0x76, 0x34, 0x99, 0x8b, 0xdd, 0xba, 0xf6, 0xfa, 0xed, 0x4d, 0x69, 0x37, 0x2c, 0xee, 0xd6,
	0xcf, 0x7f, 0xfe, 0xf5, 0x4b, 0x6b, 0xe0, 0x76, 0x3d, 0x31, 0x4a, 0xec, 0x13, 0x6b, 0x0f, 0x7d,
	0x0a, 0xbd, 0x63, 0xc2, 0x25, 0xce, 0x5a, 0xbc, 0xdd, 0x2f, 0x2c, 0xee, 0xb6, 0x4c, 0x1d, 0xa2,
	0x81, 0x4a, 0xf5, 0x9e, 0x87, 0xc1, 0x0b, 0x74, 0x1f, 0xa0, 0xdc, 0x17, 0x9a, 0xc7, 0xda, 0x02,
	0xb1, 0x27, 0x73, 0xb5, 0xcd, 0xe7, 0xf9, 0x36, 0x9f, 0xdf, 0x13, 0xdb, 0xdc, 0x9d, 0x48, 0xc8,
	0x4d, 0xdb, 0x84, 0x14, 0x94, 0xbe, 0x04, 0xf8, 0x9c, 0x44, 0x44, 0xa3, 0xae, 0xb3, 0x6a, 0xc2,
	0xd3, 0x14, 0xf7, 0x2a, 0x14, 0xbf, 0x81, 0x77, 0x74, 0x7d, 0xfa, 0xb5, 0xb0, 0x6b, 0x10, 0xdf,
	0x56, 0x0a, 0xaa, 0x00, 0x77, 0x47, 0xe2, 0x4c, 0xd0, 0xd8, 0xc0, 0xf1, 0x7c, 0x9d, 0x7d, 0xc7,
	0x42, 0x9f, 0xc1, 0x86, 0x86, 0x64, 0xa8, 0x81, 0x8b, 0xa9, 0xdc, 0x48, 0xc2, 0x6d, 0x20, 0x2d,
	0xfa, 0x1d, 0x0b, 0x7d, 0x05, 0x83, 0x53, 0xb5, 0xf3, 0xff, 0x5d, 0x85, 0x15, 0xc5, 0xd0, 0xd7,
	0x30, 0x3c, 0x4f, 0xfe, 0x43, 0xbc, 0x62, 0xb8, 0xe4, 0x4a, 0x30, 0x87, 0xcb, 0x78, 0x38, 0x7a,
	0xb8, 0x0c, 0x8b, 0x31, 0x5c, 0x62, 0x91, 0x19, 0xc3, 0x25, 0x71, 0xd6, 0xe2, 0xb5, 0x44, 0xc2,
	0x62, 0x0c, 0x97, 0x4c, 0x55, 0x3c, 0x8a, 0x31, 0x68, 0xc8, 0xbf, 0x79, 0x0c, 0x0c, 0x30, 0x35,
	0x06, 0x22, 0xbd, 0x36, 0x06, 0x26, 0x62, 0xd3, 0x18, 0x94, 0x38, 0xe6, 0x18, 0x9c, 0xc3, 0xb0,
	0xf2, 0x0b, 0x40, 0xb7, 0x0d, 0xa9, 0xaa, 0xab, 0xd9, 0xde, 0x36, 0x91, 0x73, 0xc1, 0xc6, 0xf2,
	0x82, 0x91, 0xdb, 0x2f, 0x50, 0x85, 0x66, 0xc7, 0x00, 0xc7, 0x24, 0x1f, 0x56, 0x74, 0x5d, 0x62,
	0x8d, 0xa7, 0x7e, 0x46, 0x68, 0x54, 0xc0, 0xa8, 0x92, 0x1f, 0xc3, 0xb0, 0xf2, 0x9b, 0xd0, 0xfc,
	0xae, 0xfb, 0x75, 0x34, 0x6a, 0x79, 0x5b, 0x62, 0x6f, 0xdb, 0x35, 0x6c, 0xc1, 0xf3, 0x3e, 0x0c,
	0x55, 0x7b, 0x5e, 0x4b, 0xf5, 0x86, 0xb7, 0xbf, 0x57, 0x27, 0xfd, 0x05, 0x0c, 0xca, 0xea, 0x9b,
	0x9f, 0x57, 0x55, 0x01, 0x3d, 0x79, 0xa8, 0x6f, 0xb4, 0x67, 0xd1, 0x95, 0x29, 0x07, 0x7f, 0x07,
	0x00, 0x00, 0xff, 0xff, 0xb2, 0x36, 0xa9, 0x55, 0x79, 0x0a, 0x00, 0x00,
}
