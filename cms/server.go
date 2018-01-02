package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/memcachier/mc"
	"github.com/thurt/demo-blog-platform/cms/authentication"
	"github.com/thurt/demo-blog-platform/cms/authorization"
	"github.com/thurt/demo-blog-platform/cms/mysqlprovider"
	"github.com/thurt/demo-blog-platform/cms/mysqlprovider_internal"
	pb "github.com/thurt/demo-blog-platform/cms/proto"
	"github.com/thurt/demo-blog-platform/cms/usecases"
	trace "golang.org/x/net/trace"
	"google.golang.org/grpc"
)

const (
	PORT = 10000
)

var MYSQL_CONNECTION string
var (
	MEMCACHE_HOST     string
	MEMCACHE_USER     string
	MEMCACHE_PASSWORD string
)

func main() {
	//connect to memcache
	MEMCACHE_HOST = os.Getenv("MEMCACHE_HOST")
	cn, err := mc.Dial("tcp", MEMCACHE_HOST)
	if err != nil {
		panic(err)
	}
	defer cn.Close()

	MEMCACHE_USER = os.Getenv("MEMCACHE_USER")
	MEMCACHE_PASSWORD = os.Getenv("MEMCACHE_PASSWORD")
	err = cn.Auth(MEMCACHE_USER, MEMCACHE_PASSWORD)
	if err != nil {
		panic(err)
	}
	log.Println("Connected to memcache host")

	// connect to db
	MYSQL_CONNECTION = os.Getenv("MYSQL_CONNECTION")
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

	// setup grpc server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", PORT))
	if err != nil {
		log.Println("failed to listen")
		panic(err.Error())
	}

	authProvider, authFunc := authentication.New(cn, 8*time.Hour)

	grpc.EnableTracing = true
	opts := []grpc.ServerOption{
		grpc.ConnectionTimeout(5 * time.Second),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_validator.UnaryServerInterceptor(),
			grpc_auth.UnaryServerInterceptor(authFunc),
		)),
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterCmsServer(grpcServer, authorization.New(usecases.New(mysqlprovider.New(db), mysqlprovider_internal.New(db), authProvider)))
	log.Printf("Started grpc server on port %d", PORT)

	// setup rest proxy server
	go func() {
		log.Println("Staring up rest-proxy")
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
