package cacher

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes/duration"
	pb "github.com/thurt/demo-blog-platform/cms/cacher/proto"
	"github.com/thurt/demo-blog-platform/cms/mock_mc"
)

func setup(t *testing.T) (*mock_mc.MockConn, *cacher) {
	// create NewMockConn
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mock := mock_mc.NewMockConn(mockCtrl)

	return mock, New(mock)
}

func TestGet(t *testing.T) {
	stubIn := &pb.GetRequest{}
	stubIn.Key = "abc"
	t.Run("must return error when memcached returns error", func(t *testing.T) {
		mock, c := setup(t)
		mock.EXPECT().Get(stubIn.GetKey()).Return("", uint32(0), uint64(0), errors.New(""))

		_, err := c.Get(context.Background(), stubIn)

		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("must return without error under normal circumstances", func(t *testing.T) {
		mock, c := setup(t)
		mock.EXPECT().Get(stubIn.GetKey()).Return("", uint32(0), uint64(0), nil)

		_, err := c.Get(context.Background(), stubIn)

		if err != nil {
			t.Error("unexpected error", err.Error())
		}

	})

}

func TestSet(t *testing.T) {
	stubIn := &pb.SetRequest{}
	stubIn.Key = "abc"
	stubIn.Value = "xyz"
	stubIn.Ttl = &duration.Duration{Seconds: int64(10)}
	t.Run("must return error when memcached returns error", func(t *testing.T) {
		mock, c := setup(t)
		mock.EXPECT().Set(stubIn.GetKey(), stubIn.GetValue(), gomock.Any(), gomock.Any(), gomock.Any()).Return(uint64(0), errors.New(""))

		_, err := c.Set(context.Background(), stubIn)

		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("must return without error under normal circumstances", func(t *testing.T) {
		mock, c := setup(t)
		mock.EXPECT().Set(stubIn.GetKey(), stubIn.GetValue(), gomock.Any(), uint32(stubIn.GetTtl().GetSeconds()), gomock.Any()).Return(uint64(0), nil)

		_, err := c.Set(context.Background(), stubIn)

		if err != nil {
			t.Error("unexpected error", err.Error())
		}
	})
}

func TestDelete(t *testing.T) {
	stubIn := &pb.DeleteRequest{}
	stubIn.Key = "abc"
	t.Run("must return error when memcached returns error", func(t *testing.T) {
		mock, c := setup(t)
		mock.EXPECT().Del(stubIn.GetKey()).Return(errors.New(""))

		_, err := c.Delete(context.Background(), stubIn)

		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("must return without error under normal circumstances", func(t *testing.T) {
		mock, c := setup(t)
		mock.EXPECT().Del(stubIn.GetKey()).Return(nil)

		_, err := c.Delete(context.Background(), stubIn)

		if err != nil {
			t.Error("unexpected error", err.Error())
		}
	})
}
