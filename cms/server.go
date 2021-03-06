package main // import "github.com/thurt/demo-blog-platform/cms"

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/validator"
	_ "github.com/lib/pq"
	"github.com/memcachier/mc"
	"github.com/thurt/demo-blog-platform/cms/authentication"
	"github.com/thurt/demo-blog-platform/cms/authorization"
	"github.com/thurt/demo-blog-platform/cms/cacher"
	"github.com/thurt/demo-blog-platform/cms/emailer"
	"github.com/thurt/demo-blog-platform/cms/hasher"
	"github.com/thurt/demo-blog-platform/cms/mysqlprovider"
	pb "github.com/thurt/demo-blog-platform/cms/proto"
	"github.com/thurt/demo-blog-platform/cms/usecases"
	trace "golang.org/x/net/trace"
	"google.golang.org/grpc"
)

const (
	PORT = 10000
)

var MYSQL_CONNECTION string
var DB_SCHEMA_VERSION uint

var (
	MEMCACHED_HOST     string
	MEMCACHED_USER     string
	MEMCACHED_PASSWORD string
)
var (
	SMTP_HOST     string
	SMTP_PORT     string
	SMTP_USER     string
	SMTP_PASSWORD string
)

func main() {
	//connect to memcache
	MEMCACHED_HOST = os.Getenv("MEMCACHED_HOST")
	mcCn, err := mc.Dial("tcp", MEMCACHED_HOST)
	if err != nil {
		panic(err)
	}
	defer mcCn.Close()

	MEMCACHED_USER = os.Getenv("MEMCACHED_USER")
	MEMCACHED_PASSWORD = os.Getenv("MEMCACHED_PASSWORD")
	err = mcCn.Auth(MEMCACHED_USER, MEMCACHED_PASSWORD)
	if err != nil {
		panic(err)
	}
	log.Println("Connected to memcache host")

	// connect to smtp mail
	SMTP_HOST = os.Getenv("SMTP_HOST")
	SMTP_PORT = os.Getenv("SMTP_PORT")
	SMTP_USER = os.Getenv("SMTP_USER")
	SMTP_PASSWORD = os.Getenv("SMTP_PASSWORD")
	SMTP_CONN := fmt.Sprintf("%s:%s", SMTP_HOST, SMTP_PORT)
	log.Println("Testing connection to smtp server " + SMTP_CONN)
	smtpCn, err := smtp.Dial(SMTP_CONN)
	if err != nil {
		log.Println("Couldn't connect to smtp server " + SMTP_CONN)
		panic(err.Error())
	}
	log.Println("Connected to smtp server " + SMTP_CONN)

	var smtpAuth smtp.Auth
	if SMTP_USER != "" {
		smtpAuth = smtp.PlainAuth("", SMTP_USER, SMTP_PASSWORD, SMTP_HOST)
		log.Println("Attempting to authorize with smtp server as user " + SMTP_USER)
		if ok, _ := smtpCn.Extension("STARTTLS"); ok {
			config := &tls.Config{ServerName: SMTP_HOST}
			if err = smtpCn.StartTLS(config); err != nil {
				panic(err.Error())
			}
			err = smtpCn.Auth(smtpAuth)
			if err != nil {
				log.Println("Couldn't authenticate with smtp host")
				panic(err.Error())
			}
			log.Println("Authenticated with smtp host as user " + SMTP_USER)
		}
	}
	log.Println("SMTP communication test completed, now disconnecting")

	smtpCn.Close()

	// connect to db
	MYSQL_CONNECTION = fmt.Sprint(os.Getenv("MYSQL_CONNECTION"), "?multiStatements=true")
	db, err := sql.Open("mysql", MYSQL_CONNECTION)
	if err != nil {
		log.Println("Couldn't connect with mysql connection string")
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Println("Couldn't ping database server")
		panic(err.Error())
	}
	log.Println("Connected to db server")
	defer db.Close()

	// maybe perform db migration
	// migration only happens when:
	//		env var DB_SCHEMA_VERSION exists, and
	//		its value differs from current version in the database
	vString := os.Getenv("DB_SCHEMA_VERSION")
	if vString == "" {
		log.Printf("Skipping database migration since no DB_SCHEMA_VERSION was provided")
	} else {
		vInt64, err := strconv.ParseUint(vString, 10, 0)
		if err != nil {
			log.Println("Failed to convert env var DB_SCHEMA_VERSION to integer")
			panic(err.Error())
		}
		DB_SCHEMA_VERSION = uint(vInt64)

		driver, err := mysql.WithInstance(db, &mysql.Config{})
		if err != nil {
			log.Println("Error setting up mysql driver for migrations")
			panic(err.Error())
		}
		m, err := migrate.NewWithDatabaseInstance(
			"file:///root/migrations/",
			"mysql", driver)
		if err != nil {
			log.Println("Error setting up database migrate instance")
			panic(err.Error())
		}

		// check the version
		currV, d, err := m.Version()
		if err != nil {
			// if there is no current version set, then set it to version 1
			if err == migrate.ErrNilVersion {
				err := m.Migrate(1)
				if err != nil {
					log.Printf("Error setting initial database version to 1")
					panic(err.Error())
				}
			} else {
				log.Printf("Error getting database migration version")
				panic(err.Error())
			}
		}
		// if dirty, panic
		if d == true {
			log.Printf("Aborting database migration checks because your database is dirty")
			panic(err.Error())
		}
		// if currV == provided version then skip, no-op
		if currV == DB_SCHEMA_VERSION {
			log.Printf("The provided version, %d, matches the currently installed version in the database", DB_SCHEMA_VERSION)
		} else {
			err = m.Migrate(DB_SCHEMA_VERSION)
			if err != nil {
				log.Printf("Error performing database migration from version %d to version %d. Your database is probably dirty now and requires manual adjustment.", currV, DB_SCHEMA_VERSION)
				panic(err.Error())
			}
			log.Printf("Successfully migrated database from version %d to version %d", currV, DB_SCHEMA_VERSION)
		}
	}

	// setup grpc server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", PORT))
	if err != nil {
		log.Println("failed to listen")
		panic(err.Error())
	}

	// create the cacher service using a memcached driver
	mcCacher := cacher.New(mcCn)

	authProvider, authFunc := authentication.New(mcCacher, 8*time.Hour)

	grpc.EnableTracing = true
	opts := []grpc.ServerOption{
		grpc.ConnectionTimeout(5 * time.Second),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_validator.UnaryServerInterceptor(),
			grpc_auth.UnaryServerInterceptor(authFunc),
		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_validator.StreamServerInterceptor(),
			grpc_auth.StreamServerInterceptor(authFunc),
		)),
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterCmsServer(grpcServer, authorization.New(usecases.New(mysqlprovider.New(db), authProvider, hasher.New(), emailer.New(SMTP_CONN, smtpAuth), mcCacher)))
	log.Printf("Started grpc server on port %d", PORT)

	// setup rest proxy server
	go func() {
		log.Println("Starting up rest-proxy")
		err = ProxyServe()
		if err != nil {
			log.Println("proxy error")
			panic(err.Error())
		}
	}()

	// serve grpc tracing using default mux
	go func() {
		trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
			// RemoteAddr is commonly in the form "IP" or "IP:port".
			// If it is in the form "IP:port", split off the port.
			host, _, err := net.SplitHostPort(req.RemoteAddr)
			if err != nil {
				host = req.RemoteAddr
			}
			switch host {
			case "localhost", "127.0.0.1", "::1", "172.18.0.1":
				return true, true
			default:
				return false, false
			}
		}
		err := http.ListenAndServe(":8181", nil)
		if err != nil {
			log.Println("grpc tracing server error")
			panic(err.Error())
		}
	}()

	// finally server the grpc
	err = grpcServer.Serve(lis) // this call is permanently blocking unless an error occurs -- so only error handling code should follow
	if err != nil {
		log.Println("grpc server net.Listen failed")
		panic(err.Error())
	}

}
