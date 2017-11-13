package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"golang.org/x/net/context"

	empty "github.com/golang/protobuf/ptypes/empty"
	pb "github.com/thurt/demo-blog-platform/api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"

	_ "github.com/go-sql-driver/mysql"
)

const (
	PORT    = 10000
	DBHost  = "db"
	DBPort  = ":3306"
	DBUser  = "root"
	DBDbase = "cms"
)

var database *sql.DB
var DBPass string

type cmsServer struct{}

func newServer() *cmsServer {
	s := new(cmsServer)
	return s
}

func (s *cmsServer) GetPage(ctx context.Context, pid *pb.PageId) (*pb.Page, error) {
	return &pb.Page{}, nil
}

func (s *cmsServer) CreatePage(ctx context.Context, cpr *pb.CreatePageRequest) (*pb.PageId, error) {
	// TODO: validate inputs
	// TODO: create a scheme to create a page_guid from the page_title (currently using hardCodedValue)
	hardCodedValue := "hard-coded"

	timeNow := time.Now()
	_, err := database.Exec("INSERT INTO pages SET page_guid=?, page_title=?, page_content=?, page_created=?, page_edited=?", hardCodedValue, cpr.GetTitle(), cpr.GetContent(), timeNow, timeNow)

	// TODO: return proper errors depending on the results of previous code (ie. sql row already exists, invalid inputs)
	if err != nil {
		return nil, grpc.Errorf(codes.Unknown, "Ouch!")
	}

	return &pb.PageId{hardCodedValue}, nil
}

func (s *cmsServer) DeletePage(ctx context.Context, pid *pb.PageId) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (s *cmsServer) GetPageComments(ctx context.Context, pid *pb.PageId) (*pb.Comments, error) {
	return &pb.Comments{}, nil
}

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
	database = db
	log.Println("Connected to db server:" + dbConn)

	// setup grpc server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", PORT))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterCmsServer(grpcServer, newServer())
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Println("Failed to start gRPC server")
		panic(err.Error())
	}
	log.Printf("Started gRPC server (port %d)", PORT)

	log.Println("Staring up proxy")
	ProxyStartup()
}
