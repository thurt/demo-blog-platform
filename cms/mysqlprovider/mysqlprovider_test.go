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
	"github.com/thurt/demo-blog-platform/cms/mock_proto"
	pb "github.com/thurt/demo-blog-platform/cms/proto"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var (
	db   *sql.DB
	mock sqlmock.Sqlmock
	p    *provider
	f    *fuzz.Fuzzer
)

var regexAny string = ".*"

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
	t.Run("requires returning zero-value struct when sql response is a sql.ErrNoRows", func(t *testing.T) {
		mock.ExpectQuery(regexAny).WillReturnError(sql.ErrNoRows)

		post, err := p.GetPost(context.Background(), stubIn)
		if err != nil {
			t.Error("unexpected error:", err.Error())
		}

		if *post != (pb.Post{}) {
			t.Error("expected a zero-value struct")
		}
	})
}

func TestGetPostBySlug(t *testing.T) {
	stubIn := &pb.PostBySlugRequest{}
	stubOut := &pb.Post{}
	f.Fuzz(stubIn)
	f.Fuzz(stubOut)
	regexSql := esc(p.q.GetPostBySlug(stubIn))
	stubRows := sqlmock.NewRows(structs.Names(stubOut))
	stubRows.AddRow(makeRowData(structs.Values(stubOut))...)

	t.Run("requires dispatching the correct sql request", func(t *testing.T) {
		mock.ExpectQuery(regexSql)

		_, _ = p.GetPostBySlug(context.Background(), stubIn)

		checkExpectations(t)
	})
	t.Run("requires returning result with correct values from sql response", func(t *testing.T) {
		mock.ExpectQuery(regexAny).WillReturnRows(stubRows)

		result, err := p.GetPostBySlug(context.Background(), stubIn)
		if err != nil {
			t.Error("unexpected error:", err.Error())
		}

		if !reflect.DeepEqual(result, stubOut) {
			t.Error("result should have same values as stub values")
		}
	})
	t.Run("requires returning error when sql response is an error", func(t *testing.T) {
		mock.ExpectQuery(regexAny).WillReturnError(errors.New(""))

		_, err := p.GetPostBySlug(context.Background(), stubIn)
		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires returning zero-value struct when sql response is a sql.ErrNoRows", func(t *testing.T) {
		mock.ExpectQuery(regexAny).WillReturnError(sql.ErrNoRows)

		post, err := p.GetPostBySlug(context.Background(), stubIn)
		if err != nil {
			t.Error("unexpected error:", err.Error())
		}

		if *post != (pb.Post{}) {
			t.Error("expected a zero-value struct")
		}
	})
}

func TestGetUnpublishedPost(t *testing.T) {
	stubIn := &pb.PostRequest{}
	stubOut := &pb.UnpublishedPost{Post: &pb.Post{}}
	f.Fuzz(stubIn)
	f.Fuzz(stubOut)
	regexSql := esc(p.q.GetUnpublishedPost(stubIn))
	stubRows := sqlmock.NewRows(append(structs.Names(stubOut.Post), "Published", "LastPublished"))
	stubRows.AddRow(makeRowData(structs.Values(stubOut))...)

	t.Run("requires dispatching the correct sql request", func(t *testing.T) {
		mock.ExpectQuery(regexSql)

		_, _ = p.GetUnpublishedPost(context.Background(), stubIn)

		checkExpectations(t)
	})
	t.Run("requires returning result with correct values from sql response", func(t *testing.T) {
		mock.ExpectQuery(regexAny).WillReturnRows(stubRows)

		result, err := p.GetUnpublishedPost(context.Background(), stubIn)
		if err != nil {
			t.Error("unexpected error:", err.Error())
		}

		if !reflect.DeepEqual(result, stubOut) {
			t.Error("result should have same values as stub values")
		}
	})
	t.Run("requires returning error when sql response is an error", func(t *testing.T) {
		mock.ExpectQuery(regexAny).WillReturnError(errors.New(""))

		_, err := p.GetUnpublishedPost(context.Background(), stubIn)
		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires returning zero-value struct when sql response is a sql.ErrNoRows", func(t *testing.T) {
		mock.ExpectQuery(regexAny).WillReturnError(sql.ErrNoRows)

		upost, err := p.GetUnpublishedPost(context.Background(), stubIn)
		if err != nil {
			t.Error("unexpected error:", err.Error())
		}

		if *upost != (pb.UnpublishedPost{}) && *upost.Post != (pb.Post{}) {
			t.Error("expected a zero-value struct")
		}
	})
}

