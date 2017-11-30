package authorization

import (
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/thurt/demo-blog-platform/cms/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var ErrPermissionDenied = status.Error(codes.PermissionDenied, codes.PermissionDenied.String())

type Role string

var (
	Admin = Role("admin")
	User  = Role("user")
)

func (r Role) String() string {
	return string(r)
}

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

func hasPermission(ctx context.Context, rolesAllowed ...Role) bool {
	md, ok := metadata.FromIncomingContext(ctx)

	if ok && md["role"] != nil && len(md["role"][0]) != 0 {
		cr := Role(md["role"][0])

		for _, r := range rolesAllowed {
			if r == cr {
				return true
			}
		}
	}

	return false
}

func (a *authorization) CreatePost(ctx context.Context, r *pb.CreatePostRequest) (*pb.PostRequest, error) {
	// requires Admin Role has permission
	if !hasPermission(ctx, Admin) {
		return nil, ErrPermissionDenied
	}

	return a.CmsServer.CreatePost(ctx, r)
}

func (a *authorization) UpdatePost(ctx context.Context, r *pb.UpdatePostRequest) (*empty.Empty, error) {
	// requires Admin Role has permission
	if !hasPermission(ctx, Admin) {
		return nil, ErrPermissionDenied
	}

	return a.CmsServer.UpdatePost(ctx, r)
}

func (a *authorization) DeletePost(ctx context.Context, r *pb.PostRequest) (*empty.Empty, error) {
	// requires Admin Role has permission
	if !hasPermission(ctx, Admin) {
		return nil, ErrPermissionDenied
	}
	return a.CmsServer.DeletePost(ctx, r)
}

func (a *authorization) PublishPost(ctx context.Context, r *pb.PostRequest) (*empty.Empty, error) {
	// requires Admin Role has permission
	if !hasPermission(ctx, Admin) {
		return nil, ErrPermissionDenied
	}
	return a.CmsServer.PublishPost(ctx, r)
}

func (a *authorization) UnPublishPost(ctx context.Context, r *pb.PostRequest) (*empty.Empty, error) {
	// requires Admin Role has permission
	if !hasPermission(ctx, Admin) {
		return nil, ErrPermissionDenied
	}
	return a.CmsServer.UnPublishPost(ctx, r)
}

func (a *authorization) DeleteUser(ctx context.Context, r *pb.UserRequest) (*empty.Empty, error) {
	// requires Admin Role has permission
	// requires User Role has permission
	if !hasPermission(ctx, Admin, User) {
		return nil, ErrPermissionDenied
	}
	return a.CmsServer.DeleteUser(ctx, r)
}

func (a *authorization) CreateComment(ctx context.Context, r *pb.CreateCommentRequest) (*pb.CommentRequest, error) {
	// requires Admin Role has permission
	// requires User Role has permission
	if !hasPermission(ctx, Admin, User) {
		return nil, ErrPermissionDenied
	}
	return a.CmsServer.CreateComment(ctx, r)
}

func (a *authorization) UpdateComment(ctx context.Context, r *pb.UpdateCommentRequest) (*empty.Empty, error) {
	// requires Admin Role has permission
	// requires User Role has permission
	if !hasPermission(ctx, Admin, User) {
		return nil, ErrPermissionDenied
	}
	return a.CmsServer.UpdateComment(ctx, r)
}

func (a *authorization) DeleteComment(ctx context.Context, r *pb.CommentRequest) (*empty.Empty, error) {
	// requires Admin Role has permission
	// requires User Role has permission
	if !hasPermission(ctx, Admin, User) {
		return nil, ErrPermissionDenied
	}
	return a.CmsServer.DeleteComment(ctx, r)
}
