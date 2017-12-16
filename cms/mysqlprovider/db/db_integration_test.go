// +build integration

package dbtest

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/gofuzz"
	"github.com/thurt/demo-blog-platform/cms/domain"
	"github.com/thurt/demo-blog-platform/cms/mysqlprovider"
	pb "github.com/thurt/demo-blog-platform/cms/proto"
	dockertest "gopkg.in/ory-am/dockertest.v3"
)

var (
	db *sql.DB
	q  *mysqlprovider.SqlQuery
	p  domain.Provider
	f  *fuzz.Fuzzer
)
var TCP_PROXY = flag.String("TCP_PROXY", "localhost", "(optional) supply an IP address which this process can use to connect to the docker container it creates for integration testing. This flag will only be useful if you are using the docker unix port of a remote machine other than localhost")

func TestMain(m *testing.M) {
	flag.Parse()

	// create a new fuzzer
	f = fuzz.New()

	// create SqlQuery
	q = &mysqlprovider.SqlQuery{}
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	// increase maxwait time to 2 minutes (default is 1 min atm)
	pool.MaxWait = 2 * time.Minute

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
		db, err = sql.Open("mysql", fmt.Sprintf("root:secret@tcp(%s:%s)/cms", *TCP_PROXY, resource.GetPort("3306/tcp")))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// create Domain Provider
	p = mysqlprovider.New(db)

	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestGetComment(t *testing.T) {
	stubIn := &pb.CommentRequest{}
	f.Fuzz(stubIn)
	stmt := q.GetComment(stubIn)

	t.Run("must be a specific error when entity does not exist", func(t *testing.T) {
		err := db.QueryRow(stmt).Scan()

		if err != sql.ErrNoRows {
			t.Fail()
		}
	})
}

func TestGetUser(t *testing.T) {
	stubIn := &pb.UserRequest{}
	f.Fuzz(stubIn)
	stmt := q.GetUser(stubIn)

	t.Run("must be a specific error when entity does not exist", func(t *testing.T) {
		err := db.QueryRow(stmt).Scan()

		if err != sql.ErrNoRows {
			t.Fail()
		}
	})
}

func TestCRUD_Post(t *testing.T) {
	var PostId *pb.PostRequest

	t.Run("must answer with specified error when getting entity that does not exist", func(t *testing.T) {
		stubIn := &pb.PostRequest{}
		f.Fuzz(stubIn)
		_, err := p.GetPost(context.Background(), stubIn)

		if err != sql.ErrNoRows {
			t.Fail()
		}
	})
	t.Run("must answer without error when creating entity", func(t *testing.T) {
		stubIn := &pb.CreatePostRequest{}
		f.Fuzz(stubIn)

		pid, err := p.CreatePost(context.Background(), stubIn)
		if err != nil {
			t.Error("unexpected error:", err.Error())
		}

		PostId = pid
	})
	t.Run("must answer without error when getting entity that exists", func(t *testing.T) {
		_, err := p.GetPost(context.Background(), PostId)
		if err != nil {
			t.Error("uexpected error:", err.Error())
		}
	})
	t.Run("must answer without error when updating entity that exists", func(t *testing.T) {
		stubIn := &pb.UpdatePostRequest{}
		f.Fuzz(stubIn)
		stubIn.Id = PostId.GetId()

		_, err := p.UpdatePost(context.Background(), stubIn)

		if err != nil {
			t.Error("unexpected error:", err.Error())
		}
	})
	t.Run("must answer without error when deleting entity that exists", func(t *testing.T) {
		_, err := p.DeletePost(context.Background(), PostId)
		if err != nil {
			t.Error("unexpected error:", err.Error())
		}
	})
}
