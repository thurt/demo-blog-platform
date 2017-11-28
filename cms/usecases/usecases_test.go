package usecases

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/thurt/demo-blog-platform/cms/mock_proto"
	pb "github.com/thurt/demo-blog-platform/cms/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

		r := &pb.CreateUserRequest{Id: "id", Password: "password"}

		mock.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(nil, status.Error(codes.NotFound, ""))
		mock.EXPECT().CreateUser(gomock.Any(), gomock.Not(&pb.CreateUserRequest{Password: "password"}))

		_, err := uc.CreateUser(ctx, r)
		if err != nil {
			t.Error("unexpected error:", err.Error())
		}
	})
	t.Run("returns an error when user already exists", func(t *testing.T) {
		mock, uc := setup(t)

		r := &pb.CreateUserRequest{}

		mock.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(&pb.User{}, nil)

		_, err := uc.CreateUser(ctx, r)
		if err == nil {
			t.Error("expected an error")
		}
	})
}
