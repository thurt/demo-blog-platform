package domain

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	pb "github.com/thurt/demo-blog-platform/cms/proto"
)

type Provider interface {
	pb.CmsInternalServer
	CreatePost(context.Context, *pb.CreatePostWithSlug) (*pb.PostRequest, error)
	GetPost(context.Context, *pb.PostRequest) (*pb.Post, error)
	GetPostBySlug(context.Context, *pb.PostBySlugRequest) (*pb.Post, error)
	GetUnpublishedPost(context.Context, *pb.PostRequest) (*pb.Post, error)
	GetUnpublishedPostBySlug(context.Context, *pb.PostBySlugRequest) (*pb.Post, error)
	UpdatePost(context.Context, *pb.UpdatePostWithSlug) (*empty.Empty, error)
	UpdateUnpublishedPost(context.Context, *pb.UpdatePostWithSlug) (*empty.Empty, error)
	DeletePost(context.Context, *pb.PostRequest) (*empty.Empty, error)
	GetPostComments(*pb.PostRequest, pb.Cms_GetPostCommentsServer) error
	GetPosts(*empty.Empty, pb.Cms_GetPostsServer) error
	GetUnpublishedPosts(*empty.Empty, pb.Cms_GetPostsServer) error
	GetUser(context.Context, *pb.UserRequest) (*pb.User, error)
	AdminExists(context.Context, *empty.Empty) (*wrappers.BoolValue, error)
	DeleteUser(context.Context, *pb.UserRequest) (*empty.Empty, error)
	GetUserComments(*pb.UserRequest, pb.Cms_GetUserCommentsServer) error
	CreateComment(context.Context, *pb.CreateCommentRequest) (*pb.CommentRequest, error)
	GetComment(context.Context, *pb.CommentRequest) (*pb.Comment, error)
	UpdateComment(context.Context, *pb.UpdateCommentRequest) (*empty.Empty, error)
	DeleteComment(context.Context, *pb.CommentRequest) (*empty.Empty, error)
	GetComments(*empty.Empty, pb.Cms_GetCommentsServer) error
}
