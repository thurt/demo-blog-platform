package mysql_provider

import (
	"database/sql"
	"log"

	"golang.org/x/net/context"

	"github.com/VividCortex/mysqlerr"
	mysql "github.com/go-sql-driver/mysql"
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/thurt/demo-blog-platform/cms/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type provider struct{}

var database *sql.DB

func NewProvider(db *sql.DB) *provider {
	database = db
	s := new(provider)
	return s
}

func sqlErrorToGrpcError(err error) error {
	var e error

	if deviceErr, ok := err.(*mysql.MySQLError); ok {
		// these are just a few mysql errors that i found while trying to insert bad values
		// i'm not familiar with mysql errors but i suspect there will be more errors that come up commonly

		switch deviceErr.Number {
		case mysqlerr.ER_DUP_ENTRY: // tried to insert primary key value that already exists
			fallthrough
		case mysqlerr.ER_TRUNCATED_WRONG_VALUE_FOR_FIELD: // tried to insert/update a value with an incorrect type
			fallthrough
		case mysqlerr.ER_DATA_TOO_LONG: // tried to insert string that is too long
			fallthrough
		case mysqlerr.ER_NO_DEFAULT_FOR_FIELD: // tried to insert row without passing a required value
			fallthrough
		case mysqlerr.ER_ROW_IS_REFERENCED_2: // tried to update/delete a row key that is still referenced as a foreign key in another tabele
			fallthrough
		case mysqlerr.ER_NO_REFERENCED_ROW_2: // tried to supply a foreign key value that is not found in parent table
			e = grpc.Errorf(codes.InvalidArgument, err.Error())
		case mysqlerr.ER_PARSE_ERROR: // tried to execute a sql statement that has syntax error(s)
			e = grpc.Errorf(codes.Internal, err.Error())
		default: // unknown
			e = grpc.Errorf(codes.Unknown, err.Error())
		}
	} else if err == sql.ErrNoRows { // this error is specific only to QueryRow invocations
		e = grpc.Errorf(codes.NotFound, err.Error())
	} else {
		e = grpc.Errorf(codes.Unknown, err.Error())
	}

	return e
}

func (p *provider) GetPost(ctx context.Context, r *pb.PostRequest) (*pb.Post, error) {
	po := &pb.Post{}

	err := database.QueryRow("SELECT id, title, content, created, last_edited, published FROM posts WHERE id=?", r.GetId()).Scan(&po.Id, &po.Title, &po.Content, &po.Created, &po.LastEdited, &po.Published)

	if err != nil {
		log.Println(err)
		return nil, sqlErrorToGrpcError(err)
	}

	return po, nil
}

func (p *provider) CreatePost(ctx context.Context, r *pb.CreatePostRequest) (*pb.PostRequest, error) {
	// TODO: create a scheme to create an id from the title (currently using hardCodedValue)
	hardCodedValue := "hard-coded"

	_, err := database.Exec("INSERT INTO posts SET id=?, title=?, content=?", hardCodedValue, r.GetTitle(), r.GetContent())

	// TODO: return proper errors depending on the results of previous code (ie. sql row already exists, invalid inputs)
	if err != nil {
		log.Println(err)
		return nil, sqlErrorToGrpcError(err)
	}

	return &pb.PostRequest{hardCodedValue}, nil
}

func (p *provider) DeletePost(ctx context.Context, r *pb.PostRequest) (*empty.Empty, error) {
	// TODO: validate inputs
	res, err := database.Exec("DELETE FROM posts WHERE id=?", r.GetId())

	if err != nil {
		log.Println(err)
		return nil, sqlErrorToGrpcError(err)
	}

	ra, err := res.RowsAffected()
	// error only occurs if the sql implementation does not include RowsAffected() capability
	// if no rows were affected then it is also an error
	if err != nil || ra == 0 {
		log.Println(err)
		return nil, sqlErrorToGrpcError(err)
	}

	return &empty.Empty{}, nil
}

func (s *provider) GetPostComments(r *pb.PostRequest, stream pb.Cms_GetPostCommentsServer) error {
	cs, err := database.Query("SELECT id, content, created, last_edited, user_id, post_id FROM comments WHERE post_id=?", r.GetId())

	if err != nil {
		log.Println(err)
		return sqlErrorToGrpcError(err)
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

func (p *provider) CreateComment(ctx context.Context, r *pb.CreateCommentRequest) (*pb.CommentRequest, error) {
	res, err := database.Exec("INSERT INTO comments SET content=?, user_id=?, post_id=?", r.GetContent(), r.GetUserId(), r.GetPostId())

	if err != nil {
		log.Println(err)
		return nil, sqlErrorToGrpcError(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
		return nil, grpc.Errorf(codes.Unknown, "Ouch!")
	}

	return &pb.CommentRequest{uint32(id)}, nil
}

func (p *provider) CreateUser(ctx context.Context, r *pb.CreateUserRequest) (*pb.UserRequest, error) {
	_, err := database.Exec("INSERT INTO users SET id=?, email=?, password=?", r.GetId(), r.GetEmail(), r.GetPassword())

	if err != nil {
		log.Println(err)
		return nil, sqlErrorToGrpcError(err)
	}

	return &pb.UserRequest{r.GetId()}, nil
}

func (p *provider) DeleteComment(ctx context.Context, r *pb.CommentRequest) (*empty.Empty, error) {
	res, err := database.Exec("DELETE FROM comments WHERE id=?", r.GetId())

	if err != nil {
		log.Println(err)
		return nil, sqlErrorToGrpcError(err)
	}

	ra, err := res.RowsAffected()
	// error only occurs if the sql implementation does not include RowsAffected() capability
	// if no rows were affected then it is also an error
	if err != nil || ra == 0 {
		log.Println(err)
		return nil, sqlErrorToGrpcError(err)
	}

	return &empty.Empty{}, nil
}

func (p *provider) DeleteUser(ctx context.Context, r *pb.UserRequest) (*empty.Empty, error) {
	res, err := database.Exec("DELETE FROM users WHERE id=?", r.GetId())

	if err != nil {
		log.Println(err)
		return nil, sqlErrorToGrpcError(err)
	}

	ra, err := res.RowsAffected()
	// error only occurs if the sql implementation does not include RowsAffected() capability
	// if no rows were affected then it is also an error
	if err != nil || ra == 0 {
		log.Println(err)
		return nil, sqlErrorToGrpcError(err)
	}

	return &empty.Empty{}, nil
}

func (p *provider) GetComment(ctx context.Context, r *pb.CommentRequest) (*pb.Comment, error) {
	c := &pb.Comment{}

	err := database.QueryRow("SELECT id, content, created, last_edited, user_id, post_id FROM comments WHERE id=?", r.GetId()).Scan(&c.Id, &c.Content, &c.Created, &c.LastEdited, &c.UserId, &c.PostId)

	if err != nil {
		log.Println(err)
		return nil, sqlErrorToGrpcError(err)
	}

	return c, nil
}

func (p *provider) GetComments(_ *empty.Empty, stream pb.Cms_GetCommentsServer) error {
	cs, err := database.Query("SELECT id, content, created, last_edited, user_id, post_id FROM comments")

	if err != nil {
		log.Println(err)
		return sqlErrorToGrpcError(err)
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

func (p *provider) GetPosts(_ *empty.Empty, stream pb.Cms_GetPostsServer) error {
	ps, err := database.Query("SELECT id, title, content, created, last_edited FROM posts")

	if err != nil {
		log.Println(err)
		return sqlErrorToGrpcError(err)
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

func (p *provider) GetUser(ctx context.Context, r *pb.UserRequest) (*pb.User, error) {
	u := &pb.User{}

	err := database.QueryRow("SELECT id, email, created, last_active FROM users WHERE id=?", r.GetId()).Scan(&u.Id, &u.Email, &u.Created, &u.LastActive)

	if err != nil {
		log.Println(err)
		return nil, sqlErrorToGrpcError(err)
	}

	return u, nil
}

func (p *provider) GetUserComments(r *pb.UserRequest, stream pb.Cms_GetUserCommentsServer) error {
	cs, err := database.Query("SELECT id, content, created, last_edited, user_id, post_id FROM comments WHERE user_id=?", r.GetId())

	if err != nil {
		log.Println(err)
		return sqlErrorToGrpcError(err)
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

func (p *provider) PublishPost(ctx context.Context, r *pb.PostRequest) (*empty.Empty, error) {
	_, err := database.Exec("UPDATE posts SET published=TRUE WHERE id=?", r.GetId())

	if err != nil {
		log.Println(err)
		return nil, sqlErrorToGrpcError(err)
	}

	return &empty.Empty{}, nil
}

func (p *provider) UnPublishPost(ctx context.Context, r *pb.PostRequest) (*empty.Empty, error) {
	_, err := database.Exec("UPDATE posts SET published=FALSE WHERE id=?", r.GetId())

	if err != nil {
		log.Println(err)
		return nil, sqlErrorToGrpcError(err)
	}

	return &empty.Empty{}, nil
}

func (p *provider) UpdateComment(ctx context.Context, r *pb.UpdateCommentRequest) (*empty.Empty, error) {
	_, err := database.Exec("UPDATE comments SET content=? WHERE id=?", r.GetContent(), r.GetId())

	if err != nil {
		log.Println(err)
		return nil, sqlErrorToGrpcError(err)
	}

	return &empty.Empty{}, nil
}

func (p *provider) UpdatePost(ctx context.Context, r *pb.UpdatePostRequest) (*empty.Empty, error) {
	_, err := database.Exec("UPDATE posts SET title=?, content=? WHERE id=?", r.GetTitle(), r.GetContent(), r.GetId())

	if err != nil {
		log.Println(err)
		return nil, sqlErrorToGrpcError(err)
	}

	return &empty.Empty{}, nil
}
