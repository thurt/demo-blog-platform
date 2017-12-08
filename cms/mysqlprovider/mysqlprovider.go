package mysqlprovider

import (
	"database/sql"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/thurt/demo-blog-platform/cms/domain"
	helper "github.com/thurt/demo-blog-platform/cms/mysqlprovider_helper"
	pb "github.com/thurt/demo-blog-platform/cms/proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	defaultRole uint32 = 2 // User
)

type provider struct {
	db *sql.DB
	q  sqlQueryI
}

type sqlQuery struct{}
type sqlQueryI interface {
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

func New(db *sql.DB) domain.Provider {
	s := &provider{db, NewSqlQuery()}
	return s
}

func NewSqlQuery() sqlQueryI {
	return &sqlQuery{}
}

func (q *sqlQuery) GetPost() string {
	return "SELECT id, title, content, created, last_edited, published FROM posts WHERE id=?"
}
func (p *provider) GetPost(ctx context.Context, r *pb.PostRequest) (*pb.Post, error) {
	po := &pb.Post{}

	err := p.db.QueryRow(p.q.GetPost(), r.GetId()).Scan(&po.Id, &po.Title, &po.Content, &po.Created, &po.LastEdited, &po.Published)

	if err != nil {
		log.Println(err)
		return nil, helper.SqlErrorToGrpcError(err)
	}

	return po, nil
}

func (q *sqlQuery) CreatePost() string {
	return "INSERT INTO posts SET slug=?, title=?, content=?"
}
func (p *provider) CreatePost(ctx context.Context, r *pb.CreatePostRequest) (*pb.PostRequest, error) {

	rs, err := p.db.Exec(p.q.CreatePost(), r.GetSlug(), r.GetTitle(), r.GetContent())

	if err != nil {
		log.Println(err)
		return nil, helper.SqlErrorToGrpcError(err)
	}

	id, err := rs.LastInsertId()
	if err != nil {
		log.Println(err)
		return nil, helper.SqlErrorToGrpcError(err)
	}

	return &pb.PostRequest{uint32(id)}, nil
}

func (q *sqlQuery) DeletePost() string {
	return "DELETE FROM posts WHERE id=?"
}
func (p *provider) DeletePost(ctx context.Context, r *pb.PostRequest) (*empty.Empty, error) {
	// TODO: validate inputs
	res, err := p.db.Exec(p.q.DeletePost(), r.GetId())

	if err != nil {
		log.Println(err)
		return nil, helper.SqlErrorToGrpcError(err)
	}

	ra, err := res.RowsAffected()
	// error only occurs if the sql implementation does not include RowsAffected() capability
	// if no rows were affected then it is also an error
	if err != nil || ra == 0 {
		log.Println(err)
		return nil, helper.SqlErrorToGrpcError(err)
	}

	return &empty.Empty{}, nil
}

func (q *sqlQuery) GetPostComments() string {
	return "SELECT id, content, created, last_edited, user_id, post_id FROM comments WHERE post_id=?"
}
func (p *provider) GetPostComments(r *pb.PostRequest, stream pb.Cms_GetPostCommentsServer) error {
	cs, err := p.db.Query(p.q.GetPostComments(), r.GetId())

	if err != nil {
		log.Println(err)
		return helper.SqlErrorToGrpcError(err)
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
		return status.Error(codes.Unknown, "Ouch!")
	}

	return nil
}

func (q *sqlQuery) CreateComment() string {
	return "INSERT INTO comments SET content=?, user_id=?, post_id=?"
}
func (p *provider) CreateComment(ctx context.Context, r *pb.CreateCommentRequest) (*pb.CommentRequest, error) {
	res, err := p.db.Exec(p.q.CreateComment(), r.GetContent(), r.GetUserId(), r.GetPostId())

	if err != nil {
		log.Println(err)
		return nil, helper.SqlErrorToGrpcError(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Unknown, "Ouch!")
	}

	return &pb.CommentRequest{uint32(id)}, nil
}

func (q *sqlQuery) CreateUser() string {
	return "INSERT INTO users SET id=?, email=?, password=?, role=?"
}
func (p *provider) CreateUser(ctx context.Context, r *pb.CreateUserRequest) (*pb.UserRequest, error) {
	_, err := p.db.Exec(p.q.CreateUser(), r.GetId(), r.GetEmail(), r.GetPassword(), defaultRole)

	if err != nil {
		log.Println(err)
		return nil, helper.SqlErrorToGrpcError(err)
	}

	return &pb.UserRequest{r.GetId()}, nil
}

func (q *sqlQuery) DeleteComment() string {
	return "DELETE FROM comments WHERE id=?"
}
func (p *provider) DeleteComment(ctx context.Context, r *pb.CommentRequest) (*empty.Empty, error) {
	res, err := p.db.Exec(p.q.DeleteComment(), r.GetId())

	if err != nil {
		log.Println(err)
		return nil, helper.SqlErrorToGrpcError(err)
	}

	ra, err := res.RowsAffected()
	// error only occurs if the sql implementation does not include RowsAffected() capability
	// if no rows were affected then it is also an error
	if err != nil || ra == 0 {
		log.Println(err)
		return nil, helper.SqlErrorToGrpcError(err)
	}

	return &empty.Empty{}, nil
}

func (q *sqlQuery) DeleteUser() string {
	return "DELETE FROM users WHERE id=?"
}
func (p *provider) DeleteUser(ctx context.Context, r *pb.UserRequest) (*empty.Empty, error) {
	res, err := p.db.Exec(p.q.DeleteUser(), r.GetId())

	if err != nil {
		log.Println(err)
		return nil, helper.SqlErrorToGrpcError(err)
	}

	ra, err := res.RowsAffected()
	// error only occurs if the sql implementation does not include RowsAffected() capability
	// if no rows were affected then it is also an error
	if err != nil || ra == 0 {
		log.Println(err)
		return nil, helper.SqlErrorToGrpcError(err)
	}

	return &empty.Empty{}, nil
}

func (q *sqlQuery) GetComment() string {
	return "SELECT id, content, created, last_edited, user_id, post_id FROM comments WHERE id=?"
}
func (p *provider) GetComment(ctx context.Context, r *pb.CommentRequest) (*pb.Comment, error) {
	c := &pb.Comment{}

	err := p.db.QueryRow(p.q.GetComment(), r.GetId()).Scan(&c.Id, &c.Content, &c.Created, &c.LastEdited, &c.UserId, &c.PostId)

	if err != nil {
		log.Println(err)
		return nil, helper.SqlErrorToGrpcError(err)
	}

	return c, nil
}

func (q *sqlQuery) GetComments() string {
	return "SELECT id, content, created, last_edited, user_id, post_id FROM comments"
}
func (p *provider) GetComments(_ *empty.Empty, stream pb.Cms_GetCommentsServer) error {
	cs, err := p.db.Query(p.q.GetComments())

	if err != nil {
		log.Println(err)
		return helper.SqlErrorToGrpcError(err)
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
		return status.Error(codes.Unknown, "Ouch!")
	}

	return nil
}

func (q *sqlQuery) GetPosts() string {
	return "SELECT id, title, content, created, last_edited FROM posts"
}
func (p *provider) GetPosts(_ *empty.Empty, stream pb.Cms_GetPostsServer) error {

	ps, err := p.db.Query(p.q.GetPosts())

	if err != nil {
		log.Println(err)
		return helper.SqlErrorToGrpcError(err)
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
		return status.Error(codes.Unknown, "Ouch!")
	}

	return nil
}

func (q *sqlQuery) GetUser() string {
	return "SELECT id, email, created, last_active, role FROM users WHERE id=?"
}
func (p *provider) GetUser(ctx context.Context, r *pb.UserRequest) (*pb.User, error) {
	u := &pb.User{}

	err := p.db.QueryRow(p.q.GetUser(), r.GetId()).Scan(&u.Id, &u.Email, &u.Created, &u.LastActive, &u.Role)

	if err != nil {
		log.Println(err)
		return nil, helper.SqlErrorToGrpcError(err)
	}

	return u, nil
}

func (q *sqlQuery) GetUserComments() string {
	return "SELECT id, content, created, last_edited, user_id, post_id FROM comments WHERE user_id=?"
}
func (p *provider) GetUserComments(r *pb.UserRequest, stream pb.Cms_GetUserCommentsServer) error {
	cs, err := p.db.Query(p.q.GetUserComments(), r.GetId())

	if err != nil {
		log.Println(err)
		return helper.SqlErrorToGrpcError(err)
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
		return status.Error(codes.Unknown, "Ouch!")
	}

	return nil
}

func (q *sqlQuery) PublishPost() string {
	return "UPDATE posts SET published=TRUE WHERE id=?"
}
func (p *provider) PublishPost(ctx context.Context, r *pb.PostRequest) (*empty.Empty, error) {
	_, err := p.db.Exec(p.q.PublishPost(), r.GetId())

	if err != nil {
		log.Println(err)
		return nil, helper.SqlErrorToGrpcError(err)
	}

	return &empty.Empty{}, nil
}

func (q *sqlQuery) UnPublishPost() string {
	return "UPDATE posts SET published=FALSE WHERE id=?"
}
func (p *provider) UnPublishPost(ctx context.Context, r *pb.PostRequest) (*empty.Empty, error) {
	_, err := p.db.Exec(p.q.UnPublishPost(), r.GetId())

	if err != nil {
		log.Println(err)
		return nil, helper.SqlErrorToGrpcError(err)
	}

	return &empty.Empty{}, nil
}

func (q *sqlQuery) UpdateComment() string {
	return "UPDATE comments SET content=? WHERE id=?"
}
func (p *provider) UpdateComment(ctx context.Context, r *pb.UpdateCommentRequest) (*empty.Empty, error) {
	_, err := p.db.Exec(p.q.UpdateComment(), r.GetContent(), r.GetId())

	if err != nil {
		log.Println(err)
		return nil, helper.SqlErrorToGrpcError(err)
	}

	return &empty.Empty{}, nil
}

func (q *sqlQuery) UpdatePost() string {
	return "UPDATE posts SET slug=?, title=?, content=? WHERE id=?"
}
func (p *provider) UpdatePost(ctx context.Context, r *pb.UpdatePostRequest) (*empty.Empty, error) {
	_, err := p.db.Exec(p.q.UpdatePost(), r.GetSlug(), r.GetTitle(), r.GetContent(), r.GetId())

	if err != nil {
		log.Println(err)
		return nil, helper.SqlErrorToGrpcError(err)
	}

	return &empty.Empty{}, nil
}
