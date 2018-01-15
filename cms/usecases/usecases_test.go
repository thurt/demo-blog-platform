package usecases

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/thurt/demo-blog-platform/cms/mock_domain"
	"github.com/thurt/demo-blog-platform/cms/mock_proto"
	"github.com/thurt/demo-blog-platform/cms/password"
	pb "github.com/thurt/demo-blog-platform/cms/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ctx context.Context = context.Background()

func setup(t *testing.T) (*mock_domain.MockProvider, *mock_proto.MockCmsAuthServer, *useCases) {
	// create NewMockCmsServer
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mock := mock_domain.NewMockProvider(mockCtrl)

	// create NewMockCmsAuthServer
	mockCtrlAuth := gomock.NewController(t)
	defer mockCtrlAuth.Finish()
	mockAuth := mock_proto.NewMockCmsAuthServer(mockCtrlAuth)

	uc := New(mock, mockAuth)
	return mock, mockAuth, uc
}

func TestCreatePost(t *testing.T) {
	t.Run("requires a Slug to be created from the Title and added to the request", func(t *testing.T) {
		mock, _, uc := setup(t)

		r := &pb.CreatePostRequest{Title: "Hello World!"}

		mock.EXPECT().CreatePost(ctx, &pb.CreatePostRequest{Slug: "hello-world", Title: "Hello World!"})

		_, _ = uc.CreatePost(ctx, r)
	})

}

func TestUpdatePost(t *testing.T) {
	t.Run("requires a Slug to be created from the Title and added to the request", func(t *testing.T) {
		mock, _, uc := setup(t)

		r := &pb.UpdatePostRequest{Title: "Hello World!"}

		mock.EXPECT().UpdatePost(ctx, &pb.UpdatePostRequest{Slug: "hello-world", Title: "Hello World!"})

		_, _ = uc.UpdatePost(ctx, r)
	})
}

func TestCreateUser(t *testing.T) {
	t.Run("requires that user id does not exist", func(t *testing.T) {
		mock, _, uc := setup(t)

		r := &pb.CreateUserRequest{Id: "id"}

		mock.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(&pb.User{Id: "id"}, nil)

		_, err := uc.CreateUser(ctx, r)
		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires that password is hashed", func(t *testing.T) {
		mock, _, uc := setup(t)

		r := &pb.CreateUserRequest{Id: "id", Password: "password"}

		mock.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(&pb.User{}, nil)
		mock.EXPECT().CreateUser(gomock.Any(), gomock.Not(&pb.CreateUserRequest{Password: "password"}))

		_, err := uc.CreateUser(ctx, r)
		if err != nil {
			t.Error("unexpected error:", err.Error())
		}
	})
}

func TestCreateComment(t *testing.T) {
	t.Run("requires a valid user id", func(t *testing.T) {
		mock, _, uc := setup(t)

		r := &pb.CreateCommentRequest{UserId: "id"}

		mock.EXPECT().GetUser(gomock.Any(), &pb.UserRequest{"id"}).Return(nil, status.Error(codes.NotFound, ""))

		_, err := uc.CreateComment(ctx, r)

		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires a valid post id", func(t *testing.T) {
		mock, _, uc := setup(t)

		r := &pb.CreateCommentRequest{PostId: 0}

		mock.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(&pb.User{}, nil)
		mock.EXPECT().GetPost(gomock.Any(), &pb.PostRequest{0}).Return(nil, status.Error(codes.NotFound, ""))

		_, err := uc.CreateComment(ctx, r)
		if err == nil {
			t.Error("expected an error")
		}
	})
}

