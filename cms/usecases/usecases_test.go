package usecases

import (
	"encoding/hex"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	mock_proto_cacher "github.com/thurt/demo-blog-platform/cms/cacher/mock_proto"
	"github.com/thurt/demo-blog-platform/cms/mock_domain"
	"github.com/thurt/demo-blog-platform/cms/mock_proto"
	pb "github.com/thurt/demo-blog-platform/cms/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ctx context.Context = context.Background()

type TestFixture struct {
	Provider *mock_domain.MockProvider
	Auth     *mock_proto.MockCmsAuthServer
	Hasher   *mock_proto.MockHasherServer
	Emailer  *mock_proto.MockEmailerServer
	Cacher   *mock_proto_cacher.MockCacherServer
}

func newTestFixture(t *testing.T) (*TestFixture, *useCases) {
	// create NewMockCmsServer
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mock := mock_domain.NewMockProvider(mockCtrl)

	// create NewMockCmsAuthServer
	mockCtrlAuth := gomock.NewController(t)
	defer mockCtrlAuth.Finish()
	mockAuth := mock_proto.NewMockCmsAuthServer(mockCtrlAuth)

	// create NewMockHasher
	mockCtrlHasher := gomock.NewController(t)
	defer mockCtrlHasher.Finish()
	mockHasher := mock_proto.NewMockHasherServer(mockCtrlHasher)

	// create NewMockEmailer
	mockCtrlEmailer := gomock.NewController(t)
	defer mockCtrlEmailer.Finish()
	mockEmailer := mock_proto.NewMockEmailerServer(mockCtrlEmailer)

	// create NewMockCacher
	mockCtrlCacher := gomock.NewController(t)
	defer mockCtrlCacher.Finish()
	mockCacher := mock_proto_cacher.NewMockCacherServer(mockCtrlCacher)

	uc := &useCases{mock, mockAuth, mockHasher, mockEmailer, mockCacher}
	tf := &TestFixture{mock, mockAuth, mockHasher, mockEmailer, mockCacher}
	return tf, uc
}

func TestCreatePost(t *testing.T) {
	t.Run("requires a Slug to be created from the Title and added to the request", func(t *testing.T) {
		tf, uc := newTestFixture(t)

		r := &pb.CreatePostRequest{Title: "Hello World!"}

		tf.Provider.EXPECT().CreatePost(ctx, &pb.CreatePostWithSlug{Slug: "hello-world", Post: r})

		_, _ = uc.CreatePost(ctx, r)
	})
}

func TestUpdatePost(t *testing.T) {
	t.Run("requires a Slug to be created from the Title and added to the request", func(t *testing.T) {
		tf, uc := newTestFixture(t)

		r := &pb.UpdatePostRequest{Title: "Hello World!"}

		tf.Provider.EXPECT().UpdatePost(ctx, &pb.UpdatePostWithSlug{Slug: "hello-world", Post: r})

		_, _ = uc.UpdatePost(ctx, r)
	})
	t.Run("requires that published Posts contain a title", func(t *testing.T) {
		_, uc := newTestFixture(t)
		r := &pb.UpdatePostRequest{Title: "", Published: true}

		_, err := uc.UpdatePost(ctx, r)

		if err == nil {
			t.Error("expected an error")
		}
	})
}

func TestCreateComment(t *testing.T) {
	t.Run("requires a valid user id", func(t *testing.T) {
		tf, uc := newTestFixture(t)

		r := &pb.CreateCommentRequest{UserId: "id"}

		tf.Provider.EXPECT().GetUser(gomock.Any(), &pb.UserRequest{"id"}).Return(nil, status.Error(codes.NotFound, ""))

		_, err := uc.CreateComment(ctx, r)

		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires a valid post id", func(t *testing.T) {
		tf, uc := newTestFixture(t)

		r := &pb.CreateCommentRequest{PostId: 0}

		tf.Provider.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(&pb.User{}, nil)
		tf.Provider.EXPECT().GetPost(gomock.Any(), &pb.PostRequest{0}).Return(nil, status.Error(codes.NotFound, ""))

		_, err := uc.CreateComment(ctx, r)
		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("Comment cannot be created for a Post that is not published", func(t *testing.T) {
		tf, uc := newTestFixture(t)

		r := &pb.CreateCommentRequest{}

		tf.Provider.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(&pb.User{}, nil)
		tf.Provider.EXPECT().GetPost(gomock.Any(), gomock.Any()).Return(&pb.Post{Id: 1, Published: false}, nil)

		_, err := uc.CreateComment(ctx, r)
		if err == nil {
			t.Error("expected an error")
		}
	})
}

func TestAuthUser(t *testing.T) {
	t.Run("must answer with a grpc error when given a non-existant user id", func(t *testing.T) {
		tf, uc := newTestFixture(t)

		r := &pb.AuthUserRequest{Id: "id", Password: "password"}

		tf.Provider.EXPECT().GetUser(gomock.Any(), &pb.UserRequest{Id: r.GetId()}).Return(nil, errors.New(""))

		_, err := uc.AuthUser(ctx, r)

		if err == nil {
			t.Error("expected an error")
		}
		_, ok := status.FromError(err)
		if !ok {
			t.Error("must answer with a grpc error")
		}

	})
	t.Run("must answer with a grpc error when given an invalid password", func(t *testing.T) {
		tf, uc := newTestFixture(t)

		r := &pb.AuthUserRequest{Id: "id", Password: "wrong_password"}

		tf.Provider.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(&pb.User{Id: r.GetId()}, nil)
		tf.Provider.EXPECT().GetUserPassword(gomock.Any(), &pb.UserRequest{Id: r.GetId()}).Return(&pb.UserPassword{"hashed_password"}, nil)
		tf.Hasher.EXPECT().Validate(gomock.Any(), &pb.StrAndHash{"wrong_password", "hashed_password"}).Return(nil, errors.New(""))

		_, err := uc.AuthUser(ctx, r)

		if err == nil {
			t.Error("expected an error")
		}
		_, ok := status.FromError(err)
		if !ok {
			t.Error("must answer with a grpc error")
		}
	})
	t.Run("must answer with a grpc error when error occurs trying to get user password", func(t *testing.T) {
		tf, uc := newTestFixture(t)

		r := &pb.AuthUserRequest{Id: "id", Password: "right_password"}

		tf.Provider.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(&pb.User{Id: r.GetId()}, nil)
		tf.Provider.EXPECT().GetUserPassword(gomock.Any(), gomock.Any()).Return(nil, errors.New(""))

		_, err := uc.AuthUser(ctx, r)

		if err == nil {
			t.Error("expected an error")
		}
		_, ok := status.FromError(err)
		if !ok {
			t.Error("must answer with a grpc error")
		}
	})
	t.Run("must answer with a grpc error when error occurs trying to activate new token for user", func(t *testing.T) {
		tf, uc := newTestFixture(t)

		r := &pb.AuthUserRequest{Id: "id", Password: "right_password"}

		tf.Provider.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(&pb.User{Id: r.GetId()}, nil)
		tf.Provider.EXPECT().GetUserPassword(gomock.Any(), gomock.Any()).Return(&pb.UserPassword{"hashed_password"}, nil)
		tf.Hasher.EXPECT().Validate(gomock.Any(), &pb.StrAndHash{"right_password", "hashed_password"}).Return(&empty.Empty{}, nil)
		tf.Auth.EXPECT().ActivateNewTokenForUser(gomock.Any(), gomock.Any()).Return(nil, errors.New(""))

		_, err := uc.AuthUser(ctx, r)

		if err == nil {
			t.Error("expected an error")
		}
		_, ok := status.FromError(err)
		if !ok {
			t.Error("must answer with a grpc error")
		}
	})
	t.Run("must answer with a grpc error when error occurs trying to update user last active", func(t *testing.T) {
		tf, uc := newTestFixture(t)

		r := &pb.AuthUserRequest{Id: "id", Password: "right_password"}

		tf.Provider.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(&pb.User{Id: r.GetId()}, nil)
		tf.Provider.EXPECT().GetUserPassword(gomock.Any(), gomock.Any()).Return(&pb.UserPassword{"hashed_password"}, nil)
		tf.Hasher.EXPECT().Validate(gomock.Any(), &pb.StrAndHash{"right_password", "hashed_password"}).Return(&empty.Empty{}, nil)
		tf.Auth.EXPECT().ActivateNewTokenForUser(gomock.Any(), gomock.Any()).Return(&pb.AccessToken{}, nil)
		tf.Provider.EXPECT().UpdateUserLastActive(gomock.Any(), gomock.Any()).Return(nil, errors.New(""))

		_, err := uc.AuthUser(ctx, r)

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
		tf, uc := newTestFixture(t)

		r := &pb.UserRequest{}

		tf.Provider.EXPECT().GetUser(gomock.Any(), r).Return(nil, errors.New(""))

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
		tf, uc := newTestFixture(t)

		r := &pb.UserRequest{}

		tf.Provider.EXPECT().GetUser(gomock.Any(), r).Return(&pb.User{}, nil)

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
		tf, uc := newTestFixture(t)

		r := &pb.CommentRequest{}

		tf.Provider.EXPECT().GetComment(gomock.Any(), r).Return(&pb.Comment{}, nil)

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
		tf, uc := newTestFixture(t)

		r := &empty.Empty{}

		tf.Provider.EXPECT().AdminExists(gomock.Any(), r).Return(nil, errors.New(""))

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
		tf, uc := newTestFixture(t)

		// return true (admin exists) to test whether function answers properly
		tf.Provider.EXPECT().AdminExists(gomock.Any(), gomock.Any()).Return(&wrappers.BoolValue{true}, nil)

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
		tf, uc := newTestFixture(t)

		r := &pb.CreateUserRequest{}

		tf.Provider.EXPECT().AdminExists(gomock.Any(), gomock.Any()).Return(&wrappers.BoolValue{false}, nil)
		tf.Hasher.EXPECT().Hash(gomock.Any(), gomock.Any()).Return(&wrappers.StringValue{}, nil)
		tf.Provider.EXPECT().CreateNewUser(gomock.Any(), &pb.CreateUserWithRole{Role: pb.UserRole_ADMIN, User: r}).Return(nil, errors.New(""))

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
		tf, uc := newTestFixture(t)

		r := &pb.CreateUserRequest{Password: "password"}
		// because the implementation directly mutates the request instance, i need to retain a copy of the request with the original values in order to make an expectation that the password value has actually changed
		r_orig := &pb.CreateUserRequest{Password: r.GetPassword()}

		tf.Provider.EXPECT().AdminExists(gomock.Any(), gomock.Any()).Return(&wrappers.BoolValue{false}, nil)
		tf.Hasher.EXPECT().Hash(gomock.Any(), gomock.Any()).Return(&wrappers.StringValue{"hashed_password"}, nil)
		tf.Provider.EXPECT().CreateNewUser(gomock.Any(), gomock.Not(&pb.CreateUserWithRole{Role: pb.UserRole_ADMIN, User: r_orig}))

		_, err := uc.Setup(ctx, r)
		if err != nil {
			t.Error("unexpected error:", err.Error())
		}
	})
}

func TestGetPosts(t *testing.T) {
	r := &pb.GetPostsOptions{}
	mockStreamOut := mock_proto.NewMockCms_GetPostsServer()

	t.Run("must answer with a grpc error when receiving an error", func(t *testing.T) {
		tf, uc := newTestFixture(t)

		tf.Provider.EXPECT().GetPosts(gomock.Any(), gomock.Any()).Return(errors.New(""))

		err := uc.GetPosts(r, mockStreamOut)
		if err == nil {
			t.Error("must anwser with an error")
		}
		_, ok := status.FromError(err)
		if !ok {
			t.Error("must answer with a grpc error")
		}
	})
}

func TestGetPost(t *testing.T) {
	t.Run("must answer with a grpc error when receiving an error", func(t *testing.T) {
		tf, uc := newTestFixture(t)

		r := &pb.PostRequest{}

		tf.Provider.EXPECT().GetPost(gomock.Any(), r).Return(nil, errors.New(""))

		_, err := uc.GetPost(ctx, r)

		if err == nil {
			t.Error("must anwser with an error")
		}
		_, ok := status.FromError(err)
		if !ok {
			t.Error("must answer with a grpc error")
		}
	})
	t.Run("must answer with a grpc error when receiving a zero-value Post", func(t *testing.T) {
		tf, uc := newTestFixture(t)

		r := &pb.PostRequest{}

		tf.Provider.EXPECT().GetPost(gomock.Any(), r).Return(&pb.Post{}, nil)

		_, err := uc.GetPost(ctx, r)
		if err == nil {
			t.Error("expected an error")
		}
		_, ok := status.FromError(err)
		if !ok {
			t.Error("expected a grpc error")
		}
	})
}

func TestGetPostBySlug(t *testing.T) {
	t.Run("must answer with a grpc error when receiving an error", func(t *testing.T) {
		tf, uc := newTestFixture(t)

		r := &pb.PostBySlugRequest{}

		tf.Provider.EXPECT().GetPostBySlug(gomock.Any(), r).Return(nil, errors.New(""))

		_, err := uc.GetPostBySlug(ctx, r)

		if err == nil {
			t.Error("must anwser with an error")
		}
		_, ok := status.FromError(err)
		if !ok {
			t.Error("must answer with a grpc error")
		}
	})
	t.Run("must answer with a grpc error when receiving a zero-value Post", func(t *testing.T) {
		tf, uc := newTestFixture(t)

		r := &pb.PostBySlugRequest{}

		tf.Provider.EXPECT().GetPostBySlug(gomock.Any(), r).Return(&pb.Post{}, nil)

		_, err := uc.GetPostBySlug(ctx, r)
		if err == nil {
			t.Error("expected an error")
		}
		_, ok := status.FromError(err)
		if !ok {
			t.Error("expected a grpc error")
		}
	})
}

func TestRegisterNewUser(t *testing.T) {
	t.Run("must answer with a grpc error when receiving an error when getting user", func(t *testing.T) {
		tf, uc := newTestFixture(t)
		r := &pb.CreateUserRequest{Id: "id"}

		tf.Provider.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(nil, errors.New(""))

		_, err := uc.RegisterNewUser(ctx, r)
		if err == nil {
			t.Error("expected an error")
		}
		_, ok := status.FromError(err)
		if !ok {
			t.Error("must answer with a grpc error")
		}
	})
	t.Run("must answer with a grpc error when receiving a user id that already exists", func(t *testing.T) {
		tf, uc := newTestFixture(t)

		r := &pb.CreateUserRequest{Id: "id"}

		tf.Emailer.EXPECT().Send(gomock.Any(), gomock.Any()).Return(&empty.Empty{}, nil)
		tf.Provider.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(&pb.User{Id: "id"}, nil)

		_, err := uc.RegisterNewUser(ctx, r)
		if err == nil {
			t.Error("expected an error")
		}
		_, ok := status.FromError(err)
		if !ok {
			t.Error("must answer with a grpc error")
		}
	})
	t.Run("must answer with a grpc error when receiving an error when sending email", func(t *testing.T) {
		tf, uc := newTestFixture(t)

		tf.Provider.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(&pb.User{}, nil)
		tf.Hasher.EXPECT().Hash(gomock.Any(), gomock.Any()).Return(&wrappers.StringValue{"hashed_password"}, nil)
		tf.Cacher.EXPECT().Set(gomock.Any(), gomock.Any()).Return(&empty.Empty{}, nil)
		tf.Emailer.EXPECT().Send(gomock.Any(), gomock.Any()).Return(nil, errors.New(""))

		r := &pb.CreateUserRequest{}

		_, err := uc.RegisterNewUser(ctx, r)
		if err == nil {
			t.Error("expected an error")
		}
		_, ok := status.FromError(err)
		if !ok {
			t.Error("must answer with a grpc error")
		}
	})
}

func TestVerifyNewUser(t *testing.T) {
	stubIn := &pb.VerifyNewUserRequest{}
	stubIn.Token = "abc"
	t.Run("must answer with a grpc error when receiving an error when getting user", func(t *testing.T) {
		tf, uc := newTestFixture(t)

		tf.Cacher.EXPECT().Get(gomock.Any(), gomock.Any()).Return(&wrappers.StringValue{""}, nil)
		tf.Provider.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(nil, errors.New(""))

		_, err := uc.VerifyNewUser(context.Background(), stubIn)
		if err == nil {
			t.Error("expected an error")
		}
		_, ok := status.FromError(err)
		if !ok {
			t.Error("must answer with a grpc error")
		}
	})
	t.Run("must answer with a grpc error when receiving a user id that already exists", func(t *testing.T) {
		tf, uc := newTestFixture(t)

		tf.Cacher.EXPECT().Get(gomock.Any(), gomock.Any()).Return(&wrappers.StringValue{""}, nil)
		tf.Provider.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(&pb.User{Id: "id"}, nil)

		_, err := uc.VerifyNewUser(context.Background(), stubIn)
		if err == nil {
			t.Error("expected an error")
		}
		_, ok := status.FromError(err)
		if !ok {
			t.Error("must answer with a grpc error")
		}
	})
	t.Run("must answer with a grpc error when receiving an error when sending email", func(t *testing.T) {
		tf, uc := newTestFixture(t)

		tf.Cacher.EXPECT().Get(gomock.Any(), gomock.Any()).Return(&wrappers.StringValue{""}, nil)
		tf.Provider.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(&pb.User{}, nil)
		tf.Provider.EXPECT().CreateNewUser(gomock.Any(), gomock.Any()).Return(&pb.UserRequest{}, nil)
		tf.Emailer.EXPECT().Send(gomock.Any(), gomock.Any()).Return(nil, errors.New(""))

		_, err := uc.VerifyNewUser(context.Background(), stubIn)
		if err == nil {
			t.Error("expected an error")
		}
		_, ok := status.FromError(err)
		if !ok {
			t.Error("must answer with a grpc error")
		}
	})
}

func TestLogout(t *testing.T) {
	stubIn := &pb.AccessToken{AccessToken: "0987654321"}
	t.Run("must answer with a grpc error when receiving an error when deactivating token", func(t *testing.T) {
		tf, uc := newTestFixture(t)

		tf.Auth.EXPECT().DeactivateToken(gomock.Any(), &wrappers.StringValue{stubIn.GetAccessToken()}).Return(nil, errors.New(""))

		_, err := uc.Logout(context.Background(), stubIn)
		if err == nil {
			t.Error("expected an error")
		}
		_, ok := status.FromError(err)
		if !ok {
			t.Error("must answer with a grpc error")
		}
	})
	t.Run("must return successfully without error in normal circumstances", func(t *testing.T) {
		tf, uc := newTestFixture(t)

		tf.Auth.EXPECT().DeactivateToken(gomock.Any(), &wrappers.StringValue{stubIn.GetAccessToken()}).Return(&empty.Empty{}, nil)

		_, err := uc.Logout(context.Background(), stubIn)
		if err != nil {
			t.Error("unexpected error", err.Error())
		}
	})
}

func TestGenRandHexValue(t *testing.T) {
	bytes := uint(2)
	t.Run("must return a valid hex value", func(t *testing.T) {
		s := genRandHexValue(bytes)
		_, err := hex.DecodeString(s)
		if err != nil {
			t.Error("unexpected error", err.Error())
		}
	})
	t.Run("must return a value of specified length", func(t *testing.T) {
		v := genRandHexValue(bytes)
		// note: there are 2 hexadecimal units per byte
		if len(v) != int(bytes*2) {
			t.Errorf("expected length %d, got length %d", bytes*2, len(v))
		}
	})
}
