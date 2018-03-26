package usecases

import (
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/gosimple/slug"
	"github.com/satori/go.uuid"
	"github.com/thurt/demo-blog-platform/cms/domain"
	pb "github.com/thurt/demo-blog-platform/cms/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type useCases struct {
	domain.Provider
	auth    pb.CmsAuthServer
	hasher  pb.HasherServer
	emailer pb.EmailerServer
}

func New(provider domain.Provider, authProvider pb.CmsAuthServer, hasher pb.HasherServer, emailer pb.EmailerServer) *useCases {
	uc := &useCases{provider, authProvider, hasher, emailer}
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
	cpws := &pb.CreatePostWithSlug{Post: r}
	cpws.Slug = slugMake(r.GetTitle())
	return u.Provider.CreatePost(ctx, cpws)
}

func (u *useCases) UpdatePost(ctx context.Context, r *pb.UpdatePostRequest) (*empty.Empty, error) {
	if r.GetPublished() == true && r.GetTitle() == "" {
		return nil, status.Error(codes.InvalidArgument, "A published post must have a title")
	}
	// requires a Slug to be created from the Title and added to the request
	upws := &pb.UpdatePostWithSlug{Post: r}
	upws.Slug = slugMake(r.GetTitle())
	return u.Provider.UpdatePost(ctx, upws)
}

func (u *useCases) GetUser(ctx context.Context, r *pb.UserRequest) (*pb.User, error) {
	user, err := u.Provider.GetUser(ctx, r)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, err.Error())
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
		log.Println(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	if *user != (pb.User{}) {
		return nil, status.Errorf(codes.AlreadyExists, "The provided user id %q already exists", r.GetId())
	}

	hashedPassword, err := u.hasher.Hash(ctx, &wrappers.StringValue{r.GetPassword()})
	if err != nil {
		return nil, err
	}

	r.Password = hashedPassword.GetValue()
	ur, err := u.Provider.CreateUser(ctx, &pb.CreateUserWithRole{User: r, Role: pb.UserRole_USER})
	if err != nil {
		return nil, err
	}

	_, err = u.emailer.Send(ctx, &pb.Email{
		To:      r.GetEmail(),
		From:    "no-reply@demo-blog-platform.com",
		Subject: "Thanks for joining Demo Blog!",
		Body:    "Hi, thanks for joining. \n\nThis is a confirmation email showing that you have successfully completed registration for Demo Blog with user id " + r.GetId() + ".",
	})

	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return ur, nil
}

func (u *useCases) CreateComment(ctx context.Context, r *pb.CreateCommentRequest) (*pb.CommentRequest, error) {
	// requires a valid user id
	_, err := u.Provider.GetUser(ctx, &pb.UserRequest{Id: r.GetUserId()})
	if err != nil {
		return nil, err
	}

	// requires a valid post id
	post, err := u.Provider.GetPost(ctx, &pb.PostRequest{Id: r.GetPostId()})
	if err != nil {
		return nil, err
	}
	if *post == (pb.Post{}) {
		return nil, status.Errorf(codes.NotFound, "The provided post id %q does not exist", r.GetPostId())
	}

	// Comment cannot be created for a Post that is not published
	if post.GetPublished() == false {
		return nil, status.Error(codes.InvalidArgument, "A Comment cannot be created for a Post that is not published")
	}

	return u.Provider.CreateComment(ctx, r)
}

func (u *useCases) AuthUser(ctx context.Context, r *pb.AuthUserRequest) (*pb.AccessToken, error) {
	var ErrUnauthenticated = status.Error(codes.Unauthenticated, "The provided username or password is incorrect")
	ur := &pb.UserRequest{r.GetId()}

	user, err := u.GetUser(ctx, ur)
	if err != nil {
		return nil, ErrUnauthenticated
	}

	p, err := u.Provider.GetUserPassword(ctx, ur)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	_, err = u.hasher.Validate(ctx, &pb.StrAndHash{r.GetPassword(), p.GetPassword()})
	if err != nil {
		return nil, ErrUnauthenticated
	}

	a, err := u.auth.ActivateNewTokenForUser(ctx, user)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	_, err = u.Provider.UpdateUserLastActive(ctx, ur)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, err.Error())
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
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return adminExists, nil
}

func (u *useCases) Setup(ctx context.Context, r *pb.CreateUserRequest) (*pb.UserRequest, error) {
	adminExists, err := u.Provider.AdminExists(ctx, &empty.Empty{})
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	if adminExists.GetValue() == true {
		return nil, status.Errorf(codes.Aborted, "Setup can only be performed once")
	}

	// requires that password is hashed
	hashedPassword, err := u.hasher.Hash(ctx, &wrappers.StringValue{r.GetPassword()})
	if err != nil {
		return nil, err
	}
	r.Password = hashedPassword.GetValue()

	user, err := u.Provider.CreateUser(ctx, &pb.CreateUserWithRole{Role: pb.UserRole_ADMIN, User: r})
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return user, nil
}

func (u *useCases) GetPosts(r *pb.GetPostsOptions, stream pb.Cms_GetPostsServer) error {
	err := u.Provider.GetPosts(r, stream)
	if err != nil {
		log.Println(err)
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func (u *useCases) GetPost(ctx context.Context, r *pb.PostRequest) (*pb.Post, error) {
	post, err := u.Provider.GetPost(ctx, r)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	if *post == (pb.Post{}) {
		return nil, status.Errorf(codes.NotFound, "The provided post id %q does not exist", r.GetId())
	}
	return post, nil
}

func (u *useCases) GetPostBySlug(ctx context.Context, r *pb.PostBySlugRequest) (*pb.Post, error) {
	post, err := u.Provider.GetPostBySlug(ctx, r)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	if *post == (pb.Post{}) {
		return nil, status.Errorf(codes.NotFound, "The provided post slug %q does not exist", r.GetSlug())
	}
	return post, nil
}

func (u *useCases) RegisterNewUser(ctx context.Context, r *pb.CreateUserRequest) (*empty.Empty, error) {
	// requires that user id does not exist
	user, err := u.Provider.GetUser(ctx, &pb.UserRequest{Id: r.GetId()})
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	if *user != (pb.User{}) {
		return nil, status.Errorf(codes.AlreadyExists, "The provided user id %q already exists", r.GetId())
	}

	hashedPassword, err := u.hasher.Hash(ctx, &wrappers.StringValue{r.GetPassword()})
	if err != nil {
		return nil, err
	}

	r.Password = hashedPassword.GetValue()
	at, err := u.auth.ActivateNewTokenForCreateUserWithRole(ctx, &pb.CreateUserWithRole{User: r, Role: pb.UserRole_USER})
	if err != nil {
		return nil, err
	}

	_, err = u.emailer.Send(ctx, &pb.Email{
		To:      r.GetEmail(),
		From:    "no-reply@demo-blog-platform.com",
		Subject: "Verify Your Email",
		Body:    "In order to complete you registration with user id, " + r.GetId() + ", you must copy the following value into the prompt as instructed on the Demo Blog Platform website: \n\n" + at.GetAccessToken(),
	})

	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

func (u *useCases) VerifyNewUser(ctx context.Context, _ *empty.Empty) (*pb.UserRequest, error) {
	return nil, status.Error(codes.Unimplemented, codes.Unimplemented.String())
}
