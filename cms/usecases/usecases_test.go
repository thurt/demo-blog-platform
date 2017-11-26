package usecases

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/thurt/demo-blog-platform/cms/mock_proto"
	pb "github.com/thurt/demo-blog-platform/cms/proto"
	"golang.org/x/net/context"
)

var ctx context.Context = context.Background()

func setup(t *testing.T) (*mock_proto.MockCmsServer, UseCases) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mock := mock_proto.NewMockCmsServer(mockCtrl)
	uc := New(mock)
	return mock, uc
}

func TestCreatePost(t *testing.T) {
	t.Run("creates a Slug from the provided Title", func(t *testing.T) {
		mock, uc := setup(t)

		r := &pb.CreatePostRequest{Title: "Hello World!"}

		mock.EXPECT().CreatePost(ctx, &pb.CreatePostRequest{Slug: "hello-world", Title: "Hello World!"})

		_, _ = uc.CreatePost(ctx, r)
	})

}

func TestUpdatePost(t *testing.T) {
	t.Run("creates a Slug from the provided Title", func(t *testing.T) {
		mock, uc := setup(t)

		r := &pb.UpdatePostRequest{Title: "Hello World!"}

		mock.EXPECT().UpdatePost(ctx, &pb.UpdatePostRequest{Slug: "hello-world", Title: "Hello World!"})

		_, _ = uc.UpdatePost(ctx, r)
	})
}

func TestCreateUser(t *testing.T) {
	t.Run("hashes the password", func(t *testing.T) {
		mock, uc := setup(t)

		r := &pb.CreateUserRequest{Password: "password"}

		mock.EXPECT().CreateUser(ctx, gomock.Not(&pb.CreateUserRequest{Password: "password"}))

		_, _ = uc.CreateUser(ctx, r)
	})
}
