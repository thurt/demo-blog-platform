package hasher

import (
	"golang.org/x/net/context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	pb "github.com/thurt/demo-blog-platform/cms/proto"
	"golang.org/x/crypto/bcrypt"
)

type hasher struct{}

func New() pb.HasherServer {
	return &hasher{}
}

func (h *hasher) Hash(ctx context.Context, str *wrappers.StringValue) (*wrappers.StringValue, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str.GetValue()), 14)
	return &wrappers.StringValue{string(bytes)}, err
}

func (h *hasher) Validate(ctx context.Context, strAndHash *pb.StrAndHash) (*empty.Empty, error) {
	err := bcrypt.CompareHashAndPassword([]byte(strAndHash.GetHash()), []byte(strAndHash.GetStr()))
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
