package authorization

import (
	"errors"
	"testing"

	"context"

	"github.com/golang/mock/gomock"
	"github.com/thurt/demo-blog-platform/cms/mock_proto"
	pb "github.com/thurt/demo-blog-platform/cms/proto"
	"github.com/thurt/demo-blog-platform/cms/reqContext"
	"google.golang.org/grpc/status"
)

func setup(t *testing.T) (*mock_proto.MockCmsServer, Authorization) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mock := mock_proto.NewMockCmsServer(mockCtrl)
	a := New(mock)
	return mock, a
}

func TestDeleteComment(t *testing.T) {
	t.Run("requires permission", func(t *testing.T) {
		mock, a := setup(t)
		ctx := reqContext.NewFromUser(context.Background(), &pb.User{Role: pb.UserRole_UNKNOWN})
		r := &pb.CommentRequest{}

		mock.EXPECT().DeleteComment(ctx, r)

		_, err := a.DeleteComment(ctx, r)
		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires Admin Role has permission", func(t *testing.T) {
		mock, a := setup(t)
		ctx := reqContext.NewFromUser(context.Background(), &pb.User{Role: pb.UserRole_ADMIN})
		r := &pb.CommentRequest{}

		mock.EXPECT().DeleteComment(ctx, r)

		_, err := a.DeleteComment(ctx, r)
		if err != nil {
			t.Error("unexpected error:", err)
		}
	})
	t.Run("requires User Role has permission", func(t *testing.T) {
		mock, a := setup(t)
		ctx := reqContext.NewFromUser(context.Background(), &pb.User{Role: pb.UserRole_USER})
		r := &pb.CommentRequest{}

		mock.EXPECT().GetComment(ctx, r).Return(&pb.Comment{}, nil)
		mock.EXPECT().DeleteComment(ctx, r)

		_, err := a.DeleteComment(ctx, r)
		if err != nil {
			t.Error("unexpected error:", err)
		}
	})
	t.Run("User Role is not allowed to delete comments by other users", func(t *testing.T) {
		mock, a := setup(t)
		ctx := reqContext.NewFromUser(context.Background(), &pb.User{Role: pb.UserRole_USER, Id: "user"})
		r := &pb.CommentRequest{}

		mock.EXPECT().GetComment(ctx, r).Return(&pb.Comment{UserId: "not_user"}, nil)

		_, err := a.DeleteComment(ctx, r)
		if err == nil {
			t.Error("expected an error")
		}
	})
}

func TestUpdateComment(t *testing.T) {
	t.Run("requires permission", func(t *testing.T) {
		mock, a := setup(t)
		ctx := reqContext.NewFromUser(context.Background(), &pb.User{Role: pb.UserRole_UNKNOWN})
		r := &pb.UpdateCommentRequest{}

		mock.EXPECT().UpdateComment(ctx, r)

		_, err := a.UpdateComment(ctx, r)
		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires Admin Role has permission", func(t *testing.T) {
		mock, a := setup(t)
		ctx := reqContext.NewFromUser(context.Background(), &pb.User{Role: pb.UserRole_ADMIN})
		r := &pb.UpdateCommentRequest{}

		mock.EXPECT().UpdateComment(ctx, r)

		_, err := a.UpdateComment(ctx, r)
		if err != nil {
			t.Error("unexpected error:", err)
		}
	})
	t.Run("requires User Role has permission", func(t *testing.T) {
		mock, a := setup(t)
		ctx := reqContext.NewFromUser(context.Background(), &pb.User{Role: pb.UserRole_USER})
		r := &pb.UpdateCommentRequest{}

		mock.EXPECT().UpdateComment(ctx, r)

		_, err := a.UpdateComment(ctx, r)
		if err != nil {
			t.Error("unexpected error:", err)
		}
	})

}

