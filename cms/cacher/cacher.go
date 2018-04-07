package cacher

import (
	"golang.org/x/net/context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	pb "github.com/thurt/demo-blog-platform/cms/cacher/proto"
)

type cacher struct {
	mc Conn
}

func New(mc Conn) *cacher {
	return &cacher{mc}
}

func (c *cacher) Get(ctx context.Context, r *pb.GetRequest) (*wrappers.StringValue, error) {
	val, _, _, err := c.mc.Get(r.GetKey())
	if err != nil {
		return nil, err
	}
	return &wrappers.StringValue{val}, nil
}

func (c *cacher) Set(ctx context.Context, r *pb.SetRequest) (*empty.Empty, error) {
	_, err := c.mc.Set(r.GetKey(), r.GetValue(), 0, uint32(r.GetTtl().GetSeconds()), 0)
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
