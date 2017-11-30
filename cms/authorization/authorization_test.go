package authorization

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/thurt/demo-blog-platform/cms/mock_proto"
	pb "github.com/thurt/demo-blog-platform/cms/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

func setup(t *testing.T) (*mock_proto.MockCmsServer, Authorization) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mock := mock_proto.NewMockCmsServer(mockCtrl)
	a := New(mock)
	return mock, a
}

func TestDeleteComment(t *testing.T) {
	t.Skip()
}

func TestUpdateComment(t *testing.T) {
	t.Skip()
}

func TestCreateComment(t *testing.T) {
	t.Skip()
}

func TestDeleteUser(t *testing.T) {
	t.Skip()
}

func TestUnPublishPost(t *testing.T) {
	t.Skip()
}

func TestPublishPost(t *testing.T) {
	t.Skip()
}

func TestDeletePost(t *testing.T) {
	t.Skip()
}

func TestUpdatePost(t *testing.T) {
	t.Skip()
}

func TestCreatePost(t *testing.T) {
	t.Run("requires Admin Role has permission", func(t *testing.T) {
		mock, a := setup(t)

		md := metadata.Pairs("role", "admin")
		ctx := metadata.NewIncomingContext(context.Background(), md)
		r := &pb.CreatePostRequest{}

		mock.EXPECT().CreatePost(ctx, r)

		_, err := a.CreatePost(ctx, r)
		if err != nil {
			t.Error("unexpected error:", err)
		}
	})
}
