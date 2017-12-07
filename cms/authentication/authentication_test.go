package authentication

import (
	"reflect"
	"testing"
	"time"

	pb "github.com/thurt/demo-blog-platform/cms/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

func TestActivateNewTokenForUser(t *testing.T) {
	t.Run("requires that a new token is added", func(t *testing.T) {
		mockH := TokenHash{}
		aP, _ := New(mockH, 10)

		_, err := aP.ActivateNewTokenForUser(context.Background(), &pb.User{})
		if err != nil {
			t.Error("unexpected error")
		}

		if len(mockH) != 1 {
			t.Error("expected hash length to equal 1")
		}
	})
}

func TestAuthFunc(t *testing.T) {
	t.Run("requires well-formed grpc metadata-based context", func(t *testing.T) {
		mockH := TokenHash{}
		_, aF := New(mockH, 10*time.Second)

		nonMetadataCtx := context.WithValue(context.Background(), "authorization", []string{"Bearer 0987654321"})

		_, err := aF(nonMetadataCtx)
		if err == nil {
			t.Error("expected an error")
		}

	})
	t.Run("requires a no-op when no authorization is provided", func(t *testing.T) {
		mockH := TokenHash{}
		_, aF := New(mockH, 10*time.Second)

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
		mockH := TokenHash{}
		_, aF := New(mockH, 10*time.Second)

		// note: there are other malforms that could be checked
		malformedCtx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("authorization", "Bearer 0987 654321"))

		_, err := aF(malformedCtx)
		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires authorization that exists when authorization is provided", func(t *testing.T) {
		mockH := TokenHash{"0987654321": &Identification{}}
		_, aF := New(mockH, 10*time.Second)

		ctxWithNonExistantAuth := metadata.NewIncomingContext(context.Background(), metadata.Pairs("authorization", "Bearer 1234567890"))

		_, err := aF(ctxWithNonExistantAuth)
		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires authorization that is not expired", func(t *testing.T) {
		mockH := TokenHash{"0987654321": &Identification{expiryTime: time.Now()}}
		_, aF := New(mockH, 1*time.Second)

		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("authorization", "Bearer 0987654321"))

		_, err := aF(ctx)
		if err == nil {
			t.Error("expected an error")
		}
		if len(mockH) != 0 {
			t.Error("expected hash length 0, instead got", len(mockH))
		}
	})
	t.Run("requires returning a new context for a valid authorization", func(t *testing.T) {
		user := &pb.User{Id: "id"}
		mockH := TokenHash{"0987654321": &Identification{user: user, expiryTime: time.Now().Add(1 * time.Hour)}}
		_, aF := New(mockH, 1*time.Second)

		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("authorization", "Bearer 0987654321"))

		newCtx, err := aF(ctx)
		if err != nil {
			t.Error("unexpected error:", err)
		}
		if reflect.DeepEqual(ctx, newCtx) {
			t.Error("expected new context to be different")
		}
	})
}