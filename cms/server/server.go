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
	pb "github.com/thurt/demo-blog-platform/cms/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

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
	thisPage := &pb.Page{}
	// TODO: validate inputs

	err := database.QueryRow("SELECT page_title, page_content, page_date, page_guid FROM pages WHERE page_guid=?", pid.GetGuid()).Scan(&thisPage.Title, &thisPage.Content, &thisPage.Created, &thisPage.Guid)

	// TODO: return proper errors depending on results of previous code
	if err != nil {
		log.Println(err)
		return nil, grpc.Errorf(codes.Unknown, "Ouch!")
	}

	return thisPage, nil
}

func (s *cmsServer) CreatePage(ctx context.Context, cpr *pb.CreatePageRequest) (*pb.PageId, error) {
	// TODO: validate inputs
	// TODO: create a scheme to create a page_guid from the page_title (currently using hardCodedValue)
	hardCodedValue := "hard-coded"

	_, err := database.Exec("INSERT INTO pages SET page_guid=?, page_title=?, page_content=?", hardCodedValue, cpr.GetTitle(), cpr.GetContent())

	// TODO: return proper errors depending on the results of previous code (ie. sql row already exists, invalid inputs)
	if err != nil {
		log.Println(err)
		return nil, grpc.Errorf(codes.Unknown, "Ouch!")
	}

	return &pb.PageId{hardCodedValue}, nil
}

func (s *cmsServer) DeletePage(ctx context.Context, pid *pb.PageId) (*empty.Empty, error) {
	// TODO: validate inputs
	_, err := database.Exec("DELETE FROM pages WHERE page_guid=?", pid.GetGuid())

	// TODO: return proper errors depending on results of previous code
	if err != nil {
		log.Println(err)
		return nil, grpc.Errorf(codes.Unknown, "Ouch!")
	}
	return &empty.Empty{}, nil
}

func (s *cmsServer) GetPageComments(ctx context.Context, pid *pb.PageId) (*pb.Comments, error) {
	// TODO: validate inputs

	comments, err := database.Query("SELECT comment_name,comment_text,comment_date FROM comments WHERE page_guid=?", pid.GetGuid())

	// TODO: return proper errors depending on results of previous code

	if err != nil {
		log.Println(err)
		return nil, grpc.Errorf(codes.Unknown, "Ouch!")
	}
	defer comments.Close()

	var Comments []*pb.Comment

	for comments.Next() {
		thisComment := &pb.Comment{}
		comments.Scan(&thisComment.Author, &thisComment.Content, &thisComment.Created)
		Comments = append(Comments, thisComment)
	}

	if err = comments.Err(); err != nil {
		log.Println(err)
		return nil, grpc.Errorf(codes.Unknown, "Ouch!")
	}

	return &pb.Comments{Comments}, nil
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
		log.Println("failed to listen")
		panic(err.Error())
	}
	opts := []grpc.ServerOption{grpc.ConnectionTimeout(5 * time.Second)}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterCmsServer(grpcServer, newServer())
	log.Printf("Started grpc server on port %d", PORT)

	go func() {
		log.Println("Staring up rest-proxy")
		err = ProxyServe()
		if err != nil {
			log.Println("proxy error")
			panic(err.Error())
		}
	}()

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Println("grpc server net.Listen failed")
		panic(err.Error())
	}
}
