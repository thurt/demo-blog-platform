package mysqlprovider

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

	pb "github.com/thurt/demo-blog-platform/cms/proto"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var (
	db   *sql.DB
	mock sqlmock.Sqlmock
	p    *pb.CmsServer
)

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags

	var err error
	// initialize the mock db and mock api
	db, mock, err = sqlmock.New()
	if err != nil {
		panic(fmt.Errorf("an error '%s' was not expected when opening a stub database connection", err))
	}

	// create the service provider
	p = New(db)

	os.Exit(m.Run())
}

// TestStatementsSyntax parses each SQL statement in statements to assert valid SQL syntax
//func TestStatementsSyntax(t *testing.T) {
//	for k, v := range statements {
//		_, err := sqlparser.Parse(v)
//		if err != nil {
//			t.Errorf("%s: %s", k, err)
//		}
//
//	}
//}

func TestGetPost(t *testing.T) {
	rs := sqlmock.NewRows([]string{"", "", "", "", "", ""}).AddRow(0, 0, 0, 0, 0, 0)
	mock.ExpectQuery("SELECT").WithArgs("id_value").WillReturnRows(rs)
	r := &pb.PostRequest{Id: "id_value"}

	_, err := p.GetPost(context.Background(), r)
	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}
}
