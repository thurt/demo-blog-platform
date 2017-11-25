package mysqlprovider

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"
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
	p    *Provider
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
	p = New(db).(*Provider)

	os.Exit(m.Run())
}

func checkExpectations(t *testing.T) {
	err := mock.ExpectationsWereMet()
	if err != nil {
		t.Error(t.Name(), err)
	}
}

// rpl replaces any "?" in the sql statement to ".+"
// this transformation is required because the mock library makes expectations based on a regex string instead of a regular string
func rpl(stmt string) string {
	return strings.Replace(stmt, "?", ".+", -1)
}

func checkQuerySyntax(query string, t *testing.T) {
	if _, err := sqlparser.Parse(query); err != nil {
		t.Error(err)
	}
}

func TestGetPost(t *testing.T) {
	mock.ExpectQuery(p.q.GetPost()).WithArgs("under_test")
	r := &pb.PostRequest{Id: "under_test"}

	_, _ = p.GetPost(context.Background(), r)

	checkExpectations(t)
}

func TestGetComment(t *testing.T) {
	mock.ExpectQuery(p.q.GetComment()).WithArgs(0)
	r := &pb.CommentRequest{Id: 0}

	_, _ = p.GetComment(context.Background(), r)

	checkExpectations(t)
}

func TestGetUser(t *testing.T) {
	mock.ExpectQuery(p.q.GetUser()).WithArgs("under_test")
	r := &pb.UserRequest{Id: "under_test"}

	_, _ = p.GetUser(context.Background(), r)

	checkExpectations(t)
}

func TestDeleteUser(t *testing.T) {
	mock.ExpectExec(p.q.DeleteUser()).WithArgs("under_test")
	r := &pb.UserRequest{Id: "under_test"}

	_, _ = p.DeleteUser(context.Background(), r)

	checkExpectations(t)
}

func TestDeletePost(t *testing.T) {
	mock.ExpectExec(p.q.DeletePost()).WithArgs("under_test")
	r := &pb.PostRequest{Id: "under_test"}

	_, _ = p.DeletePost(context.Background(), r)

	checkExpectations(t)
}

func TestDeleteComment(t *testing.T) {
	mock.ExpectExec(p.q.DeleteComment()).WithArgs(0)
	r := &pb.CommentRequest{Id: 0}

	_, _ = p.DeleteComment(context.Background(), r)

	checkExpectations(t)
}

func TestCreatePost(t *testing.T) {
	// TODO: change hard-coded -- see comment in source code
	mock.ExpectExec(rpl(p.q.CreatePost())).WithArgs("hard-coded", "title", "content")
	r := &pb.CreatePostRequest{Title: "title", Content: "content"}

	_, _ = p.CreatePost(context.Background(), r)

	checkExpectations(t)
}

func TestCreateComment(t *testing.T) {
	mock.ExpectExec(rpl(p.q.CreateComment())).WithArgs("content", "user_id", "post_id")
	r := &pb.CreateCommentRequest{Content: "content", UserId: "user_id", PostId: "post_id"}

	_, _ = p.CreateComment(context.Background(), r)

	checkExpectations(t)
}

func TestCreateUser(t *testing.T) {
	mock.ExpectExec(rpl(p.q.CreateUser())).WithArgs("id", "email", "password")
	r := &pb.CreateUserRequest{Id: "id", Email: "email", Password: "password"}

	_, _ = p.CreateUser(context.Background(), r)

	checkExpectations(t)
}

func TestPublishPost(t *testing.T) {
	mock.ExpectExec(rpl(p.q.PublishPost())).WithArgs("id")
	r := &pb.PostRequest{Id: "id"}

	_, _ = p.PublishPost(context.Background(), r)

	checkExpectations(t)
}

func TestUnPublishPost(t *testing.T) {
	mock.ExpectExec(rpl(p.q.UnPublishPost())).WithArgs("id")
	r := &pb.PostRequest{Id: "id"}

	_, _ = p.UnPublishPost(context.Background(), r)

	checkExpectations(t)
}

func TestUpdateComment(t *testing.T) {
	mock.ExpectExec(rpl(p.q.UpdateComment())).WithArgs("content", 0)
	r := &pb.UpdateCommentRequest{Content: "content", Id: 0}

	_, _ = p.UpdateComment(context.Background(), r)

	checkExpectations(t)
}

func TestUpdatePost(t *testing.T) {
	mock.ExpectExec(rpl(p.q.UpdatePost())).WithArgs("title", "content", "id")
	r := &pb.UpdatePostRequest{Title: "title", Content: "content", Id: "id"}

	_, _ = p.UpdatePost(context.Background(), r)

	checkExpectations(t)
}

func TestGetPostComments(t *testing.T) {
	mock.ExpectQuery(rpl(p.q.GetPostComments())).WithArgs("id")
	r := &pb.PostRequest{Id: "id"}
	s := &mockCms_GetPostCommentsServer{}

	_ = p.GetPostComments(r, s)

	checkExpectations(t)
}

func TestGetComments(t *testing.T) {
	mock.ExpectQuery(rpl(p.q.GetComments()))
	r := &empty.Empty{}
	s := &mockCms_GetCommentsServer{}

	_ = p.GetComments(r, s)

	checkExpectations(t)
}

func TestGetPosts(t *testing.T) {
	mock.ExpectQuery(rpl(p.q.GetPosts()))
	r := &empty.Empty{}
	s := &mockCms_GetPostsServer{}

	_ = p.GetPosts(r, s)

	checkExpectations(t)
}

func TestGetUserComments(t *testing.T) {
	mock.ExpectQuery(rpl(p.q.GetUserComments())).WithArgs("id")
	r := &pb.UserRequest{Id: "id"}
	s := &mockCms_GetUserCommentsServer{}

	_ = p.GetUserComments(r, s)

	checkExpectations(t)
}
