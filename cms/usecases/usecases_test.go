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
	t.Run("requires a Slug to be created from the Title and added to the request", func(t *testing.T) {
		mock, uc := setup(t)

		r := &pb.CreatePostRequest{Title: "Hello World!"}

		mock.EXPECT().CreatePost(ctx, &pb.CreatePostRequest{Slug: "hello-world", Title: "Hello World!"})

		_, _ = uc.CreatePost(ctx, r)
	})

}

func TestUpdatePost(t *testing.T) {
	t.Run("requires a Slug to be created from the Title and added to the request", func(t *testing.T) {
		mock, uc := setup(t)

		r := &pb.UpdatePostRequest{Title: "Hello World!"}

		mock.EXPECT().UpdatePost(ctx, &pb.UpdatePostRequest{Slug: "hello-world", Title: "Hello World!"})

		_, _ = uc.UpdatePost(ctx, r)
	})
}

func TestCreateUser(t *testing.T) {
	t.Run("requires that user id does not exist", func(t *testing.T) {
		mock, uc := setup(t)

		r := &pb.CreateUserRequest{}

		mock.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(&pb.User{}, nil)

		_, err := uc.CreateUser(ctx, r)
		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires that password is hashed", func(t *testing.T) {
		mock, uc := setup(t)

		r := &pb.CreateUserRequest{Password: "password"}

		mock.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(nil, status.Error(codes.NotFound, ""))
		mock.EXPECT().CreateUser(gomock.Any(), gomock.Not(&pb.CreateUserRequest{Password: "password"}))

		_, err := uc.CreateUser(ctx, r)
		if err != nil {
			t.Error("unexpected error:", err.Error())
		}
	})
}

func TestCreateComment(t *testing.T) {
	t.Run("requires a valid user id", func(t *testing.T) {
		mock, uc := setup(t)

		r := &pb.CreateCommentRequest{UserId: "id"}

		mock.EXPECT().GetUser(gomock.Any(), &pb.UserRequest{"id"}).Return(nil, status.Error(codes.NotFound, ""))

		_, err := uc.CreateComment(ctx, r)

		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires a valid post id", func(t *testing.T) {
		mock, uc := setup(t)

		r := &pb.CreateCommentRequest{PostId: 0}

		mock.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(&pb.User{}, nil)
		mock.EXPECT().GetPost(gomock.Any(), &pb.PostRequest{0}).Return(nil, status.Error(codes.NotFound, ""))

		_, err := uc.CreateComment(ctx, r)
		if err == nil {
			t.Error("expected an error")
		}
	})
}