func TestCreateComment(t *testing.T) {
	t.Run("requires permission", func(t *testing.T) {
		mock, a := setup(t)
		ctx := reqContext.NewFromUser(context.Background(), &pb.User{Role: pb.UserRole_UNKNOWN})
		r := &pb.CreateCommentRequest{}

		mock.EXPECT().CreateComment(ctx, r)

		_, err := a.CreateComment(ctx, r)
		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires Admin Role has permission", func(t *testing.T) {
		mock, a := setup(t)
		ctx := reqContext.NewFromUser(context.Background(), &pb.User{Role: pb.UserRole_ADMIN})
		r := &pb.CreateCommentRequest{}

		mock.EXPECT().CreateComment(ctx, r)

		_, err := a.CreateComment(ctx, r)
		if err != nil {
			t.Error("unexpected error:", err)
		}
	})
	t.Run("requires User Role has permission", func(t *testing.T) {
		mock, a := setup(t)
		ctx := reqContext.NewFromUser(context.Background(), &pb.User{Role: pb.UserRole_USER})
		r := &pb.CreateCommentRequest{}

		mock.EXPECT().CreateComment(ctx, r)

		_, err := a.CreateComment(ctx, r)
		if err != nil {
			t.Error("unexpected error:", err)
		}
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("requires permission", func(t *testing.T) {
		mock, a := setup(t)
		ctx := reqContext.NewFromUser(context.Background(), &pb.User{Role: pb.UserRole_UNKNOWN})
		r := &pb.UserRequest{}

		mock.EXPECT().DeleteUser(ctx, r)

		_, err := a.DeleteUser(ctx, r)
		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires Admin Role has permission", func(t *testing.T) {
		mock, a := setup(t)
		ctx := reqContext.NewFromUser(context.Background(), &pb.User{Role: pb.UserRole_ADMIN})
		r := &pb.UserRequest{}

		mock.EXPECT().DeleteUser(ctx, r)

		_, err := a.DeleteUser(ctx, r)
		if err != nil {
			t.Error("unexpected error:", err)
		}
	})
	t.Run("requires User Role has permission", func(t *testing.T) {
		mock, a := setup(t)
		ctx := reqContext.NewFromUser(context.Background(), &pb.User{Role: pb.UserRole_USER})
		r := &pb.UserRequest{}

		mock.EXPECT().DeleteUser(ctx, r)

		_, err := a.DeleteUser(ctx, r)
		if err != nil {
			t.Error("unexpected error:", err)
		}
	})
	t.Run("must answer with grpc error when User Role tries to delete other User", func(t *testing.T) {
		mock, a := setup(t)
		ctx := reqContext.NewFromUser(context.Background(), &pb.User{Id: "id", Role: pb.UserRole_USER})
		r := &pb.UserRequest{Id: "different_id"}

		mock.EXPECT().DeleteUser(ctx, r)

		_, err := a.DeleteUser(ctx, r)
		if err == nil {
			t.Error("must be an error")
		}
		_, ok := status.FromError(err)
		if !ok {
			t.Error("must be a grpc error")
		}
	})
}

func TestDeletePost(t *testing.T) {
	t.Run("requires permission", func(t *testing.T) {
		mock, a := setup(t)
		ctx := reqContext.NewFromUser(context.Background(), &pb.User{Role: pb.UserRole_UNKNOWN})
		r := &pb.PostRequest{}

		mock.EXPECT().DeletePost(ctx, r)

		_, err := a.DeletePost(ctx, r)
		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires Admin Role has permission", func(t *testing.T) {
		mock, a := setup(t)
		ctx := reqContext.NewFromUser(context.Background(), &pb.User{Role: pb.UserRole_ADMIN})
		r := &pb.PostRequest{}

		mock.EXPECT().DeletePost(ctx, r)

		_, err := a.DeletePost(ctx, r)
		if err != nil {
			t.Error("unexpected error:", err)
		}
	})
}

func TestUpdatePost(t *testing.T) {
	t.Run("requires permission", func(t *testing.T) {
		mock, a := setup(t)
		ctx := reqContext.NewFromUser(context.Background(), &pb.User{Role: pb.UserRole_UNKNOWN})
		r := &pb.UpdatePostRequest{}

		mock.EXPECT().UpdatePost(ctx, r)

		_, err := a.UpdatePost(ctx, r)
		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires Admin Role has permission", func(t *testing.T) {
		mock, a := setup(t)
		ctx := reqContext.NewFromUser(context.Background(), &pb.User{Role: pb.UserRole_ADMIN})
		r := &pb.UpdatePostRequest{}

		mock.EXPECT().UpdatePost(ctx, r)

		_, err := a.UpdatePost(ctx, r)
		if err != nil {
			t.Error("unexpected error:", err)
		}
	})
}

func TestCreatePost(t *testing.T) {
	t.Run("requires permission", func(t *testing.T) {
		mock, a := setup(t)
		ctx := reqContext.NewFromUser(context.Background(), &pb.User{Role: pb.UserRole_UNKNOWN})
		r := &pb.CreatePostRequest{}

		mock.EXPECT().CreatePost(ctx, r)

		_, err := a.CreatePost(ctx, r)
		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires Admin Role has permission", func(t *testing.T) {
		mock, a := setup(t)
		ctx := reqContext.NewFromUser(context.Background(), &pb.User{Role: pb.UserRole_ADMIN})
		r := &pb.CreatePostRequest{}

		mock.EXPECT().CreatePost(ctx, r)

		_, err := a.CreatePost(ctx, r)
		if err != nil {
			t.Error("unexpected error:", err)
		}
	})
}

func TestGetPosts(t *testing.T) {
	r := &pb.GetPostsOptions{IncludeUnPublished: true}
	mockStream := mock_proto.NewMockCms_GetPostsServer()

	t.Run("(options) requires permission when IncludeUnPublished is true", func(t *testing.T) {
		mock, a := setup(t)
		ctx := reqContext.NewFromUser(context.Background(), &pb.User{Role: pb.UserRole_UNKNOWN})
		mockStream.SetContext(ctx)

		mock.EXPECT().GetPosts(gomock.Any(), gomock.Any())

		err := a.GetPosts(r, mockStream)
		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("(options) requires Admin Role has permission when IncludeUnPublished is true", func(t *testing.T) {
		mock, a := setup(t)
		ctx := reqContext.NewFromUser(context.Background(), &pb.User{Role: pb.UserRole_ADMIN})
		mockStream.SetContext(ctx)

		mock.EXPECT().GetPosts(gomock.Any(), gomock.Any())

		err := a.GetPosts(r, mockStream)
		if err != nil {
			t.Error("unexpected error:", err)
		}
	})
}

func TestGetPost(t *testing.T) {
	t.Run("must return error when CmsServer.GetPost returns error", func(t *testing.T) {
		mock, a := setup(t)
		ctx := reqContext.NewFromUser(context.Background(), &pb.User{Role: pb.UserRole_ADMIN})
		r := &pb.PostRequest{}

		mock.EXPECT().GetPost(ctx, r).Return(nil, errors.New(""))

		_, err := a.GetPost(ctx, r)
		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires permission when returned Post is unpublished", func(t *testing.T) {
		mock, a := setup(t)
		ctx := reqContext.NewFromUser(context.Background(), &pb.User{Role: pb.UserRole_UNKNOWN})
		r := &pb.PostRequest{}

		mock.EXPECT().GetPost(ctx, r)

		_, err := a.GetPost(ctx, r)
		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires Admin Role has permission when returned Post is unpublished", func(t *testing.T) {
		mock, a := setup(t)
		ctx := reqContext.NewFromUser(context.Background(), &pb.User{Role: pb.UserRole_ADMIN})
		r := &pb.PostRequest{}

		mock.EXPECT().GetPost(ctx, r)

		_, err := a.GetPost(ctx, r)
		if err != nil {
			t.Error("unexpected error:", err)
		}
	})
}

func TestGetPostBySlug(t *testing.T) {
	t.Run("must return error when CmsServer.GetPostBySlug returns error", func(t *testing.T) {
		mock, a := setup(t)
		r := &pb.PostBySlugRequest{}

		mock.EXPECT().GetPostBySlug(gomock.Any(), r).Return(nil, errors.New(""))

		_, err := a.GetPostBySlug(context.Background(), r)
		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires permission when returned Post is unpublished", func(t *testing.T) {
		mock, a := setup(t)
		ctx := reqContext.NewFromUser(context.Background(), &pb.User{Role: pb.UserRole_UNKNOWN})
		r := &pb.PostBySlugRequest{}

		mock.EXPECT().GetPostBySlug(ctx, r)

		_, err := a.GetPostBySlug(ctx, r)
		if err == nil {
			t.Error("expected an error")
		}

	})
	t.Run("requires Admin Role has permission when returned Post is unpublished", func(t *testing.T) {
		mock, a := setup(t)
		ctx := reqContext.NewFromUser(context.Background(), &pb.User{Role: pb.UserRole_ADMIN})
		r := &pb.PostBySlugRequest{}

		mock.EXPECT().GetPostBySlug(ctx, r)

		_, err := a.GetPostBySlug(ctx, r)
		if err != nil {
			t.Error("unexpected error:", err)
		}
	})
}
