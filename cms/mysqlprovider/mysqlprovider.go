package mysqlprovider

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

type Provider struct {
	db *sql.DB
	q  SqlQuery
}

type sqlQuery struct{}
type SqlQuery interface {
	GetPost() string
	CreatePost() string
	DeletePost() string
	CreateComment() string
	CreateUser() string
	DeleteComment() string
	DeleteUser() string
	GetComment() string
	GetPostComments() string
	GetComments() string
	GetPosts() string
	GetUser() string
	GetUserComments() string
	PublishPost() string
	UnPublishPost() string
	UpdateComment() string
	UpdatePost() string
}

func New(db *sql.DB) pb.CmsServer {
	s := &Provider{db, &sqlQuery{}}
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

func (q *sqlQuery) GetPost() string {
	return "SELECT id, title, content, created, last_edited, published FROM posts WHERE id=?"
}
func (p *Provider) GetPost(ctx context.Context, r *pb.PostRequest) (*pb.Post, error) {
	po := &pb.Post{}

	err := p.db.QueryRow(p.q.GetPost(), r.GetId()).Scan(&po.Id, &po.Title, &po.Content, &po.Created, &po.LastEdited, &po.Published)

	if err != nil {
		log.Println(err)
		return nil, sqlErrorToGrpcError(err)
	}

	return po, nil
}

func (q *sqlQuery) CreatePost() string {
	return "INSERT INTO posts SET title=?, content=?"
}
func (p *Provider) CreatePost(ctx context.Context, r *pb.CreatePostRequest) (*pb.PostRequest, error) {
	// TODO: create a scheme to create an id from the title (currently using hardCodedValue)
	hardCodedValue := "hard-coded"

	_, err := p.db.Exec(p.q.CreatePost(), hardCodedValue, r.GetTitle(), r.GetContent())

	// TODO: return proper errors depending on the results of previous code (ie. sql row already exists, invalid inputs)
	if err != nil {
		log.Println(err)
		return nil, sqlErrorToGrpcError(err)
	}

	return &pb.PostRequest{hardCodedValue}, nil
}

func (q *sqlQuery) DeletePost() string {
	return "DELETE FROM posts WHERE id=?"
}
func (p *Provider) DeletePost(ctx context.Context, r *pb.PostRequest) (*empty.Empty, error) {
	// TODO: validate inputs
	res, err := p.db.Exec(p.q.DeletePost(), r.GetId())

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

func (q *sqlQuery) GetPostComments() string {
	return "SELECT id, content, created, last_edited, user_id, post_id FROM comments WHERE post_id=?"
}
func (p *Provider) GetPostComments(r *pb.PostRequest, stream pb.Cms_GetPostCommentsServer) error {
	cs, err := p.db.Query(p.q.GetPostComments(), r.GetId())

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

func (q *sqlQuery) CreateComment() string {
	return "INSERT INTO comments SET content=?, user_id=?, post_id=?"
}
func (p *Provider) CreateComment(ctx context.Context, r *pb.CreateCommentRequest) (*pb.CommentRequest, error) {
	res, err := p.db.Exec(p.q.CreateComment(), r.GetContent(), r.GetUserId(), r.GetPostId())

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

func (q *sqlQuery) CreateUser() string {
	return "INSERT INTO users SET id=?, email=?, password=?"
}
func (p *Provider) CreateUser(ctx context.Context, r *pb.CreateUserRequest) (*pb.UserRequest, error) {
	_, err := p.db.Exec(p.q.CreateUser(), r.GetId(), r.GetEmail(), r.GetPassword())

	if err != nil {
		log.Println(err)
		return nil, sqlErrorToGrpcError(err)
	}

	return &pb.UserRequest{r.GetId()}, nil
}

func (q *sqlQuery) DeleteComment() string {
	return "DELETE FROM comments WHERE id=?"
}
func (p *Provider) DeleteComment(ctx context.Context, r *pb.CommentRequest) (*empty.Empty, error) {
	res, err := p.db.Exec(p.q.DeleteComment(), r.GetId())

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

func (q *sqlQuery) DeleteUser() string {
	return "DELETE FROM users WHERE id=?"
}
func (p *Provider) DeleteUser(ctx context.Context, r *pb.UserRequest) (*empty.Empty, error) {
	res, err := p.db.Exec(p.q.DeleteUser(), r.GetId())

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

func (q *sqlQuery) GetComment() string {
	return "SELECT id, content, created, last_edited, user_id, post_id FROM comments WHERE id=?"
}
func (p *Provider) GetComment(ctx context.Context, r *pb.CommentRequest) (*pb.Comment, error) {
	c := &pb.Comment{}

	err := p.db.QueryRow(p.q.GetComment(), r.GetId()).Scan(&c.Id, &c.Content, &c.Created, &c.LastEdited, &c.UserId, &c.PostId)

	if err != nil {
		log.Println(err)
		return nil, sqlErrorToGrpcError(err)
	}

	return c, nil
}

func (q *sqlQuery) GetComments() string {
	return "SELECT id, content, created, last_edited, user_id, post_id FROM comments"
}
func (p *Provider) GetComments(_ *empty.Empty, stream pb.Cms_GetCommentsServer) error {
	cs, err := p.db.Query(p.q.GetComments())

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

func (q *sqlQuery) GetPosts() string {
	return "SELECT id, title, content, created, last_edited FROM posts"
}
func (p *Provider) GetPosts(_ *empty.Empty, stream pb.Cms_GetPostsServer) error {

	ps, err := p.db.Query(p.q.GetPosts())

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

func (q *sqlQuery) GetUser() string {
	return "SELECT id, email, created, last_active FROM users WHERE id=?"
}
func (p *Provider) GetUser(ctx context.Context, r *pb.UserRequest) (*pb.User, error) {
	u := &pb.User{}

	err := p.db.QueryRow(p.q.GetUser(), r.GetId()).Scan(&u.Id, &u.Email, &u.Created, &u.LastActive)

	if err != nil {
		log.Println(err)
		return nil, sqlErrorToGrpcError(err)
	}

	return u, nil
}

func (q *sqlQuery) GetUserComments() string {
	return "SELECT id, content, created, last_edited, user_id, post_id FROM comments WHERE user_id=?"
}
func (p *Provider) GetUserComments(r *pb.UserRequest, stream pb.Cms_GetUserCommentsServer) error {
	cs, err := p.db.Query(p.q.GetUserComments(), r.GetId())

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

func (q *sqlQuery) PublishPost() string {
	return "UPDATE posts SET published=TRUE WHERE id=?"
}
func (p *Provider) PublishPost(ctx context.Context, r *pb.PostRequest) (*empty.Empty, error) {
	_, err := p.db.Exec(p.q.PublishPost(), r.GetId())

	if err != nil {
		log.Println(err)
		return nil, sqlErrorToGrpcError(err)
	}

	return &empty.Empty{}, nil
}

func (q *sqlQuery) UnPublishPost() string {
	return "UPDATE posts SET published=FALSE WHERE id=?"
}
func (p *Provider) UnPublishPost(ctx context.Context, r *pb.PostRequest) (*empty.Empty, error) {
	_, err := p.db.Exec(p.q.UnPublishPost(), r.GetId())

	if err != nil {
		log.Println(err)
		return nil, sqlErrorToGrpcError(err)
	}

	return &empty.Empty{}, nil
}

func (q *sqlQuery) UpdateComment() string {
	return "UPDATE comments SET content=? WHERE id=?"
}
func (p *Provider) UpdateComment(ctx context.Context, r *pb.UpdateCommentRequest) (*empty.Empty, error) {
	_, err := p.db.Exec(p.q.UpdateComment(), r.GetContent(), r.GetId())

	if err != nil {
		log.Println(err)
		return nil, sqlErrorToGrpcError(err)
	}

	return &empty.Empty{}, nil
}

func (q *sqlQuery) UpdatePost() string {
	return "UPDATE posts SET title=?, content=? WHERE id=?"
}
func (p *Provider) UpdatePost(ctx context.Context, r *pb.UpdatePostRequest) (*empty.Empty, error) {
	_, err := p.db.Exec(p.q.UpdatePost(), r.GetTitle(), r.GetContent(), r.GetId())

	if err != nil {
		log.Println(err)
		return nil, sqlErrorToGrpcError(err)
	}

	return &empty.Empty{}, nil
}