func TestGetUnpublishedPostBySlug(t *testing.T) {
	stubIn := &pb.PostBySlugRequest{}
	stubOut := &pb.UnpublishedPost{Post: &pb.Post{}}
	f.Fuzz(stubIn)
	f.Fuzz(stubOut)
	regexSql := esc(p.q.GetUnpublishedPostBySlug(stubIn))
	stubRows := sqlmock.NewRows(append(structs.Names(stubOut.Post), "Published", "LastPublished"))
	stubRows.AddRow(makeRowData(structs.Values(stubOut))...)

	t.Run("requires dispatching the correct sql request", func(t *testing.T) {
		mock.ExpectQuery(regexSql)

		_, _ = p.GetUnpublishedPostBySlug(context.Background(), stubIn)

		checkExpectations(t)
	})
	t.Run("requires returning result with correct values from sql response", func(t *testing.T) {
		mock.ExpectQuery(regexAny).WillReturnRows(stubRows)

		result, err := p.GetUnpublishedPostBySlug(context.Background(), stubIn)
		if err != nil {
			t.Error("unexpected error:", err.Error())
		}

		if !reflect.DeepEqual(result, stubOut) {
			t.Error("result should have same values as stub values")
		}
	})
	t.Run("requires returning error when sql response is an error", func(t *testing.T) {
		mock.ExpectQuery(regexAny).WillReturnError(errors.New(""))

		_, err := p.GetUnpublishedPostBySlug(context.Background(), stubIn)
		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires returning zero-value struct when sql response is a sql.ErrNoRows", func(t *testing.T) {
		mock.ExpectQuery(regexAny).WillReturnError(sql.ErrNoRows)

		upost, err := p.GetUnpublishedPostBySlug(context.Background(), stubIn)
		if err != nil {
			t.Error("unexpected error:", err.Error())
		}

		if *upost != (pb.UnpublishedPost{}) && *upost.Post != (pb.Post{}) {
			t.Error("expected a zero-value struct")
		}
	})
}

