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

func (s *cmsServer) GetPost(ctx context.Context, r *pb.PostRequest) (*pb.Post, error) {
	p := &pb.Post{}

	err := database.QueryRow("SELECT page_title, page_content, page_date, page_guid FROM pages WHERE page_guid=?", r.GetId()).Scan(&p.Title, &p.Content, &p.Created, &p.Id)

	// TODO: return proper errors depending on results of previous code
	if err != nil {
		log.Println(err)
		return nil, grpc.Errorf(codes.Unknown, "Ouch!")
	}

	return p, nil
}

func (s *cmsServer) CreatePost(ctx context.Context, r *pb.CreatePostRequest) (*pb.PostRequest, error) {
	// TODO: create a scheme to create a page_guid from the page_title (currently using hardCodedValue)
	hardCodedValue := "hard-coded"

	_, err := database.Exec("INSERT INTO pages SET page_guid=?, page_title=?, page_content=?", hardCodedValue, r.GetTitle(), r.GetContent())

	// TODO: return proper errors depending on the results of previous code (ie. sql row already exists, invalid inputs)
	if err != nil {
		log.Println(err)
		return nil, grpc.Errorf(codes.Unknown, "Ouch!")
	}

	return &pb.PostRequest{hardCodedValue}, nil
}

func (s *cmsServer) DeletePost(ctx context.Context, pr *pb.PostRequest) (*empty.Empty, error) {
	// TODO: validate inputs
	_, err := database.Exec("DELETE FROM pages WHERE page_guid=?", pr.GetId())

	// TODO: return proper errors depending on results of previous code
	if err != nil {
		log.Println(err)
		return nil, grpc.Errorf(codes.Unknown, "Ouch!")
	}
	return &empty.Empty{}, nil
}

func (s *cmsServer) GetPostComments(r *pb.PostRequest, stream pb.Cms_GetPostCommentsServer) error {
	cs, err := database.Query("SELECT id, content, created, last_edited, user_id, post_id FROM comments WHERE post_id=?", r.GetId())

	// TODO: return proper errors depending on results of previous code

	if err != nil {
		log.Println(err)
		return grpc.Errorf(codes.Unknown, "Ouch!")
	}
	defer cs.Close()

	for cs.Next() {
		c := &pb.Comment{}
		cs.Scan(&c.Id, &c.Content, &c.Created, &c.LastEdited, &c.UserId, &c.PostId)
		err := stream.Send(c)
		if err != nil {
			return err
		}
	}

	if err = cs.Err(); err != nil {
		log.Println(err)
		return grpc.Errorf(codes.Unknown, "Ouch!")
	}

	return nil
}

func (s *cmsServer) CreateComment(ctx context.Context, r *pb.CreateCommentRequest) (*pb.CommentRequest, error) {
	res, err := database.Exec("INSERT INTO comments SET content=?, user_id=?, post_id=?", r.GetContent(), r.GetUserId(), r.GetPostId())

	if err != nil {
		log.Println(err)
		return nil, grpc.Errorf(codes.Unknown, "Ouch!")
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
		return nil, grpc.Errorf(codes.Unknown, "Ouch!")
	}

	return &pb.CommentRequest{uint32(id)}, nil
}

func (s *cmsServer) CreateUser(ctx context.Context, r *pb.CreateUserRequest) (*pb.UserRequest, error) {
	_, err := database.Exec("INSERT INTO users SET id=?, email=?, password=?", r.GetId(), r.GetEmail(), r.GetPassword())

	if err != nil {
		log.Println(err)
		return nil, grpc.Errorf(codes.Unknown, "Ouch!")
	}

	return &pb.UserRequest{r.GetId()}, nil
}

func (s *cmsServer) DeleteComment(ctx context.Context, r *pb.CommentRequest) (*empty.Empty, error) {
	_, err := database.Exec("DELETE FROM comments WHERE id=?", r.GetId())

	if err != nil {
		log.Println(err)
		return nil, grpc.Errorf(codes.Unknown, "Ouch!")
	}

	return &empty.Empty{}, nil
}

func (s *cmsServer) DeleteUser(ctx context.Context, r *pb.UserRequest) (*empty.Empty, error) {
	_, err := database.Exec("DELETE FROM users WHERE id=?", r.GetId())

	if err != nil {
		log.Println(err)
		return nil, grpc.Errorf(codes.Unknown, "Ouch!")
	}

	return &empty.Empty{}, nil
}

func (s *cmsServer) GetComment(ctx context.Context, r *pb.CommentRequest) (*pb.Comment, error) {
	c := &pb.Comment{}

	err := database.QueryRow("SELECT id, content, created, last_edited, user_id, post_id FROM comments WHERE id=?", r.GetId()).Scan(&c.Id, &c.Content, &c.Created, &c.LastEdited, &c.UserId, &c.PostId)

	if err != nil {
		log.Println(err)
		return nil, grpc.Errorf(codes.Unknown, "Ouch!")
	}

	return c, nil
}

