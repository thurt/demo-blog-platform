package mysqlprovider_internal

import (
	"database/sql"
	"fmt"
	"os"
	"regexp"
	"testing"

	pb "github.com/thurt/demo-blog-platform/cms/proto"
	"golang.org/x/net/context"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var (
	db   *sql.DB
	mock sqlmock.Sqlmock
	p    *provider_internal
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
	p = New(db).(*provider_internal)

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

func TestGetUserPassword(t *testing.T) {
	r := &pb.UserRequest{"id"}
	mock.ExpectQuery(esc(p.q.GetUserPassword())).WithArgs(r.GetId())

	_, _ = p.GetUserPassword(context.Background(), r)

	checkExpectations(t)
}
