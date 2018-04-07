package usecases

import (
	"crypto/rand"
	"fmt"
	"log"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/gosimple/slug"
	"github.com/satori/go.uuid"
	pb_cacher "github.com/thurt/demo-blog-platform/cms/cacher/proto"
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
	cacher  pb_cacher.CacherServer
}

func New(provider domain.Provider, authProvider pb.CmsAuthServer, hasher pb.HasherServer, emailer pb.EmailerServer, cacher pb_cacher.CacherServer) *useCases {
	uc := &useCases{provider, authProvider, hasher, emailer, cacher}
	return uc
}

func genRandHexValue(bytes uint) string {
	if bytes == 0 {
		return ""
	}
	b := make([]byte, bytes)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
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

	user, err := u.Provider.CreateNewUser(ctx, &pb.CreateUserWithRole{Role: pb.UserRole_ADMIN, User: r})
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

	cuwr := &pb.CreateUserWithRole{User: r, Role: pb.UserRole_USER}
	tv := genRandHexValue(3)
	_, err = u.cacher.Set(ctx, &pb_cacher.SetRequest{
		Key:   tv,
		Value: cuwr.String(),
		Ttl:   &duration.Duration{Seconds: int64((24 * time.Hour).Seconds())},
	})
	if err != nil {
		return nil, err
	}

	_, err = u.emailer.Send(ctx, &pb.Email{
		To:      r.GetEmail(),
		From:    "no-reply@demo-blog-platform.com",
		Subject: "Verify Your Email",
		Body:    "In order to complete you registration with user id, " + r.GetId() + ", you must copy the following value into the prompt as instructed on the Demo Blog Platform website: \n\n" + tv,
	})

	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

func (u *useCases) VerifyNewUser(ctx context.Context, r *pb.VerifyNewUserRequest) (*pb.UserRequest, error) {
	v, err := u.cacher.Get(ctx, &pb_cacher.GetRequest{Key: r.GetToken()})
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	// unmarshal value returned from cacher
	cuwr := &pb.CreateUserWithRole{}
	err = proto.UnmarshalText(v.GetValue(), cuwr)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	// requires that user id does not exist
	user, err := u.Provider.GetUser(ctx, &pb.UserRequest{Id: cuwr.GetUser().GetId()})
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	if *user != (pb.User{}) {
		return nil, status.Errorf(codes.AlreadyExists, "The provided user id %q already exists", cuwr.GetUser().GetId())
	}

	ur, err := u.Provider.CreateNewUser(ctx, cuwr)
	if err != nil {
		return nil, err
	}

	_, err = u.emailer.Send(ctx, &pb.Email{
		To:      cuwr.GetUser().GetEmail(),
		From:    "no-reply@demo-blog-platform.com",
		Subject: "Registration Complete",
		Body:    "Hi, thanks for joining Demo Blog! \n\nThis is a confirmation email that you have successfully completed registration for Demo Blog with user id " + cuwr.GetUser().GetId() + ".",
	})

	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return ur, nil
}

func (u *useCases) Logout(ctx context.Context, r *pb.AccessToken) (*empty.Empty, error) {
	_, err := u.auth.DeactivateToken(ctx, &wrappers.StringValue{r.GetAccessToken()})
	if err != nil {
		return nil, status.Error(codes.Internal, codes.Internal.String())
	}
	return &empty.Empty{}, nil
}
