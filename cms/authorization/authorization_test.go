package authorization

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/thurt/demo-blog-platform/cms/mock_proto"
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
	t.Skip()
}
