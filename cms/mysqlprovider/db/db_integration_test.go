package db_test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/thurt/demo-blog-platform/cms/mysqlprovider"
	pb "github.com/thurt/demo-blog-platform/cms/proto"
	dockertest "gopkg.in/ory-am/dockertest.v3"
)

var db *sql.DB

func TestMain(m *testing.M) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	opts := &dockertest.RunOptions{
		Name:       "db-test",
		Repository: "demoblogplatform_db",
		Env:        []string{"MYSQL_ROOT_PASSWORD=secret"},
	}
	resource, err := pool.RunWithOptions(opts)
	//	resource, err := pool.Run("mysql", "5.7", []string{"MYSQL_ROOT_PASSWORD=secret"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error
		db, err = sql.Open("mysql", fmt.Sprintf("root:secret@tcp(172.17.0.3:3306)/cms"))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestGetPostAssertions(t *testing.T) {
	q := &mysqlprovider.SqlQuery{}
	t.Run("must be an error when entity does not exist", func(t *testing.T) {
		err := db.QueryRow(q.GetPost(&pb.PostRequest{Id: 1})).Scan()

		if err != sql.ErrNoRows {
			t.Fail()
		}
	})
}

func TestGetCommentAssertions(t *testing.T) {
	q := &mysqlprovider.SqlQuery{}
	t.Run("must be an error when entity does not exist", func(t *testing.T) {
		err := db.QueryRow(q.GetComment(&pb.CommentRequest{Id: 1})).Scan()

		if err != sql.ErrNoRows {
			t.Fail()
		}
	})
}

func TestGetUserAssertions(t *testing.T) {
	q := &mysqlprovider.SqlQuery{}
	t.Run("must be an error when entity does not exist", func(t *testing.T) {
		err := db.QueryRow(q.GetUser(&pb.UserRequest{Id: "id"})).Scan()

		if err != sql.ErrNoRows {
			t.Fail()
		}
	})
}
