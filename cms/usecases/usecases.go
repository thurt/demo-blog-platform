package usecases

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/gosimple/slug"
	"github.com/satori/go.uuid"
	pb "github.com/thurt/demo-blog-platform/cms/proto"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrAlreadyExists error = status.Error(codes.AlreadyExists, codes.AlreadyExists.String())

type useCases struct {
	pb.CmsServer
}

type UseCases interface {
	CreatePost(context.Context, *pb.CreatePostRequest) (*pb.PostRequest, error)
	UpdatePost(context.Context, *pb.UpdatePostRequest) (*empty.Empty, error)
	CreateUser(context.Context, *pb.CreateUserRequest) (*pb.UserRequest, error)
	CreateComment(context.Context, *pb.CreateCommentRequest) (*pb.CommentRequest, error)
}

func New(provider pb.CmsServer) *useCases {
	uc := &useCases{provider}
	return uc
}

func slugMake(str string) string {
	var s string

	if "" == str {
		s = uuid.NewV4().String()
	} else {
		slug.MaxLength = 36
		s = slug.Make(str)
	}
	return s
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (u *useCases) CreatePost(ctx context.Context, r *pb.CreatePostRequest) (*pb.PostRequest, error) {
	// requires a Slug to be created from the Title and added to the request
	r.Slug = slugMake(r.GetTitle())
	return u.CmsServer.CreatePost(ctx, r)
}

func (u *useCases) UpdatePost(ctx context.Context, r *pb.UpdatePostRequest) (*empty.Empty, error) {
	// requires a Slug to be created from the Title and added to the request
	r.Slug = slugMake(r.GetTitle())
	return u.CmsServer.UpdatePost(ctx, r)
}

func (u *useCases) CreateUser(ctx context.Context, r *pb.CreateUserRequest) (*pb.UserRequest, error) {
	// requires that user id does not exist
	_, err := u.GetUser(ctx, &pb.UserRequest{Id: r.GetId()})
	if err == nil {
		return nil, ErrAlreadyExists
	} else {
		s, ok := status.FromError(err)
		if !ok || s.Code() != codes.NotFound {
			return nil, err
		}
	}

	// requires that password is hashed
	hashedPassword, err := hashPassword(r.GetPassword())
	if err != nil {
		return nil, err
	}

	r.Password = hashedPassword
	return u.CmsServer.CreateUser(ctx, r)
}

func (u *useCases) CreateComment(ctx context.Context, r *pb.CreateCommentRequest) (*pb.CommentRequest, error) {
	// requires a valid user id
	_, err := u.GetUser(ctx, &pb.UserRequest{Id: r.GetUserId()})
	if err != nil {
		return nil, err
	}

	// requires a valid post id
	_, err = u.GetPost(ctx, &pb.PostRequest{Id: r.GetPostId()})
	if err != nil {
		return nil, err
	}

	return u.CmsServer.CreateComment(ctx, r)
}
