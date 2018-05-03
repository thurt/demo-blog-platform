package authentication

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"context"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/thurt/demo-blog-platform/cms/cacher/mock_proto"
	pb_cacher "github.com/thurt/demo-blog-platform/cms/cacher/proto"
	pb "github.com/thurt/demo-blog-platform/cms/proto"
	"google.golang.org/grpc/metadata"
)

func setup(t *testing.T) *mock_proto.MockCacherServer {
	// create NewMockConn
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mock := mock_proto.NewMockCacherServer(mockCtrl)
	return mock
}

func TestActivateNewTokenForUser(t *testing.T) {
	stubIn := &pb.User{}
	tokenExpiry := 10 * time.Second
	t.Run("requires that a new token is added", func(t *testing.T) {
		mock := setup(t)
		aP, _ := New(mock, tokenExpiry)

		mock.EXPECT().Set(gomock.Any(), gomock.Any())

		_, err := aP.ActivateNewTokenForUser(context.Background(), stubIn)
		if err != nil {
			t.Error("unexpected error")
		}
	})
}

func TestActivateNewTokenForCreateUserWithRole(t *testing.T) {
	stubIn := &pb.CreateUserWithRole{}
	tokenExpiry := 10 * time.Second
	t.Run("requires that a new token is added", func(t *testing.T) {
		mock := setup(t)
		aP, _ := New(mock, tokenExpiry)

		mock.EXPECT().Set(gomock.Any(), gomock.Any())

		_, err := aP.ActivateNewTokenForCreateUserWithRole(context.Background(), stubIn)
		if err != nil {
			t.Error("unexpected error", err.Error())
		}
	})
}

func TestDeactivateToken(t *testing.T) {
	stubIn := &wrappers.StringValue{"0987654321"}
	t.Run("must return without error under normal circumstances", func(t *testing.T) {
		mock := setup(t)
		aP, _ := New(mock, 10*time.Second)

		mock.EXPECT().Delete(gomock.Any(), &pb_cacher.DeleteRequest{stubIn.GetValue()}).Return(&empty.Empty{}, nil)

		_, err := aP.DeactivateToken(context.Background(), stubIn)
		if err != nil {
			t.Error("unexpected an error", err.Error())
		}
	})
	t.Run("must return a grpc error when memcached returns error", func(t *testing.T) {
		mock := setup(t)
		aP, _ := New(mock, 10*time.Second)

		mock.EXPECT().Delete(gomock.Any(), &pb_cacher.DeleteRequest{stubIn.GetValue()}).Return(nil, errors.New(""))

		_, err := aP.DeactivateToken(context.Background(), stubIn)
		if err == nil {
			t.Error("expected an error")
		}
	})
}

func TestAuthFunc(t *testing.T) {
	t.Run("requires well-formed grpc metadata-based context", func(t *testing.T) {
		mock := setup(t)
		_, aF := New(mock, 10*time.Second)

		nonMetadataCtx := context.WithValue(context.Background(), "authorization", []string{"Bearer 0987654321"})

		_, err := aF(nonMetadataCtx)
		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires a no-op when no authorization is provided", func(t *testing.T) {
		mock := setup(t)
		_, aF := New(mock, 10*time.Second)

		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs())

		newCtx, err := aF(ctx)
		if err != nil {
			t.Error("unexpected error:", err)
		}

		if !reflect.DeepEqual(ctx, newCtx) {
			t.Error("expected context to be unchanged")
		}
	})
	t.Run("requires well-formed authorization when authorization is provided", func(t *testing.T) {
		mock := setup(t)
		_, aF := New(mock, 10*time.Second)

		// note: there are other malforms that could be checked
		malformedCtx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("authorization", "Bearer 0987 654321"))

		_, err := aF(malformedCtx)
		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires authorization that exists when authorization is provided", func(t *testing.T) {
		mock := setup(t)
		tokenExpiry := 10 * time.Second
		_, aF := New(mock, tokenExpiry)

		ctxWithNonExistantAuth := metadata.NewIncomingContext(context.Background(), metadata.Pairs("authorization", "Bearer 1234567890"))

		mock.EXPECT().Get(gomock.Any(), &pb_cacher.GetRequest{"1234567890"}).Return(nil, errors.New(""))

		_, err := aF(ctxWithNonExistantAuth)
		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires returning a new context for a valid authorization", func(t *testing.T) {
		user := &pb.User{Id: "id"}
		mock := setup(t)
		_, aF := New(mock, 10*time.Second)

		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("authorization", "Bearer 0987654321"))

		mock.EXPECT().Get(gomock.Any(), &pb_cacher.GetRequest{"0987654321"}).Return(&wrappers.StringValue{user.String()}, nil)

		newCtx, err := aF(ctx)
		if err != nil {
			t.Error("unexpected error:", err)
		}
		if reflect.DeepEqual(ctx, newCtx) {
			t.Error("expected new context to be different")
		}
	})
}
