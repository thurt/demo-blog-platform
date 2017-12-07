package authentication

import (
	"crypto/rand"
	"fmt"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	pb "github.com/thurt/demo-blog-platform/cms/proto"
	"github.com/thurt/demo-blog-platform/cms/reqContext"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrUnauthenticated = status.Error(codes.Unauthenticated, codes.Unauthenticated.String())

type Identification struct {
	accessToken *pb.AccessToken
	user        *pb.User
	expiryTime  time.Time
}

type TokenHash map[string]*Identification

type authProvider struct {
	h           TokenHash
	tokenExpiry time.Duration
	tokenType   string
}

func newAuthFunc(h TokenHash) grpc_auth.AuthFunc {
	return func(ctx context.Context) (context.Context, error) {
		t, err := reqContext.GetAuthorizationToken(ctx)

		if err != nil {
			return nil, err
		}
		if t == "" {
			return ctx, nil
		}

		id, ok := h[t]
		if !ok {
			return nil, ErrUnauthenticated
		}

		if time.Now().After(id.expiryTime) {
			delete(h, t)
			return nil, ErrUnauthenticated
		}

		return reqContext.NewFromUser(ctx, id.user), nil
	}
}

func New(hash TokenHash, tokenExpiry time.Duration) (pb.CmsAuthServer, grpc_auth.AuthFunc) {
	return &authProvider{hash, tokenExpiry, "Bearer"}, newAuthFunc(hash)
}

func (a *authProvider) ActivateNewTokenForUser(ctx context.Context, r *pb.User) (*pb.AccessToken, error) {
	at := a.genAccessToken()

	a.h[at.GetAccessToken()] = &Identification{
		accessToken: at,
		user:        r,
		expiryTime:  time.Now().Add(a.tokenExpiry),
	}
	return at, nil
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