func (s *cmsServer) GetComments(_ *empty.Empty, stream pb.Cms_GetCommentsServer) error {
	cs, err := database.Query("SELECT id, content, created, last_edited, user_id, post_id FROM comments")

	// TODO: return proper errors depending on results of previous code

	if err != nil {
		log.Println(err)
		return grpc.Errorf(codes.Unknown, "Ouch!")
	}
	defer cs.Close()

	for cs.Next() {
		c := &pb.Comment{}
		cs.Scan(&c.Id, &c.Content, &c.Created, &c.LastEdited, &c.UserId, &c.PostId)
		err := stream.Send(c)
		if err != nil {
			return err
		}
	}

	if err = cs.Err(); err != nil {
		log.Println(err)
		return grpc.Errorf(codes.Unknown, "Ouch!")
	}

	return nil
}

func (s *cmsServer) GetPosts(_ *empty.Empty, stream pb.Cms_GetPostsServer) error {
	ps, err := database.Query("SELECT id, title, content, created, last_edited FROM posts")

	// TODO: return proper errors depending on results of previous code

	if err != nil {
		log.Println(err)
		return grpc.Errorf(codes.Unknown, "Ouch!")
	}
	defer ps.Close()

	for ps.Next() {
		p := &pb.Post{}
		ps.Scan(&p.Id, &p.Title, &p.Content, &p.Created, &p.LastEdited)
		err := stream.Send(p)
		if err != nil {
			return err
		}
	}

	if err = ps.Err(); err != nil {
		log.Println(err)
		return grpc.Errorf(codes.Unknown, "Ouch!")
	}

	return nil
}

func (s *cmsServer) GetUser(ctx context.Context, r *pb.UserRequest) (*pb.User, error) {
	u := &pb.User{}

	err := database.QueryRow("SELECT id, email, created, last_active FROM users WHERE id=?", r.GetId()).Scan(&u.Id, &u.Email, &u.Created, &u.LastActive)

	if err != nil {
		log.Println(err)
		return nil, grpc.Errorf(codes.Unknown, "Ouch!")
	}

	return u, nil

}

func (s *cmsServer) GetUserComments(r *pb.UserRequest, stream pb.Cms_GetUserCommentsServer) error {
	cs, err := database.Query("SELECT id, content, created, last_edited, user_id, post_id FROM comments WHERE user_id=?", r.GetId())

	// TODO: return proper errors depending on results of previous code

	if err != nil {
		log.Println(err)
		return grpc.Errorf(codes.Unknown, "Ouch!")
	}
	defer cs.Close()

	for cs.Next() {
		c := &pb.Comment{}
		cs.Scan(&c.Id, &c.Content, &c.Created, &c.LastEdited, &c.UserId, &c.PostId)
		err := stream.Send(c)
		if err != nil {
			return err
		}
	}

	if err = cs.Err(); err != nil {
		log.Println(err)
		return grpc.Errorf(codes.Unknown, "Ouch!")
	}

	return nil
}

func (s *cmsServer) PublishPost(ctx context.Context, r *pb.PostRequest) (*empty.Empty, error) {
	_, err := database.Exec("UPDATE posts SET published=TRUE WHERE id=?", r.GetId())

	if err != nil {
		log.Println(err)
		return nil, grpc.Errorf(codes.Unknown, "Ouch!")
	}

	return &empty.Empty{}, nil
}

func (s *cmsServer) UnPublishPost(ctx context.Context, r *pb.PostRequest) (*empty.Empty, error) {
	_, err := database.Exec("UPDATE posts SET published=FALSE WHERE id=?", r.GetId())

	if err != nil {
		log.Println(err)
		return nil, grpc.Errorf(codes.Unknown, "Ouch!")
	}

	return &empty.Empty{}, nil
}

func (s *cmsServer) UpdateComment(ctx context.Context, r *pb.UpdateCommentRequest) (*empty.Empty, error) {
	_, err := database.Exec("UPDATE comments SET content=? WHERE id=?", r.GetContent(), r.GetId())

	if err != nil {
		log.Println(err)
		return nil, grpc.Errorf(codes.Unknown, "Ouch!")
	}

	return &empty.Empty{}, nil
}

func (s *cmsServer) UpdatePost(ctx context.Context, r *pb.UpdatePostRequest) (*empty.Empty, error) {
	_, err := database.Exec("UPDATE posts SET title=?, content=? WHERE id=?", r.GetTitle(), r.GetContent(), r.GetId())

	if err != nil {
		log.Println(err)
		return nil, grpc.Errorf(codes.Unknown, "Ouch!")
	}

	return &empty.Empty{}, nil
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
