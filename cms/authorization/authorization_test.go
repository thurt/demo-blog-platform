package authorization

import (
	"errors"
	"testing"

	"context"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes/empty"
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

func TestGetUnpublishedPosts(t *testing.T) {
	r := &empty.Empty{}
	mockStream := mock_proto.NewMockCms_GetUnpublishedPostsServer()

	t.Run("requires permission", func(t *testing.T) {
		mock, a := setup(t)
		ctx := reqContext.NewFromUser(context.Background(), &pb.User{Role: pb.UserRole_UNKNOWN})
		mockStream.SetContext(ctx)

		mock.EXPECT().GetUnpublishedPosts(gomock.Any(), gomock.Any())

		err := a.GetUnpublishedPosts(r, mockStream)
		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires Admin Role has permission", func(t *testing.T) {
		mock, a := setup(t)
		ctx := reqContext.NewFromUser(context.Background(), &pb.User{Role: pb.UserRole_ADMIN})
		mockStream.SetContext(ctx)

		mock.EXPECT().GetUnpublishedPosts(gomock.Any(), gomock.Any())

		err := a.GetUnpublishedPosts(r, mockStream)
		if err != nil {
			t.Error("unexpected error:", err)
		}
	})
}

func TestGetUnpublishedPost(t *testing.T) {
	t.Run("must return error when CmsServer.GetUnpublishedPost returns error", func(t *testing.T) {
		mock, a := setup(t)
		ctx := reqContext.NewFromUser(context.Background(), &pb.User{Role: pb.UserRole_ADMIN})
		r := &pb.PostRequest{}

		mock.EXPECT().GetUnpublishedPost(ctx, r).Return(nil, errors.New(""))

		_, err := a.GetUnpublishedPost(ctx, r)
		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires permission when returned Post is unpublished", func(t *testing.T) {
		mock, a := setup(t)
		ctx := reqContext.NewFromUser(context.Background(), &pb.User{Role: pb.UserRole_UNKNOWN})
		r := &pb.PostRequest{}

		mock.EXPECT().GetUnpublishedPost(ctx, r)

		_, err := a.GetUnpublishedPost(ctx, r)
		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires Admin Role has permission when returned Post is unpublished", func(t *testing.T) {
		mock, a := setup(t)
		ctx := reqContext.NewFromUser(context.Background(), &pb.User{Role: pb.UserRole_ADMIN})
		r := &pb.PostRequest{}

		mock.EXPECT().GetUnpublishedPost(ctx, r)

		_, err := a.GetUnpublishedPost(ctx, r)
		if err != nil {
			t.Error("unexpected error:", err)
		}
	})
}

func TestGetUnpublishedPostBySlug(t *testing.T) {
	t.Run("must return error when CmsServer.GetPostBySlug returns error", func(t *testing.T) {
		mock, a := setup(t)
		r := &pb.PostBySlugRequest{}

		mock.EXPECT().GetUnpublishedPostBySlug(gomock.Any(), r).Return(nil, errors.New(""))

		_, err := a.GetUnpublishedPostBySlug(context.Background(), r)
		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires permission when returned Post is unpublished", func(t *testing.T) {
		mock, a := setup(t)
		ctx := reqContext.NewFromUser(context.Background(), &pb.User{Role: pb.UserRole_UNKNOWN})
		r := &pb.PostBySlugRequest{}

		mock.EXPECT().GetUnpublishedPostBySlug(ctx, r)

		_, err := a.GetUnpublishedPostBySlug(ctx, r)
		if err == nil {
			t.Error("expected an error")
		}

	})
	t.Run("requires Admin Role has permission when returned Post is unpublished", func(t *testing.T) {
		mock, a := setup(t)
		ctx := reqContext.NewFromUser(context.Background(), &pb.User{Role: pb.UserRole_ADMIN})
		r := &pb.PostBySlugRequest{}

		mock.EXPECT().GetUnpublishedPostBySlug(ctx, r)

		_, err := a.GetUnpublishedPostBySlug(ctx, r)
		if err != nil {
			t.Error("unexpected error:", err)
		}
	})
}
