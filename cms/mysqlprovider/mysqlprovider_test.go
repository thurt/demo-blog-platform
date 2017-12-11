package mysqlprovider

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"reflect"

	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/fatih/structs"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/gofuzz"
	pb "github.com/thurt/demo-blog-platform/cms/proto"
	"google.golang.org/grpc"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var (
	db   *sql.DB
	mock sqlmock.Sqlmock
	p    *provider
	f    *fuzz.Fuzzer
)

var regexAny string = ".*"

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

	// create the fuzzer
	f = fuzz.New()

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

// makeRowData creates a slice of driver.Value by casting provided values to driver.Value,
// which is the expected input type for sqlmock.Rows{}.AddRow
func makeRowData(values []interface{}) []driver.Value {
	driverVals := make([]driver.Value, len(values))
	for i, v := range values {
		driverVals[i] = v.(driver.Value)
	}
	return driverVals
}

func TestGetPost(t *testing.T) {
	stubIn := &pb.PostRequest{}
	stubOut := &pb.Post{}
	f.Fuzz(stubIn)
	f.Fuzz(stubOut)
	regexSql := esc(p.q.GetPost(stubIn))
	stubRows := sqlmock.NewRows(structs.Names(stubOut))
	stubRows.AddRow(makeRowData(structs.Values(stubOut))...)

	t.Run("requires dispatching the correct sql request", func(t *testing.T) {
		mock.ExpectQuery(regexSql)

		_, _ = p.GetPost(context.Background(), stubIn)

		checkExpectations(t)
	})
	t.Run("requires returning result with correct values from sql response", func(t *testing.T) {
		mock.ExpectQuery(regexAny).WillReturnRows(stubRows)

		result, err := p.GetPost(context.Background(), stubIn)
		if err != nil {
			t.Error("unexpected error:", err.Error())
		}

		if !reflect.DeepEqual(result, stubOut) {
			t.Error("result should have same values as stub values")
		}
	})
	t.Run("requires returning error when sql response is an error", func(t *testing.T) {
		mock.ExpectQuery(regexAny).WillReturnError(errors.New(""))

		_, err := p.GetPost(context.Background(), stubIn)
		if err == nil {
			t.Error("expected an error")
		}
	})
}

func TestGetComment(t *testing.T) {
	r := &pb.CommentRequest{Id: 0}
	mock.ExpectQuery(p.q.GetComment(r)).WillReturnRows(&sqlmock.Rows{})

	_, _ = p.GetComment(context.Background(), r)

	checkExpectations(t)
}

func TestGetUser(t *testing.T) {
	r := &pb.UserRequest{Id: "id"}
	mock.ExpectQuery(p.q.GetUser(r)).WillReturnRows(&sqlmock.Rows{})

	_, _ = p.GetUser(context.Background(), r)

	checkExpectations(t)
}

func TestDeleteUser(t *testing.T) {
	r := &pb.UserRequest{Id: "id"}
	mock.ExpectExec(p.q.DeleteUser(r)).WillReturnResult(sqlmock.NewResult(1, 1))

	_, _ = p.DeleteUser(context.Background(), r)

	checkExpectations(t)
}

func TestDeletePost(t *testing.T) {
	stubIn := &pb.PostRequest{}
	stubOut := &empty.Empty{}
	f.Fuzz(stubIn)
	f.Fuzz(stubOut)
	regexSql := esc(p.q.DeletePost(stubIn))
	stubResult := sqlmock.NewResult(int64(stubIn.Id), 1)

	t.Run("requires dispatching the correct sql request", func(t *testing.T) {
		mock.ExpectExec(regexSql)

		_, _ = p.DeletePost(context.Background(), stubIn)

		checkExpectations(t)
	})
	t.Run("requires returning result with correct values from sql response", func(t *testing.T) {
		mock.ExpectExec(regexAny).WillReturnResult(stubResult)

		result, err := p.DeletePost(context.Background(), stubIn)
		if err != nil {
			t.Error("unexpected error:", err.Error())
		}

		if !reflect.DeepEqual(result, stubOut) {
			t.Error("result should have same values as stub values")
		}
	})
	t.Run("requires returning error when sql response is an error", func(t *testing.T) {
		mock.ExpectExec(regexAny).WillReturnError(errors.New(""))

		_, err := p.DeletePost(context.Background(), stubIn)
		if err == nil {
			t.Error("expected an error")
		}
	})
}

func TestDeleteComment(t *testing.T) {
	r := &pb.CommentRequest{Id: 0}
	mock.ExpectExec(p.q.DeleteComment(r)).WillReturnResult(sqlmock.NewResult(1, 1))

	_, _ = p.DeleteComment(context.Background(), r)

	checkExpectations(t)
}

