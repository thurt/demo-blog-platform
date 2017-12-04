package mysqlprovider

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/thurt/demo-blog-platform/cms/proto"
	"github.com/xwb1989/sqlparser"
	"google.golang.org/grpc"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var (
	db   *sql.DB
	mock sqlmock.Sqlmock
	p    *provider
)

type mockCms_GetPostsServer struct {
	grpc.ServerStream
	Results []*pb.Post
}

func (m *mockCms_GetPostsServer) Send(p *pb.Post) error {
	m.Results = append(m.Results, p)
	return nil
}

type mockCms_GetPostCommentsServer struct {
	grpc.ServerStream
	Results []*pb.Comment
}

func (m *mockCms_GetPostCommentsServer) Send(c *pb.Comment) error {
	m.Results = append(m.Results, c)
	return nil
}

type mockCms_GetCommentsServer struct {
	grpc.ServerStream
	Results []*pb.Comment
}

func (m *mockCms_GetCommentsServer) Send(c *pb.Comment) error {
	m.Results = append(m.Results, c)
	return nil
}

type mockCms_GetUserCommentsServer struct {
	grpc.ServerStream
	Results []*pb.Comment
}

func (m *mockCms_GetUserCommentsServer) Send(c *pb.Comment) error {
	m.Results = append(m.Results, c)
	return nil
}

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags

	var err error
	// initialize the mock db and mock api
	db, mock, err = sqlmock.New()
	if err != nil {
		panic(fmt.Errorf("an error '%s' was not expected when opening a stub database connection", err))
	}

	// create the service provider
	p = New(db).(*provider)

	os.Exit(m.Run())
}

func checkExpectations(t *testing.T) {
	err := mock.ExpectationsWereMet()
	if err != nil {
		t.Error(t.Name(), err)
	}
}

func esc(stmt string) string {
	return regexp.QuoteMeta(stmt)
}

func checkQuerySyntax(query string, t *testing.T) {
	if _, err := sqlparser.Parse(query); err != nil {
		t.Error(err)
	}
}

func TestGetPost(t *testing.T) {
	mock.ExpectQuery(esc(p.q.GetPost())).WithArgs(0)
	r := &pb.PostRequest{Id: 0}

	_, _ = p.GetPost(context.Background(), r)

	checkExpectations(t)
}

func TestGetComment(t *testing.T) {
	mock.ExpectQuery(esc(p.q.GetComment())).WithArgs(0)
	r := &pb.CommentRequest{Id: 0}

	_, _ = p.GetComment(context.Background(), r)

	checkExpectations(t)
}

func TestGetUser(t *testing.T) {
	mock.ExpectQuery(esc(p.q.GetUser())).WithArgs("under_test")
	r := &pb.UserRequest{Id: "under_test"}

	_, _ = p.GetUser(context.Background(), r)

	checkExpectations(t)
}

func TestDeleteUser(t *testing.T) {
	mock.ExpectExec(esc(p.q.DeleteUser())).WithArgs("under_test")
	r := &pb.UserRequest{Id: "under_test"}

	_, _ = p.DeleteUser(context.Background(), r)

	checkExpectations(t)
}

func TestDeletePost(t *testing.T) {
	mock.ExpectExec(esc(p.q.DeletePost())).WithArgs(0)
	r := &pb.PostRequest{Id: 0}

	_, _ = p.DeletePost(context.Background(), r)

	checkExpectations(t)
}

func TestDeleteComment(t *testing.T) {
	mock.ExpectExec(esc(p.q.DeleteComment())).WithArgs(0)
	r := &pb.CommentRequest{Id: 0}

	_, _ = p.DeleteComment(context.Background(), r)

	checkExpectations(t)
}

func TestCreatePost(t *testing.T) {
	mock.ExpectExec(esc(p.q.CreatePost())).WithArgs("a-great-title", "A Great Title!", "content")
	r := &pb.CreatePostRequest{Title: "A Great Title!", Content: "content", Slug: "a-great-title"}

	_, _ = p.CreatePost(context.Background(), r)

	checkExpectations(t)
}

func TestCreateComment(t *testing.T) {
	mock.ExpectExec(esc(p.q.CreateComment())).WithArgs("content", "user_id", 0)
	r := &pb.CreateCommentRequest{Content: "content", UserId: "user_id", PostId: 0}

	_, _ = p.CreateComment(context.Background(), r)

	checkExpectations(t)
}

func TestCreateUser(t *testing.T) {
	mock.ExpectExec(esc(p.q.CreateUser())).WithArgs("id", "email", "password", defaultRole)
	r := &pb.CreateUserRequest{Id: "id", Email: "email", Password: "password"}

	_, _ = p.CreateUser(context.Background(), r)

	checkExpectations(t)
}

func TestPublishPost(t *testing.T) {
	mock.ExpectExec(esc(p.q.PublishPost())).WithArgs(0)
	r := &pb.PostRequest{Id: 0}

	_, _ = p.PublishPost(context.Background(), r)

	checkExpectations(t)
}

func TestUnPublishPost(t *testing.T) {
	mock.ExpectExec(esc(p.q.UnPublishPost())).WithArgs(0)
	r := &pb.PostRequest{Id: 0}

	_, _ = p.UnPublishPost(context.Background(), r)

	checkExpectations(t)
}

func TestUpdateComment(t *testing.T) {
	mock.ExpectExec(esc(p.q.UpdateComment())).WithArgs("content", 0)
	r := &pb.UpdateCommentRequest{Content: "content", Id: 0}

	_, _ = p.UpdateComment(context.Background(), r)

	checkExpectations(t)
}

func TestUpdatePost(t *testing.T) {
	mock.ExpectExec(esc(p.q.UpdatePost())).WithArgs("a-great-title", "A Great Title!", "content", 0)
	r := &pb.UpdatePostRequest{Title: "A Great Title!", Content: "content", Id: 0, Slug: "a-great-title"}

	_, _ = p.UpdatePost(context.Background(), r)

	checkExpectations(t)
}

func TestGetPostComments(t *testing.T) {
	mock.ExpectQuery(esc(p.q.GetPostComments())).WithArgs(0)
	r := &pb.PostRequest{Id: 0}
	s := &mockCms_GetPostCommentsServer{}

	_ = p.GetPostComments(r, s)

	checkExpectations(t)
}

func TestGetComments(t *testing.T) {
	mock.ExpectQuery(esc(p.q.GetComments()))
	r := &empty.Empty{}
	s := &mockCms_GetCommentsServer{}

	_ = p.GetComments(r, s)

	checkExpectations(t)
}

func TestGetPosts(t *testing.T) {
	mock.ExpectQuery(esc(p.q.GetPosts()))
	r := &empty.Empty{}
	s := &mockCms_GetPostsServer{}

	_ = p.GetPosts(r, s)

	checkExpectations(t)
}

func TestGetUserComments(t *testing.T) {
	mock.ExpectQuery(esc(p.q.GetUserComments())).WithArgs("id")
	r := &pb.UserRequest{Id: "id"}
	s := &mockCms_GetUserCommentsServer{}

	_ = p.GetUserComments(r, s)

	checkExpectations(t)
}
