package reqContext

import (
	"context"
	"strings"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	pb "github.com/thurt/demo-blog-platform/cms/proto"
)

var ErrUnauthenticated = status.Error(codes.Unauthenticated, "Your credentials are malformed and invalid. Try to authenticate again.")

func NewFromUser(ctx context.Context, u *pb.User) context.Context {
	md := metadata.Pairs("user", u.String())
	newMD := metadata.NewIncomingContext(ctx, md)
	return newMD
}
func NewFromUserString(ctx context.Context, u string) context.Context {
	md := metadata.Pairs("user", u)
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
func HasPermission(ctx context.Context, rolesAllowed ...pb.UserRole) bool {
	u, err := GetUser(ctx)
	if err != nil {
		return false
	}

	ur := u.GetRole()

	for _, r := range rolesAllowed {
		if r == ur {
			return true
		}
	}

	return false
}

func GetAuthorizationToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Error(codes.Internal, "malformed metadata from incoming context")
	}

	if md["authorization"] == nil || len(md["authorization"][0]) == 0 {
		return "", nil
	}

	authStr := strings.TrimSpace(md["authorization"][0])

	strParts := strings.Split(authStr, " ")
	if len(strParts) != 2 || strParts[0] != "Bearer" {
		return "", ErrUnauthenticated
	}
	token := strParts[1]

	return token, nil
}
