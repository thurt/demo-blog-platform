package authorization

import (
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/thurt/demo-blog-platform/cms/proto"
	"golang.org/x/net/context"
)

type authorization struct {
	pb.CmsServer
}

type Authorization interface {
	CreatePost(context.Context, *pb.CreatePostRequest) (*pb.PostRequest, error)
	UpdatePost(context.Context, *pb.UpdatePostRequest) (*empty.Empty, error)
	DeletePost(context.Context, *pb.PostRequest) (*empty.Empty, error)
	PublishPost(context.Context, *pb.PostRequest) (*empty.Empty, error)
	UnPublishPost(context.Context, *pb.PostRequest) (*empty.Empty, error)
	DeleteUser(context.Context, *pb.UserRequest) (*empty.Empty, error)
	CreateComment(context.Context, *pb.CreateCommentRequest) (*pb.CommentRequest, error)
	UpdateComment(context.Context, *pb.UpdateCommentRequest) (*empty.Empty, error)
	DeleteComment(context.Context, *pb.CommentRequest) (*empty.Empty, error)
}

func New(server pb.CmsServer) pb.CmsServer {
	a := &authorization{server}
	return a
}
