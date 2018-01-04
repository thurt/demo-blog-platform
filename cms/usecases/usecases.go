package usecases

import (
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/gosimple/slug"
	"github.com/satori/go.uuid"
	"github.com/thurt/demo-blog-platform/cms/domain"
	"github.com/thurt/demo-blog-platform/cms/password"
	pb "github.com/thurt/demo-blog-platform/cms/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type useCases struct {
	domain.Provider
	internal pb.CmsInternalServer
	auth     pb.CmsAuthServer
}

func New(provider domain.Provider, internalProvider pb.CmsInternalServer, authProvider pb.CmsAuthServer) *useCases {
	uc := &useCases{provider, internalProvider, authProvider}
	return uc
}

func slugMake(str string) string {
	var s string

	if "" == str {
		s = uuid.Must(uuid.NewV4()).String()
	} else {
		slug.MaxLength = 36
		s = slug.Make(str)
	}
	return s
}

func (u *useCases) CreatePost(ctx context.Context, r *pb.CreatePostRequest) (*pb.PostRequest, error) {
	// requires a Slug to be created from the Title and added to the request
	r.Slug = slugMake(r.GetTitle())
	return u.Provider.CreatePost(ctx, r)
}

func (u *useCases) UpdatePost(ctx context.Context, r *pb.UpdatePostRequest) (*empty.Empty, error) {
	// requires a Slug to be created from the Title and added to the request
	r.Slug = slugMake(r.GetTitle())
	return u.Provider.UpdatePost(ctx, r)
}

func (u *useCases) GetUser(ctx context.Context, r *pb.UserRequest) (*pb.User, error) {
	user, err := u.Provider.GetUser(ctx, r)
	if err != nil {
		return nil, status.Error(codes.Internal, codes.Internal.String())
	}
	if *user == (pb.User{}) {
		return nil, status.Errorf(codes.NotFound, "The provided user id %q does not exist", r.GetId())
	}
	return user, nil
}

func (u *useCases) CreateUser(ctx context.Context, r *pb.CreateUserRequest) (*pb.UserRequest, error) {
	// requires that user id does not exist
	user, err := u.Provider.GetUser(ctx, &pb.UserRequest{Id: r.GetId()})
	if err != nil {
		return nil, status.Error(codes.Internal, codes.Internal.String())
	}

	if *user != (pb.User{}) {
		return nil, status.Errorf(codes.AlreadyExists, "The provided user id %q already exists", r.GetId())
	}

	// requires that password is hashed
	hashedPassword, err := password.Hash(r.GetPassword())
	if err != nil {
		return nil, err
	}

	r.Password = hashedPassword
	return u.Provider.CreateUser(ctx, &pb.CreateUserWithRole{User: r, Role: pb.UserRole_USER})
}

func (u *useCases) CreateComment(ctx context.Context, r *pb.CreateCommentRequest) (*pb.CommentRequest, error) {
	// requires a valid user id
	_, err := u.Provider.GetUser(ctx, &pb.UserRequest{Id: r.GetUserId()})
	if err != nil {
		return nil, err
	}

	// requires a valid post id
	_, err = u.GetPost(ctx, &pb.PostRequest{Id: r.GetPostId()})
	if err != nil {
		return nil, err
	}

	return u.Provider.CreateComment(ctx, r)
}

func (u *useCases) AuthUser(ctx context.Context, r *pb.AuthUserRequest) (*pb.AccessToken, error) {
	ur := &pb.UserRequest{r.GetId()}

	user, err := u.Provider.GetUser(ctx, ur)
	if err != nil {
		return nil, err
	}

	p, err := u.internal.GetUserPassword(ctx, &pb.UserRequest{r.GetId()})
	if err != nil {
		return nil, err
	}

	err = password.Validate(r.GetPassword(), p.GetPassword())
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "The provided password does not match")
	}

	a, err := u.auth.ActivateNewTokenForUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (u *useCases) GetComment(ctx context.Context, r *pb.CommentRequest) (*pb.Comment, error) {
	comment, _ := u.Provider.GetComment(ctx, r)
	if *comment == (pb.Comment{}) {
		return nil, status.Errorf(codes.NotFound, "The provided comment id %q does not exist", r.GetId())
	}
	return comment, nil
}

func (u *useCases) IsSetup(ctx context.Context, r *empty.Empty) (*wrappers.BoolValue, error) {
	adminExists, err := u.Provider.AdminExists(ctx, r)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, codes.Internal.String())
	}
	return adminExists, nil
}

func (u *useCases) Setup(ctx context.Context, r *pb.CreateUserRequest) (*pb.UserRequest, error) {
	adminExists, err := u.Provider.AdminExists(ctx, &empty.Empty{})
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, codes.Internal.String())
	}
	if adminExists.GetValue() == true {
		return nil, status.Errorf(codes.Aborted, "Setup can only be performed when an admin account does not already exist")
	}

	user, err := u.Provider.CreateUser(ctx, &pb.CreateUserWithRole{Role: pb.UserRole_ADMIN, User: r})
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, codes.Internal.String())
	}

	return user, nil
}
