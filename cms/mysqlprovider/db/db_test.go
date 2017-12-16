// +build integration

package db_test

import (
	"io/ioutil"
	"testing"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/thurt/demo-blog-platform/cms/mysqlprovider"
	"github.com/thurt/demo-blog-platform/cms/mysqlprovider_internal"
	pb "github.com/thurt/demo-blog-platform/cms/proto"

	"github.com/google/gofuzz"
	"github.com/xwb1989/sqlparser"
)

func checkSyntax(sql string, t *testing.T, strict bool) {
	var err error
	if strict {
		_, err = sqlparser.ParseStrictDDL(sql)
	} else {
		_, err = sqlparser.Parse(sql)
	}

	if err != nil {
		t.Error(t.Name(), err, sql)
	}
}

func TestCreateTablesSQLSyntax(t *testing.T) {
	d, err := ioutil.ReadFile("./docker-entrypoint-initdb.d/1_createTables.sql")
	if err != nil {
		panic(err)
	}

	sql := string(d)

	var stmt string
	for sql != "\n" {
		stmt, sql, err = sqlparser.SplitStatement(sql)
		if err != nil {
			panic(err)
		}

		// NOTE: using strict false here b/c the parser isnt able to parse FOREIGN KEY https://github.com/xwb1989/sqlparser/issues/26
		checkSyntax(stmt, t, false)
	}
}

func TestRequestSQLSyntax(t *testing.T) {
	q := &mysqlprovider.SqlQuery{}
	q_internal := mysqlprovider_internal.SqlQuery{}
	f := fuzz.New()

	ccr := &pb.CreateCommentRequest{}
	f.Fuzz(ccr)
	cpr := &pb.CreatePostRequest{}
	f.Fuzz(cpr)
	cur := &pb.CreateUserRequest{}
	f.Fuzz(cur)
	cr := &pb.CommentRequest{}
	f.Fuzz(cr)
	pr := &pb.PostRequest{}
	f.Fuzz(pr)
	ur := &pb.UserRequest{}
	f.Fuzz(ur)
	ucr := &pb.UpdateCommentRequest{}
	f.Fuzz(ucr)
	upr := &pb.UpdatePostRequest{}
	f.Fuzz(upr)

	checkSyntax(q_internal.GetUserPassword(), t, true)
	checkSyntax(q.CreateComment(ccr), t, true)
	checkSyntax(q.CreatePost(cpr), t, true)
	checkSyntax(q.CreateUser(cur), t, true)
	checkSyntax(q.DeleteComment(cr), t, true)
	checkSyntax(q.DeletePost(pr), t, true)
	checkSyntax(q.DeleteUser(ur), t, true)
	checkSyntax(q.GetComment(cr), t, true)
	checkSyntax(q.GetComments(), t, true)
	checkSyntax(q.GetPost(pr), t, true)
	checkSyntax(q.GetPosts(&empty.Empty{}), t, true)
	checkSyntax(q.GetPostComments(pr), t, true)
	checkSyntax(q.GetUser(ur), t, true)
	checkSyntax(q.GetUserComments(ur), t, true)
	checkSyntax(q.PublishPost(pr), t, true)
	checkSyntax(q.UnPublishPost(pr), t, true)
	checkSyntax(q.UpdateComment(ucr), t, true)
	checkSyntax(q.UpdatePost(upr), t, true)
}
