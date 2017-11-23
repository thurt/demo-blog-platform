package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/thurt/demo-blog-platform/cms/mysqlprovider"
	pb "github.com/thurt/demo-blog-platform/cms/proto"
	"google.golang.org/grpc"
)

const (
	PORT    = 10000
	DBHost  = "db"
	DBPort  = ":3306"
	DBUser  = "root"
	DBDbase = "cms"
)

var DBPass string

func main() {
	// connect to db
	DBPass = os.Getenv("DB_PASS")
	dbConn := fmt.Sprintf("%s:%s@tcp(%s)/%s", DBUser, DBPass, DBHost, DBDbase)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("Couldn't connect with mysql connection string")
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Println("Couldn't ping database server")
		panic(err.Error())
	}
	log.Println("Connected to db server:" + dbConn)

	// setup grpc server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", PORT))
	if err != nil {
		log.Println("failed to listen")
		panic(err.Error())
	}
	opts := []grpc.ServerOption{grpc.ConnectionTimeout(5 * time.Second)}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterCmsServer(grpcServer, mysqlprovider.New(db))
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

	// finally server the grpc
	err = grpcServer.Serve(lis) // this call is permanently blocking unless an error occurs -- so only error handling code should follow
	if err != nil {
		log.Println("grpc server net.Listen failed")
		panic(err.Error())
	}

}
