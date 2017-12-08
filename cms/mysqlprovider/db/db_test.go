package db_test

import (
	"io/ioutil"
	"testing"

	"github.com/thurt/demo-blog-platform/cms/mysqlprovider"
	"github.com/thurt/demo-blog-platform/cms/mysqlprovider_internal"
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
		t.Error(err, sql)
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
	q := mysqlprovider.NewSqlQuery()
	q_internal := mysqlprovider_internal.NewSqlQuery()

	checkSyntax(q_internal.GetUserPassword(), t, true)
	checkSyntax(q.CreateComment(), t, true)
	checkSyntax(q.CreatePost(), t, true)
	checkSyntax(q.CreateUser(), t, true)
	checkSyntax(q.DeleteComment(), t, true)
	checkSyntax(q.DeletePost(), t, true)
	checkSyntax(q.DeleteUser(), t, true)
	checkSyntax(q.GetComment(), t, true)
	checkSyntax(q.GetComments(), t, true)
	checkSyntax(q.GetPost(), t, true)
	checkSyntax(q.GetPosts(), t, true)
	checkSyntax(q.GetPostComments(), t, true)
	checkSyntax(q.GetUser(), t, true)
	checkSyntax(q.GetUserComments(), t, true)
	checkSyntax(q.PublishPost(), t, true)
	checkSyntax(q.UnPublishPost(), t, true)
	checkSyntax(q.UpdateComment(), t, true)
	checkSyntax(q.UpdatePost(), t, true)
}