func TestCreatePost(t *testing.T) {
	stubIn := &pb.CreatePostRequest{}
	stubOut := &pb.PostRequest{}
	f.Fuzz(stubIn)
	f.Fuzz(stubOut)
	regexSql := esc(p.q.CreatePost(stubIn))
	stubResult := sqlmock.NewResult(int64(stubOut.Id), 1)

	t.Run("requires dispatching the correct sql request", func(t *testing.T) {
		mock.ExpectExec(regexSql)

		_, _ = p.CreatePost(context.Background(), stubIn)

		checkExpectations(t)
	})
	t.Run("requires returning result with correct values from sql response", func(t *testing.T) {
		mock.ExpectExec(regexAny).WillReturnResult(stubResult)

		result, err := p.CreatePost(context.Background(), stubIn)
		if err != nil {
			t.Error("unexpected error:", err.Error())
		}

		if !reflect.DeepEqual(result, stubOut) {
			t.Error("result should have same values as stub values")
		}
	})
	t.Run("requires returning error when sql response is an error", func(t *testing.T) {
		mock.ExpectExec(regexAny).WillReturnError(errors.New(""))

		_, err := p.CreatePost(context.Background(), stubIn)
		if err == nil {
			t.Error("expected an error")
		}
	})
}

func TestCreateComment(t *testing.T) {
	stubIn := &pb.CreateCommentRequest{}
	stubOut := &pb.CommentRequest{}
	f.Fuzz(stubIn)
	f.Fuzz(stubOut)
	regexSql := esc(p.q.CreateComment(stubIn))
	stubResult := sqlmock.NewResult(int64(stubOut.Id), 1)

	t.Run("requires dispatching the correct sql request", func(t *testing.T) {
		mock.ExpectExec(regexSql)

		_, _ = p.CreateComment(context.Background(), stubIn)

		checkExpectations(t)
	})
	t.Run("requires returning result with correct values from sql response", func(t *testing.T) {
		mock.ExpectExec(regexAny).WillReturnResult(stubResult)

		result, err := p.CreateComment(context.Background(), stubIn)
		if err != nil {
			t.Error("unexpected error:", err.Error())
		}

		if !reflect.DeepEqual(result, stubOut) {
			t.Error("result should have same values as stub values")
		}
	})
	t.Run("requires returning error when sql response is an error", func(t *testing.T) {
		mock.ExpectExec(regexAny).WillReturnError(errors.New(""))

		_, err := p.CreateComment(context.Background(), stubIn)
		if err == nil {
			t.Error("expected an error")
		}
	})
}

func TestCreateUser(t *testing.T) {
	r := &pb.CreateUserRequest{Id: "id", Email: "email", Password: "password"}
	mock.ExpectExec(p.q.CreateUser(r)).WillReturnResult(sqlmock.NewResult(1, 1))

	_, _ = p.CreateUser(context.Background(), r)

	checkExpectations(t)
}

func TestPublishPost(t *testing.T) {
	r := &pb.PostRequest{Id: 0}
	mock.ExpectExec(p.q.PublishPost(r)).WillReturnResult(sqlmock.NewResult(1, 1))

	_, _ = p.PublishPost(context.Background(), r)

	checkExpectations(t)
}

func TestUnPublishPost(t *testing.T) {
	r := &pb.PostRequest{Id: 0}
	mock.ExpectExec(p.q.UnPublishPost(r)).WillReturnResult(sqlmock.NewResult(1, 1))

	_, _ = p.UnPublishPost(context.Background(), r)

	checkExpectations(t)
}

func TestUpdateComment(t *testing.T) {
	r := &pb.UpdateCommentRequest{Content: "content", Id: 0}
	mock.ExpectExec(p.q.UpdateComment(r)).WillReturnResult(sqlmock.NewResult(1, 1))

	_, _ = p.UpdateComment(context.Background(), r)

	checkExpectations(t)
}

func TestUpdatePost(t *testing.T) {
	r := &pb.UpdatePostRequest{Title: "A Great Title!", Content: "content", Id: 0, Slug: "a-great-title"}
	mock.ExpectExec(p.q.UpdatePost(r)).WillReturnResult(sqlmock.NewResult(1, 1))

	_, _ = p.UpdatePost(context.Background(), r)

	checkExpectations(t)
}

func TestGetPostComments(t *testing.T) {
	r := &pb.PostRequest{Id: 0}
	mock.ExpectQuery(p.q.GetPostComments(r)).WillReturnRows(&sqlmock.Rows{})
	s := &mockCms_GetPostCommentsServer{}

	_ = p.GetPostComments(r, s)

	checkExpectations(t)
}

func TestGetComments(t *testing.T) {
	mock.ExpectQuery(esc(p.q.GetComments())).WillReturnRows(&sqlmock.Rows{})
	r := &empty.Empty{}
	s := &mockCms_GetCommentsServer{}

	_ = p.GetComments(r, s)

	checkExpectations(t)
}

func TestGetPosts(t *testing.T) {
	mock.ExpectQuery(esc(p.q.GetPosts())).WillReturnRows(&sqlmock.Rows{})
	r := &empty.Empty{}
	s := &mockCms_GetPostsServer{}

	_ = p.GetPosts(r, s)

	checkExpectations(t)
}

func TestGetUserComments(t *testing.T) {
	r := &pb.UserRequest{Id: "id"}
	mock.ExpectQuery(p.q.GetUserComments(r)).WillReturnRows(&sqlmock.Rows{})
	s := &mockCms_GetUserCommentsServer{}

	_ = p.GetUserComments(r, s)

	checkExpectations(t)
}
