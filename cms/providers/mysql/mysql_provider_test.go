package mysql_provider

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
	p    *provider
)

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags

	var err error
	db, mock, err = sqlmock.New()
	if err != nil {
		panic(fmt.Errorf("an error '%s' was not expected when opening a stub database connection", err))
	}

	p = NewProvider(db)

	os.Exit(m.Run())
}

func TestGetPost(t *testing.T) {
	rs := sqlmock.NewRows([]string{"", "", "", "", "", ""}).AddRow(0, 0, 0, 0, 0, 0)
	mock.ExpectQuery("SELECT .+ FROM posts WHERE id=?").WithArgs("id").WillReturnRows(rs)
	r := &pb.PostRequest{Id: "id"}

	_, err := p.GetPost(context.Background(), r)
	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}
}
