package authentication

import (
	"crypto/rand"
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	pb_cacher "github.com/thurt/demo-blog-platform/cms/cacher/proto"
	pb "github.com/thurt/demo-blog-platform/cms/proto"
	"github.com/thurt/demo-blog-platform/cms/reqContext"
	"golang.org/x/net/context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrUnauthenticated = status.Error(codes.Unauthenticated, "Your credentials are expired or invalid. Try to authenticate again.")

type authProvider struct {
	cacher      pb_cacher.CacherServer
	tokenExpiry time.Duration
	tokenType   string
}

func newAuthFunc(cacher pb_cacher.CacherServer, tokenExpiry time.Duration) grpc_auth.AuthFunc {
	return func(ctx context.Context) (context.Context, error) {
		t, err := reqContext.GetAuthorizationToken(ctx)

		if err != nil {
			return nil, err
		}
		if t == "" {
			return ctx, nil
		}

		userStr, err := cacher.Get(ctx, &pb_cacher.GetRequest{t})
		// maybe want to handle some possible errors returned here
		if err != nil {
			return nil, ErrUnauthenticated
		}

		return reqContext.NewFromUserString(ctx, userStr.GetValue()), nil
	}
}

func New(cacher pb_cacher.CacherServer, tokenExpiry time.Duration) (*authProvider, grpc_auth.AuthFunc) {
	return &authProvider{cacher, tokenExpiry, "Bearer"}, newAuthFunc(cacher, tokenExpiry)
}

func (a *authProvider) ActivateNewTokenForUser(ctx context.Context, r *pb.User) (*pb.AccessToken, error) {
	at := a.genAccessToken()

	_, err := a.cacher.Set(ctx, &pb_cacher.SetRequest{
		Key:   at.GetAccessToken(),
		Value: r.String(),
		Ttl:   &duration.Duration{Seconds: int64(a.tokenExpiry.Seconds())},
	})
	// maybe want to handle some possible errors returned here
	if err != nil {
		fmt.Println("error trying to add token to mc:", err)
		return nil, err
	}

	return at, nil
}

func (a *authProvider) ActivateNewTokenForCreateUserWithRole(ctx context.Context, r *pb.CreateUserWithRole) (*pb.AccessToken, error) {
	at := a.genAccessToken()

	_, err := a.cacher.Set(ctx, &pb_cacher.SetRequest{
		Key:   at.GetAccessToken(),
		Value: r.String(),
		Ttl:   &duration.Duration{Seconds: int64(a.tokenExpiry.Seconds())},
	})

	// maybe want to handle some possible errors returned here
	if err != nil {
		fmt.Println("error trying to add token to mc:", err)
		return nil, err
	}

	return at, nil
}

func (a *authProvider) DeactivateToken(ctx context.Context, r *wrappers.StringValue) (*empty.Empty, error) {
	_, err := a.cacher.Delete(ctx, &pb_cacher.DeleteRequest{r.GetValue()})
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (a *authProvider) genAccessToken() *pb.AccessToken {
	at := &pb.AccessToken{
		AccessToken: a.genRandTokenValue(),
		ExpiresIn:   uint32(a.tokenExpiry.Seconds()),
		TokenType:   a.tokenType,
	}
	return at
}

func (a *authProvider) genRandTokenValue() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
