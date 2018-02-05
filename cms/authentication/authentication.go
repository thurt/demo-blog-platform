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

var ErrUnauthenticated = status.Error(codes.Unauthenticated, "Your credentials are expired or invalid. Try to authenticate again.")

type authProvider struct {
	mc          Conn
	tokenExpiry time.Duration
	tokenType   string
}

func newAuthFunc(mc Conn, tokenExpiry time.Duration) grpc_auth.AuthFunc {
	return func(ctx context.Context) (context.Context, error) {
		t, err := reqContext.GetAuthorizationToken(ctx)

		if err != nil {
			return nil, err
		}
		if t == "" {
			return ctx, nil
		}

		// get and touch will reset expiry to the tokenExpiry value
		// this allows a token to remain active so long as it is still being used
		userStr, _, _, err := mc.GAT(t, uint32(tokenExpiry.Seconds()))
		// maybe want to handle some possible errors returned here
		if err != nil {
			return nil, ErrUnauthenticated
		}

		return reqContext.NewFromUserString(ctx, userStr), nil
	}
}

func New(mc Conn, tokenExpiry time.Duration) (*authProvider, grpc_auth.AuthFunc) {
	return &authProvider{mc, tokenExpiry, "Bearer"}, newAuthFunc(mc, tokenExpiry)
}

func (a *authProvider) ActivateNewTokenForUser(ctx context.Context, r *pb.User) (*pb.AccessToken, error) {
	at := a.genAccessToken()

	_, err := a.mc.Add(at.GetAccessToken(), r.String(), 0, uint32(a.tokenExpiry.Seconds()))
	// maybe want to handle some possible errors returned here
	if err != nil {
		fmt.Println("error trying to add token to mc:", err)
		return nil, err
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
