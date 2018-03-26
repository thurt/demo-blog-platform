package authentication

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/thurt/demo-blog-platform/cms/mock_mc"
	pb "github.com/thurt/demo-blog-platform/cms/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

func setup(t *testing.T) *mock_mc.MockConn {
	// create NewMockConn
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mock := mock_mc.NewMockConn(mockCtrl)
	return mock
}

func TestActivateNewTokenForUser(t *testing.T) {
	t.Run("requires that a new token is added", func(t *testing.T) {
		mock := setup(t)
		aP, _ := New(mock, 10*time.Second)

		mock.EXPECT().Add(gomock.Any(), gomock.Any(), gomock.Any(), uint32(10*time.Second.Seconds())).Return(uint64(0), nil)

		_, err := aP.ActivateNewTokenForUser(context.Background(), &pb.User{})
		if err != nil {
			t.Error("unexpected error")
		}
	})
}

func TestActivateNewTokenForCreateUserWithRole(t *testing.T) {
	t.Run("requires that a new token is added", func(t *testing.T) {
		mock := setup(t)
		aP, _ := New(mock, 10*time.Second)

		mock.EXPECT().Add(gomock.Any(), gomock.Any(), gomock.Any(), uint32(10*time.Second.Seconds())).Return(uint64(0), nil)

		_, err := aP.ActivateNewTokenForCreateUserWithRole(context.Background(), &pb.CreateUserWithRole{})
		if err != nil {
			t.Error("unexpected error")
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

		mock.EXPECT().GAT("1234567890", uint32(tokenExpiry.Seconds())).Return("", uint32(0), uint64(0), errors.New(""))

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

		mock.EXPECT().GAT("0987654321", uint32(10*time.Second.Seconds())).Return(user.String(), uint32(0), uint64(0), nil)

		newCtx, err := aF(ctx)
		if err != nil {
			t.Error("unexpected error:", err)
		}
		if reflect.DeepEqual(ctx, newCtx) {
			t.Error("expected new context to be different")
		}
	})
}
