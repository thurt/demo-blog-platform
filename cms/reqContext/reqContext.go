package reqContext

import (
	"context"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	pb "github.com/thurt/demo-blog-platform/cms/proto"
)

func NewFromUser(ctx context.Context, u *pb.User) context.Context {
	md := metadata.Pairs("user", u.String())
	newMD := metadata.NewIncomingContext(ctx, md)
	return newMD
}

func GetUser(ctx context.Context) (*pb.User, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Internal, "metadata does not exist in context")
	}

	if md["user"] == nil || len(md["user"]) == 0 {
		return nil, status.Error(codes.Internal, "metadata does not contain user")
	}

	u := &pb.User{}
	err := proto.UnmarshalText(md["user"][0], u)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return u, nil
}
