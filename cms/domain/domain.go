package domain

import (
	"golang.org/x/net/context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	pb "github.com/thurt/demo-blog-platform/cms/proto"
)

type Provider interface {
	CreatePost(context.Context, *pb.CreatePostRequest) (*pb.PostRequest, error)
	GetPost(context.Context, *pb.PostRequest) (*pb.Post, error)
	UpdatePost(context.Context, *pb.UpdatePostRequest) (*empty.Empty, error)
	DeletePost(context.Context, *pb.PostRequest) (*empty.Empty, error)
	GetPostComments(*pb.PostRequest, pb.Cms_GetPostCommentsServer) error
	GetPosts(*empty.Empty, pb.Cms_GetPostsServer) error
	PublishPost(context.Context, *pb.PostRequest) (*empty.Empty, error)
	UnPublishPost(context.Context, *pb.PostRequest) (*empty.Empty, error)
	CreateUser(context.Context, *pb.CreateUserWithRole) (*pb.UserRequest, error)
	GetUser(context.Context, *pb.UserRequest) (*pb.User, error)
	AdminExists(context.Context, *empty.Empty) (*wrappers.BoolValue, error)
	DeleteUser(context.Context, *pb.UserRequest) (*empty.Empty, error)
	GetUserComments(*pb.UserRequest, pb.Cms_GetUserCommentsServer) error
	CreateComment(context.Context, *pb.CreateCommentRequest) (*pb.CommentRequest, error)
	GetComment(context.Context, *pb.CommentRequest) (*pb.Comment, error)
	UpdateComment(context.Context, *pb.UpdateCommentRequest) (*empty.Empty, error)
	DeleteComment(context.Context, *pb.CommentRequest) (*empty.Empty, error)
	GetComments(*empty.Empty, pb.Cms_GetCommentsServer) error
	GetUserPassword(context.Context, *pb.UserRequest) (*pb.UserPassword, error)
}

type UseCases interface {
	CreatePost(context.Context, *pb.CreatePostRequest) (*pb.PostRequest, error)
	UpdatePost(context.Context, *pb.UpdatePostRequest) (*empty.Empty, error)
	CreateUser(context.Context, *pb.CreateUserRequest) (*pb.UserRequest, error)
	CreateComment(context.Context, *pb.CreateCommentRequest) (*pb.CommentRequest, error)
	AuthUser(context.Context, *pb.AuthUserRequest) (*pb.AccessToken, error)
}