func TestGetUser(t *testing.T) {
	stubIn := &pb.UserRequest{}
	stubOut := &pb.User{}
	f.Fuzz(stubIn)
	f.Fuzz(stubOut)
	regexSql := esc(p.q.GetUser(stubIn))
	stubRows := sqlmock.NewRows(structs.Names(stubOut))
	stubRows.AddRow(makeRowData(structs.Values(stubOut))...)

	t.Run("requires dispatching the correct sql request", func(t *testing.T) {
		mock.ExpectQuery(regexSql)

		_, _ = p.GetUser(context.Background(), stubIn)

		checkExpectations(t)
	})
	t.Run("requires returning result with correct values from sql response", func(t *testing.T) {
		mock.ExpectQuery(regexAny).WillReturnRows(stubRows)

		result, err := p.GetUser(context.Background(), stubIn)
		if err != nil {
			t.Error("unexpected error:", err.Error())
		}

		if !reflect.DeepEqual(result, stubOut) {
			t.Error("result should have same values as stub values")
		}
	})
	t.Run("requires returning error when sql response is an error", func(t *testing.T) {
		mock.ExpectQuery(regexAny).WillReturnError(errors.New(""))

		_, err := p.GetUser(context.Background(), stubIn)
		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires returning zero-value struct when sql response is a sql.ErrNoRows", func(t *testing.T) {
		mock.ExpectQuery(regexAny).WillReturnError(sql.ErrNoRows)

		user, err := p.GetUser(context.Background(), stubIn)
		if err != nil {
			t.Error("unexpected error:", err.Error())
		}

		if *user != (pb.User{}) {
			t.Error("expected a zero-value struct")
		}
	})
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
	stubIn := &pb.CreatePostWithSlug{Post: &pb.CreatePostRequest{}}
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

func TestGetComment(t *testing.T) {
	stubIn := &pb.CommentRequest{}
	stubOut := &pb.Comment{}
	f.Fuzz(stubIn)
	f.Fuzz(stubOut)
	regexSql := esc(p.q.GetComment(stubIn))
	stubRows := sqlmock.NewRows(structs.Names(stubOut))
	stubRows.AddRow(makeRowData(structs.Values(stubOut))...)

	t.Run("requires dispatching the correct sql request", func(t *testing.T) {
		mock.ExpectQuery(regexSql)

		_, _ = p.GetComment(context.Background(), stubIn)

		checkExpectations(t)
	})
	t.Run("requires returning result with correct values from sql response", func(t *testing.T) {
		mock.ExpectQuery(regexAny).WillReturnRows(stubRows)

		result, err := p.GetComment(context.Background(), stubIn)
		if err != nil {
			t.Error("unexpected error:", err.Error())
		}

		if !reflect.DeepEqual(result, stubOut) {
			t.Error("result should have same values as stub values")
		}
	})
	t.Run("requires returning error when sql response is an error", func(t *testing.T) {
		mock.ExpectQuery(regexAny).WillReturnError(errors.New(""))

		_, err := p.GetComment(context.Background(), stubIn)
		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires returning zero-value struct when sql response is a sql.ErrNoRows", func(t *testing.T) {
		mock.ExpectQuery(regexAny).WillReturnError(sql.ErrNoRows)

		comment, err := p.GetComment(context.Background(), stubIn)
		if err != nil {
			t.Error("unexpected error:", err.Error())
		}

		if *comment != (pb.Comment{}) {
			t.Error("expected a zero-value struct")
		}
	})
}

func TestUpdateComment(t *testing.T) {
	r := &pb.UpdateCommentRequest{Content: "content", Id: 0}
	mock.ExpectExec(p.q.UpdateComment(r)).WillReturnResult(sqlmock.NewResult(1, 1))

	_, _ = p.UpdateComment(context.Background(), r)

	checkExpectations(t)
}

func TestUpdatePost(t *testing.T) {
	r := &pb.UpdatePostWithSlug{Post: &pb.UpdatePostRequest{Title: "A Great Title!", Content: "content", Id: 0}, Slug: "a-great-title"}
	mock.ExpectExec(p.q.UpdatePost(r)).WillReturnResult(sqlmock.NewResult(1, 1))

	_, _ = p.UpdatePost(context.Background(), r)

	checkExpectations(t)
}

func TestGetPostComments(t *testing.T) {
	r := &pb.PostRequest{Id: 0}
	mock.ExpectQuery(p.q.GetPostComments(r)).WillReturnRows(&sqlmock.Rows{})
	s := &mock_proto.MockCms_GetPostCommentsServer{}

	_ = p.GetPostComments(r, s)

	checkExpectations(t)
}

func TestGetComments(t *testing.T) {
	r := &empty.Empty{}
	mock.ExpectQuery(esc(p.q.GetComments(r))).WillReturnRows(&sqlmock.Rows{})
	s := &mock_proto.MockCms_GetCommentsServer{}

	_ = p.GetComments(r, s)

	checkExpectations(t)
}

func TestGetPosts(t *testing.T) {
	stubIn := &empty.Empty{}
	stubOut := []*pb.Post{&pb.Post{}, &pb.Post{}}
	f.Fuzz(stubIn)
	f.Fuzz(stubOut[0])
	f.Fuzz(stubOut[1])
	mockStreamOut := mock_proto.NewMockCms_GetPostsServer()

	t.Run("requires dispatching the correct sql request", func(t *testing.T) {
		regexSql := esc(p.q.GetPosts())
		mock.ExpectQuery(regexSql)

		_ = p.GetPosts(stubIn, mockStreamOut)

		checkExpectations(t)
	})
	t.Run("requires returning result thru stream with correct values from sql response", func(t *testing.T) {
		stubRows := sqlmock.NewRows(structs.Names(stubOut[0])).AddRow(makeRowData(structs.Values(stubOut[0]))...).AddRow(makeRowData(structs.Values(stubOut[1]))...)
		mock.ExpectQuery(regexAny).WillReturnRows(stubRows)

		err := p.GetPosts(stubIn, mockStreamOut)
		if err != nil {
			t.Error("unexpected error:", err.Error())
		}

		if !reflect.DeepEqual(mockStreamOut.Results, stubOut) {
			t.Error(fmt.Sprintf("result should have same values as stub values:\nHave:\n%v\nWant:\n%v\n", mockStreamOut.Results, stubOut))
		}
	})
	t.Run("requires returning error when sql response is an error", func(t *testing.T) {
		mock.ExpectQuery(regexAny).WillReturnError(errors.New(""))

		err := p.GetPosts(stubIn, mockStreamOut)
		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires returning error when stream.Send creates an error", func(t *testing.T) {
		mockStreamOutWithErr := mock_proto.NewMockCms_GetPostsServer().SetSendError(1, errors.New(""))
		stubRows := sqlmock.NewRows(structs.Names(stubOut[0])).AddRow(makeRowData(structs.Values(stubOut[0]))...).AddRow(makeRowData(structs.Values(stubOut[1]))...)
		mock.ExpectQuery(regexAny).WillReturnRows(stubRows)

		err := p.GetPosts(stubIn, mockStreamOutWithErr)
		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires returning error when driver scan creates an error", func(t *testing.T) {
		// creates badStub with 0 fields in order to setup a situation where *Rows.Scan (called by the func under test) should return an error: "sql: expected 0 destination arguments in Scan, got X"
		type badStub struct{}
		badStubRows := sqlmock.NewRows(structs.Names(badStub{})).AddRow(makeRowData(structs.Values(badStub{}))...)

		mock.ExpectQuery(regexAny).WillReturnRows(badStubRows)

		err := p.GetPosts(stubIn, mockStreamOut)
		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires returning error when driver row creates an error", func(t *testing.T) {
		stubRowsWithErr := sqlmock.NewRows(structs.Names(stubOut[0])).AddRow(makeRowData(structs.Values(stubOut[0]))...).RowError(0, errors.New(""))
		mock.ExpectQuery(regexAny).WillReturnRows(stubRowsWithErr)

		err := p.GetPosts(stubIn, mockStreamOut)
		if err == nil {
			t.Error("expected an error")
		}
	})
}

func TestGetUnpublishedPosts(t *testing.T) {
	stubIn := &empty.Empty{}
	stubOut := []*pb.UnpublishedPost{&pb.UnpublishedPost{Post: &pb.Post{}}, &pb.UnpublishedPost{Post: &pb.Post{}}}
	f.Fuzz(stubIn)
	f.Fuzz(stubOut[0])
	f.Fuzz(stubOut[1])
	mockStreamOut := mock_proto.NewMockCms_GetUnpublishedPostsServer()
	columns := append(structs.Names(stubOut[0].Post), "Published", "LastPublished")
	rowValues1 := makeRowData(structs.Values(stubOut[0]))
	rowValues2 := makeRowData(structs.Values(stubOut[1]))

	t.Run("requires dispatching the correct sql request", func(t *testing.T) {
		regexSql := esc(p.q.GetUnpublishedPosts())
		mock.ExpectQuery(regexSql)

		_ = p.GetUnpublishedPosts(stubIn, mockStreamOut)

		checkExpectations(t)
	})
	t.Run("requires returning result thru stream with correct values from sql response", func(t *testing.T) {
		stubRows := sqlmock.NewRows(columns).AddRow(rowValues1...).AddRow(rowValues2...)
		mock.ExpectQuery(regexAny).WillReturnRows(stubRows)

		err := p.GetUnpublishedPosts(stubIn, mockStreamOut)
		if err != nil {
			t.Error("unexpected error:", err.Error())
		}

		if !reflect.DeepEqual(mockStreamOut.Results, stubOut) {
			t.Error(fmt.Sprintf("result should have same values as stub values:\nHave:\n%v\nWant:\n%v\n", mockStreamOut.Results, stubOut))
		}
	})
	t.Run("requires returning error when sql response is an error", func(t *testing.T) {
		mock.ExpectQuery(regexAny).WillReturnError(errors.New(""))

		err := p.GetUnpublishedPosts(stubIn, mockStreamOut)
		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires returning error when stream.Send creates an error", func(t *testing.T) {
		mockStreamOutWithErr := mock_proto.NewMockCms_GetUnpublishedPostsServer().SetSendError(1, errors.New(""))
		stubRows := sqlmock.NewRows(columns).AddRow(rowValues1...).AddRow(rowValues2...)
		mock.ExpectQuery(regexAny).WillReturnRows(stubRows)

		err := p.GetUnpublishedPosts(stubIn, mockStreamOutWithErr)
		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires returning error when driver scan creates an error", func(t *testing.T) {
		// creates badStub with 0 fields in order to setup a situation where *Rows.Scan (called by the func under test) should return an error: "sql: expected 0 destination arguments in Scan, got X"
		type badStub struct{}
		badStubRows := sqlmock.NewRows(structs.Names(badStub{})).AddRow(makeRowData(structs.Values(badStub{}))...)

		mock.ExpectQuery(regexAny).WillReturnRows(badStubRows)

		err := p.GetUnpublishedPosts(stubIn, mockStreamOut)
		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires returning error when driver row creates an error", func(t *testing.T) {
		stubRowsWithErr := sqlmock.NewRows(columns).AddRow(rowValues1...).RowError(0, errors.New(""))
		mock.ExpectQuery(regexAny).WillReturnRows(stubRowsWithErr)

		err := p.GetUnpublishedPosts(stubIn, mockStreamOut)
		if err == nil {
			t.Error("expected an error")
		}
	})
}

func TestGetUserComments(t *testing.T) {
	r := &pb.UserRequest{Id: "id"}
	mock.ExpectQuery(p.q.GetUserComments(r)).WillReturnRows(&sqlmock.Rows{})
	s := &mock_proto.MockCms_GetUserCommentsServer{}

	_ = p.GetUserComments(r, s)

	checkExpectations(t)
}

func TestAdminExists(t *testing.T) {
	stubIn := &empty.Empty{}
	stubRows := sqlmock.NewRows([]string{"EXISTS(query)"})

	t.Run("requires sending the correct sql request", func(t *testing.T) {
		regexSql := esc(p.q.AdminExists(stubIn))
		mock.ExpectQuery(regexSql)

		_, _ = p.AdminExists(context.Background(), stubIn)

		checkExpectations(t)
	})
	t.Run("requires returning error when sql response is an error", func(t *testing.T) {
		mock.ExpectQuery(regexAny).WillReturnError(errors.New(""))

		_, err := p.AdminExists(context.Background(), stubIn)

		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires returning true when sql response is 1", func(t *testing.T) {
		stubRows.AddRow(1)
		mock.ExpectQuery(regexAny).WillReturnRows(stubRows)

		res, err := p.AdminExists(context.Background(), stubIn)

		if err != nil {
			t.Error("unexpected error:", err.Error())
		}

		wantValue := true
		if res.Value != wantValue {
			t.Errorf("returned wrong value. want %v, got %v", wantValue, res.Value)
		}
	})
	t.Run("requires returning false when sql response is 0", func(t *testing.T) {
		stubRows.AddRow(0)
		mock.ExpectQuery(regexAny).WillReturnRows(stubRows)

		res, err := p.AdminExists(context.Background(), stubIn)

		if err != nil {
			t.Error("unexpected error:", err.Error())
		}

		wantValue := false
		if res.Value != wantValue {
			t.Errorf("returned wrong value. want %v, got %v", wantValue, res.Value)
		}
	})
}

func TestUpdateUserLastActive(t *testing.T) {
	stubIn := &pb.UserRequest{}
	stubOut := &empty.Empty{}
	f.Fuzz(stubIn)
	stubResult := sqlmock.NewResult(0, 1)

	t.Run("requires sending the correct sql request", func(t *testing.T) {
		regexSql := esc(p.q.UpdateUserLastActive(stubIn))
		mock.ExpectExec(regexSql)

		_, _ = p.UpdateUserLastActive(context.Background(), stubIn)

		checkExpectations(t)

	})
	t.Run("requires returning error when sql response is an error", func(t *testing.T) {
		mock.ExpectExec(regexAny).WillReturnError(errors.New(""))

		_, err := p.UpdateUserLastActive(context.Background(), stubIn)
		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires returning result with correct values from sql response", func(t *testing.T) {
		mock.ExpectExec(regexAny).WillReturnResult(stubResult)

		result, err := p.UpdateUserLastActive(context.Background(), stubIn)
		if err != nil {
			t.Error("unexpected error:", err.Error())
		}

		if !reflect.DeepEqual(result, stubOut) {
			t.Error("result should have same values as stub values")
		}
	})
}

func TestUpdateUnpublishedPost(t *testing.T) {
	stubIn := &pb.UpdatePostWithSlug{Post: &pb.UpdatePostRequest{}}
	stubOut := &empty.Empty{}
	f.Fuzz(stubIn)
	stubResult := sqlmock.NewResult(0, 1)

	t.Run("requires sending the correct sql query", func(t *testing.T) {
		regexSql := esc(p.q.UpdateUnpublishedPost(stubIn))
		mock.ExpectExec(regexSql)

		_, _ = p.UpdateUnpublishedPost(context.Background(), stubIn)

		checkExpectations(t)

	})
	t.Run("requires returning error when sql response is an error", func(t *testing.T) {
		mock.ExpectExec(regexAny).WillReturnError(errors.New(""))

		_, err := p.UpdateUnpublishedPost(context.Background(), stubIn)
		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires returning result with correct values from sql response", func(t *testing.T) {
		mock.ExpectExec(regexAny).WillReturnResult(stubResult)

		result, err := p.UpdateUnpublishedPost(context.Background(), stubIn)
		if err != nil {
			t.Error("unexpected error:", err.Error())
		}

		if !reflect.DeepEqual(result, stubOut) {
			t.Error("result should have same values as stub values")
		}
	})
}

func TestCreateNewUser(t *testing.T) {
	cur := &pb.CreateUserRequest{}
	f.Fuzz(cur)
	stubIn := &pb.CreateUserWithRole{User: cur}
	stubOut := &pb.UserRequest{}
	f.Fuzz(stubIn)
	f.Fuzz(stubOut)
	stubOut.Id = stubIn.GetUser().GetId()
	stubResult := sqlmock.NewResult(1, 1)

	t.Run("requires sending the correct sql request", func(t *testing.T) {
		regexSql := esc(p.q.CreateUser(stubIn))
		mock.ExpectExec(regexSql)

		_, _ = p.CreateNewUser(context.Background(), stubIn)

		checkExpectations(t)
	})
	t.Run("requires returning error when sql response is an error", func(t *testing.T) {
		mock.ExpectExec(regexAny).WillReturnError(errors.New(""))
		_, err := p.CreateNewUser(context.Background(), stubIn)

		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("requires returning result with correct values when sql returns results", func(t *testing.T) {
		mock.ExpectExec(regexAny).WillReturnResult(stubResult)

		result, err := p.CreateNewUser(context.Background(), stubIn)
		if err != nil {
			t.Error("unexpected error:", err.Error())
		}

		if !reflect.DeepEqual(result, stubOut) {
			t.Error("result should have same values as stub values")
		}
	})
}