func TestAuthUser(t *testing.T) {
	t.Run("requires a valid user id", func(t *testing.T) {
		mock, _, uc := setup(t)

		r := &pb.AuthUserRequest{Id: "id", Password: "password"}

		mock.EXPECT().GetUser(gomock.Any(), &pb.UserRequest{Id: r.GetId()}).Return(nil, status.Error(codes.NotFound, ""))

		_, err := uc.AuthUser(ctx, r)

		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("must answer with a grpc error when given an invalid password", func(t *testing.T) {
		mock, _, uc := setup(t)

		r := &pb.AuthUserRequest{Id: "id", Password: "wrong_password"}

		// run my implementation of hashing in order to create mock's stub
		stubbedHash, err := password.Hash("right_password")
		if err != nil {
			t.Error("unexpected error during stub preparation")
		}

		mock.EXPECT().GetUser(gomock.Any(), gomock.Any())
		mock.EXPECT().GetUserPassword(gomock.Any(), &pb.UserRequest{Id: r.GetId()}).Return(&pb.UserPassword{stubbedHash}, nil)

		_, err = uc.AuthUser(ctx, r)

		if err == nil {
			t.Error("expected an error")
		}
		_, ok := status.FromError(err)
		if !ok {
			t.Error("must answer with a grpc error")
		}
	})
}

func TestGetUser(t *testing.T) {
	t.Run("must answer with a grpc error when receiving an error", func(t *testing.T) {
		mock, _, uc := setup(t)

		r := &pb.UserRequest{}

		mock.EXPECT().GetUser(gomock.Any(), r).Return(nil, errors.New(""))

		_, err := uc.GetUser(ctx, r)

		if err == nil {
			t.Error("must anwser with an error")
		}
		_, ok := status.FromError(err)
		if !ok {
			t.Error("must answer with a grpc error")
		}
	})
	t.Run("must answer with a grpc error when receiving a zero-value User", func(t *testing.T) {
		mock, _, uc := setup(t)

		r := &pb.UserRequest{}

		mock.EXPECT().GetUser(gomock.Any(), r).Return(&pb.User{}, nil)

		_, err := uc.GetUser(ctx, r)
		if err == nil {
			t.Error("expected an error")
		}
		_, ok := status.FromError(err)
		if !ok {
			t.Error("expected a grpc error")
		}
	})
}

func TestGetComment(t *testing.T) {
	t.Run("must answer with a grpc error when receiving a zero-value Comment", func(t *testing.T) {
		mock, _, uc := setup(t)

		r := &pb.CommentRequest{}

		mock.EXPECT().GetComment(gomock.Any(), r).Return(&pb.Comment{}, nil)

		_, err := uc.GetComment(ctx, r)
		if err == nil {
			t.Error("expected an error")
		}
		_, ok := status.FromError(err)
		if !ok {
			t.Error("expected a grpc error")
		}
	})
}

func TestIsSetup(t *testing.T) {
	t.Run("must answer with a grpc error when receiving an error", func(t *testing.T) {
		mock, _, uc := setup(t)

		r := &empty.Empty{}

		mock.EXPECT().AdminExists(gomock.Any(), r).Return(nil, errors.New(""))

		_, err := uc.IsSetup(ctx, r)

		if err == nil {
			t.Error("must anwser with an error")
		}
		_, ok := status.FromError(err)
		if !ok {
			t.Error("must answer with a grpc error")
		}
	})
}

func TestSetup(t *testing.T) {
	r := &pb.CreateUserRequest{}
	t.Run("must answer with a grpc error if Setup condition (admin does not exist) is not satisfied", func(t *testing.T) {
		mock, _, uc := setup(t)

		// return true (admin exists) to test whether function answers properly
		mock.EXPECT().AdminExists(gomock.Any(), gomock.Any()).Return(&wrappers.BoolValue{true}, nil)

		_, err := uc.Setup(ctx, r)
		if err == nil {
			t.Error("must anwser with an error")
		}
		_, ok := status.FromError(err)
		if !ok {
			t.Error("must answer with a grpc error")
		}
	})
	t.Run("must answer with a grpc error when receiving an error", func(t *testing.T) {
		mock, _, uc := setup(t)

		r := &pb.CreateUserRequest{}

		mock.EXPECT().AdminExists(gomock.Any(), gomock.Any()).Return(&wrappers.BoolValue{false}, nil)
		mock.EXPECT().CreateUser(gomock.Any(), &pb.CreateUserWithRole{Role: pb.UserRole_ADMIN, User: r}).Return(nil, errors.New(""))

		_, err := uc.Setup(ctx, r)

		if err == nil {
			t.Error("must anwser with an error")
		}
		_, ok := status.FromError(err)
		if !ok {
			t.Error("must answer with a grpc error")
		}
	})
	t.Run("requires that password is hashed", func(t *testing.T) {
		mock, _, uc := setup(t)

		r := &pb.CreateUserRequest{Id: "id", Password: "password"}

		mock.EXPECT().AdminExists(gomock.Any(), gomock.Any()).Return(&wrappers.BoolValue{false}, nil)
		mock.EXPECT().CreateUser(gomock.Any(), gomock.Not(&pb.CreateUserWithRole{Role: pb.UserRole_ADMIN, User: r}))

		_, err := uc.Setup(ctx, r)
		if err != nil {
			t.Error("unexpected error:", err.Error())
		}
	})
}
