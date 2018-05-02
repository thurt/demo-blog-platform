// +build integration

package dbtest

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/github"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/gofuzz"
	_ "github.com/lib/pq"
	"github.com/thurt/demo-blog-platform/cms/domain"
	"github.com/thurt/demo-blog-platform/cms/mysqlprovider"
	pb "github.com/thurt/demo-blog-platform/cms/proto"
	dockertest "gopkg.in/ory-am/dockertest.v3"
)

const (
	DB_SCHEMA_VERSION = 2
)

var (
	db *sql.DB
	p  domain.Provider
	f  *fuzz.Fuzzer
)

func TestMain(m *testing.M) {
	// (optional) supply an IP address which this process can use to connect to the docker container it creates for integration testing. This flag will only be useful if you are using the docker unix port of a remote machine other than localhost
	TCP_PROXY := os.Getenv("TCP_PROXY")
	if TCP_PROXY == "" {
		TCP_PROXY = "localhost"
	}

	// create a new fuzzer
	f = fuzz.New()

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
		Repository: "demo-blog-platform_db",
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
		db, err = sql.Open("mysql", fmt.Sprintf("root:secret@tcp(%s:%s)/cms?multiStatements=true", TCP_PROXY, resource.GetPort("3306/tcp")))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		if purgeErr := pool.Purge(resource); purgeErr != nil {
			log.Println("Could not purge resource (you may have to manually):", purgeErr)
		}
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// setup schema version in database
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Println("Error setting up mysql driver for migrations")
		panic(err.Error())
	}
	mi, err := migrate.NewWithDatabaseInstance(
		"github://:@thurt/demo-blog-platform/cms/mysqlprovider/db/migrations",
		"mysql", driver)
	if err != nil {
		log.Println("Error setting up database migrate instance")
		panic(err.Error())
	}

	err = mi.Migrate(DB_SCHEMA_VERSION)
	if err != nil {
		log.Printf("Error performing database migration to version %d. Your database is probably dirty now and requires manual adjustment.", DB_SCHEMA_VERSION)
		panic(err.Error())
	}
	log.Printf("Successfully performed database migration to version %d", DB_SCHEMA_VERSION)

	// create Domain Provider
	p = mysqlprovider.New(db)

	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestCRUD_Post(t *testing.T) {
	var PostId *pb.PostRequest

	t.Run("must answer with zero-value Post when getting entity that does not exist", func(t *testing.T) {
		stubIn := &pb.PostRequest{}
		f.Fuzz(stubIn)
		p, err := p.GetPost(context.Background(), stubIn)

		if err != nil {
			t.Error("unexpected error:", err.Error())
		}

		if *p != (pb.Post{}) {
			t.Errorf("expected a zero-value Post, instead got %+v", p)
		}

	})
	t.Run("must answer without error when creating entity", func(t *testing.T) {
		stubIn := &pb.CreatePostWithSlug{Post: &pb.CreatePostRequest{}}
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
		stubIn := &pb.UpdatePostWithSlug{Post: &pb.UpdatePostRequest{}}
		f.Fuzz(stubIn)
		stubIn.Post.Id = PostId.GetId()

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

func Test_AdminExists(t *testing.T) {
	t.Run("must answer without error when checking for entity that does not exist", func(t *testing.T) {
		stubIn := &empty.Empty{}
		f.Fuzz(stubIn)
		_, err := p.AdminExists(context.Background(), stubIn)

		if err != nil {
			t.Error("unexpected error:", err.Error())
		}
	})
	t.Run("must answer without error when checking for entity that does exist", func(t *testing.T) {
		stubIn := &empty.Empty{}
		f.Fuzz(stubIn)

		// prerequisite: must first create an admin
		preStubIn := &pb.CreateUserWithRole{User: &pb.CreateUserRequest{}}
		f.Fuzz(preStubIn)
		preStubIn.Role = pb.UserRole_ADMIN
		_, err := p.CreateNewUser(context.Background(), preStubIn)

		if err != nil {
			t.Error("unexpected error in prerequisite:", err.Error())
		}

		// continue with actual test
		_, err = p.AdminExists(context.Background(), stubIn)

		if err != nil {
			t.Error("unexpected error:", err.Error())
		}
	})
}
