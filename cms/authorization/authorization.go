package authorization

import (
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/thurt/demo-blog-platform/cms/proto"
	"github.com/thurt/demo-blog-platform/cms/reqContext"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrPermissionDenied = status.Error(codes.PermissionDenied, codes.PermissionDenied.String())

type authorization struct {
	pb.CmsServer
}

type Authorization interface {
	CreatePost(context.Context, *pb.CreatePostRequest) (*pb.PostRequest, error)
	UpdatePost(context.Context, *pb.UpdatePostRequest) (*empty.Empty, error)
	DeletePost(context.Context, *pb.PostRequest) (*empty.Empty, error)
	DeleteUser(context.Context, *pb.UserRequest) (*empty.Empty, error)
	CreateComment(context.Context, *pb.CreateCommentRequest) (*pb.CommentRequest, error)
	UpdateComment(context.Context, *pb.UpdateCommentRequest) (*empty.Empty, error)
	DeleteComment(context.Context, *pb.CommentRequest) (*empty.Empty, error)
	GetPosts(*pb.GetPostsOptions, pb.Cms_GetPostsServer) error
	GetPost(context.Context, *pb.PostRequest) (*pb.Post, error)
	GetPostBySlug(context.Context, *pb.PostBySlugRequest) (*pb.Post, error)
}

func New(server pb.CmsServer) pb.CmsServer {
	a := &authorization{server}
	return a
}

func (a *authorization) CreatePost(ctx context.Context, r *pb.CreatePostRequest) (*pb.PostRequest, error) {
	// requires Admin Role has permission
	if !reqContext.HasPermission(ctx, pb.UserRole_ADMIN) {
		return nil, ErrPermissionDenied
	}

	return a.CmsServer.CreatePost(ctx, r)
}

func (a *authorization) UpdatePost(ctx context.Context, r *pb.UpdatePostRequest) (*empty.Empty, error) {
	// requires Admin Role has permission
	if !reqContext.HasPermission(ctx, pb.UserRole_ADMIN) {
		return nil, ErrPermissionDenied
	}

	return a.CmsServer.UpdatePost(ctx, r)
}

func (a *authorization) DeletePost(ctx context.Context, r *pb.PostRequest) (*empty.Empty, error) {
	// requires Admin Role has permission
	if !reqContext.HasPermission(ctx, pb.UserRole_ADMIN) {
		return nil, ErrPermissionDenied
	}
	return a.CmsServer.DeletePost(ctx, r)
}

func (a *authorization) DeleteUser(ctx context.Context, r *pb.UserRequest) (*empty.Empty, error) {
	// requires Admin Role has permission
	// requires User Role has permission
	if !reqContext.HasPermission(ctx, pb.UserRole_ADMIN, pb.UserRole_USER) {
		return nil, ErrPermissionDenied
	}

	u, _ := reqContext.GetUser(ctx)
	if u.GetRole() == pb.UserRole_USER && r.GetId() != u.GetId() {
		return nil, status.Errorf(codes.PermissionDenied, "Role %q is not allowed to delete other users", u.GetRole(), r.GetId())
	}

	return a.CmsServer.DeleteUser(ctx, r)
}

func (a *authorization) CreateComment(ctx context.Context, r *pb.CreateCommentRequest) (*pb.CommentRequest, error) {
	// requires Admin Role has permission
	// requires User Role has permission
	if !reqContext.HasPermission(ctx, pb.UserRole_ADMIN, pb.UserRole_USER) {
		return nil, ErrPermissionDenied
	}
	return a.CmsServer.CreateComment(ctx, r)
}

func (a *authorization) UpdateComment(ctx context.Context, r *pb.UpdateCommentRequest) (*empty.Empty, error) {
	// requires Admin Role has permission
	// requires User Role has permission
	if !reqContext.HasPermission(ctx, pb.UserRole_ADMIN, pb.UserRole_USER) {
		return nil, ErrPermissionDenied
	}
	return a.CmsServer.UpdateComment(ctx, r)
}

func (a *authorization) DeleteComment(ctx context.Context, r *pb.CommentRequest) (*empty.Empty, error) {
	// requires Admin Role has permission
	// requires User Role has permission
	if !reqContext.HasPermission(ctx, pb.UserRole_ADMIN, pb.UserRole_USER) {
		return nil, ErrPermissionDenied
	}
	return a.CmsServer.DeleteComment(ctx, r)
}

func (a *authorization) GetPosts(r *pb.GetPostsOptions, stream pb.Cms_GetPostsServer) error {
	if r.GetIncludeUnPublished() == true {
		// requires Admin Role has permission
		if !reqContext.HasPermission(stream.Context(), pb.UserRole_ADMIN) {
			return ErrPermissionDenied
		}
	}
	return a.CmsServer.GetPosts(r, stream)
}

func (a *authorization) GetPost(ctx context.Context, r *pb.PostRequest) (*pb.Post, error) {
	post, err := a.CmsServer.GetPost(ctx, r)
	if err != nil {
		return nil, err
	}
	if post.GetPublished() == false && !reqContext.HasPermission(ctx, pb.UserRole_ADMIN) {
		return nil, ErrPermissionDenied
	}
	return post, nil
}

func (a *authorization) GetPostBySlug(ctx context.Context, r *pb.PostBySlugRequest) (*pb.Post, error) {
	post, err := a.CmsServer.GetPostBySlug(ctx, r)
	if err != nil {
		return nil, err
	}
	if post.GetPublished() == false && !reqContext.HasPermission(ctx, pb.UserRole_ADMIN) {
		return nil, ErrPermissionDenied
	}
	return post, nil
}
