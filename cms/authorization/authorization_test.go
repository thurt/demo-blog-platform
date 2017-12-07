package authorization

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/thurt/demo-blog-platform/cms/mock_proto"
	pb "github.com/thurt/demo-blog-platform/cms/proto"
	"github.com/thurt/demo-blog-platform/cms/reqContext"
	"golang.org/x/net/context"
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

		mock.EXPECT().DeleteComment(ctx, r)

		_, err := a.DeleteComment(ctx, r)
		if err != nil {
			t.Error("unexpected error:", err)
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
}

func TestUnPublishPost(t *testing.T) {
	t.Run("requires permission", func(t *testing.T) {
		mock, a := setup(t)
		ctx := reqContext.NewFromUser(context.Background(), &pb.User{Role: pb.UserRole_UNKNOWN})
		r := &pb.PostRequest{}

		mock.EXPECT().UnPublishPost(ctx, r)

		_, err := a.UnPublishPost(ctx, r)
		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires Admin Role has permission", func(t *testing.T) {
		mock, a := setup(t)
		ctx := reqContext.NewFromUser(context.Background(), &pb.User{Role: pb.UserRole_ADMIN})
		r := &pb.PostRequest{}

		mock.EXPECT().UnPublishPost(ctx, r)

		_, err := a.UnPublishPost(ctx, r)
		if err != nil {
			t.Error("unexpected error:", err)
		}
	})
}

func TestPublishPost(t *testing.T) {
	t.Run("requires permission", func(t *testing.T) {
		mock, a := setup(t)
		ctx := reqContext.NewFromUser(context.Background(), &pb.User{Role: pb.UserRole_UNKNOWN})
		r := &pb.PostRequest{}

		mock.EXPECT().PublishPost(ctx, r)

		_, err := a.PublishPost(ctx, r)
		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires Admin Role has permission", func(t *testing.T) {
		mock, a := setup(t)
		ctx := reqContext.NewFromUser(context.Background(), &pb.User{Role: pb.UserRole_ADMIN})
		r := &pb.PostRequest{}

		mock.EXPECT().PublishPost(ctx, r)

		_, err := a.PublishPost(ctx, r)
		if err != nil {
			t.Error("unexpected error:", err)
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
