package usecases

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/gosimple/slug"
	"github.com/satori/go.uuid"
	pb "github.com/thurt/demo-blog-platform/cms/proto"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type useCases struct {
	pb.CmsServer
}

type UseCases interface {
	CreatePost(context.Context, *pb.CreatePostRequest) (*pb.PostRequest, error)
	UpdatePost(context.Context, *pb.UpdatePostRequest) (*empty.Empty, error)
	CreateUser(context.Context, *pb.CreateUserRequest) (*pb.UserRequest, error)
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
	r.Slug = slugMake(r.GetTitle())
	return u.CmsServer.CreatePost(ctx, r)
}

func (u *useCases) UpdatePost(ctx context.Context, r *pb.UpdatePostRequest) (*empty.Empty, error) {
	r.Slug = slugMake(r.GetTitle())
	return u.CmsServer.UpdatePost(ctx, r)
}

func (u *useCases) CreateUser(ctx context.Context, r *pb.CreateUserRequest) (*pb.UserRequest, error) {
	hashedPassword, err := hashPassword(r.GetPassword())
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "Internal server error")
	}

	r.Password = hashedPassword
	return u.CmsServer.CreateUser(ctx, r)
}
