package mysqlprovider

import (
	"database/sql"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/thurt/demo-blog-platform/cms/domain"
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
	GetPost(*pb.PostRequest) string
	CreatePost(*pb.CreatePostRequest) string
	DeletePost(*pb.PostRequest) string
	CreateComment(*pb.CreateCommentRequest) string
	CreateUser(*pb.CreateUserRequest) string
	DeleteComment(*pb.CommentRequest) string
	DeleteUser(*pb.UserRequest) string
	GetComment(*pb.CommentRequest) string
	GetPostComments(*pb.PostRequest) string
	GetComments() string
	GetPosts(*empty.Empty) string
	GetUser(*pb.UserRequest) string
	GetUserComments(*pb.UserRequest) string
	PublishPost(*pb.PostRequest) string
	UnPublishPost(*pb.PostRequest) string
	UpdateComment(*pb.UpdateCommentRequest) string
	UpdatePost(*pb.UpdatePostRequest) string
}

func New(db *sql.DB) domain.Provider {
	s := &provider{db, NewSqlQuery()}
	return s
}

func NewSqlQuery() sqlQueryI {
	return &sqlQuery{}
}

func (q *sqlQuery) GetPost(r *pb.PostRequest) string {
	return fmt.Sprintf("SELECT id, title, content, created, last_edited, published, slug FROM posts WHERE id=%d", r.GetId())
}
func (p *provider) GetPost(ctx context.Context, r *pb.PostRequest) (*pb.Post, error) {
	po := &pb.Post{}
	err := p.db.QueryRow(p.q.GetPost(r)).Scan(&po.Id, &po.Title, &po.Content, &po.Created, &po.LastEdited, &po.Published, &po.Slug)
	if err != nil {
		return nil, err
	}
	return po, nil
}

func (q *sqlQuery) CreatePost(r *pb.CreatePostRequest) string {
	return fmt.Sprintf("INSERT INTO posts SET slug=%q, title=%q, content=%q", r.GetSlug(), r.GetTitle(), r.GetContent())
}
func (p *provider) CreatePost(ctx context.Context, r *pb.CreatePostRequest) (*pb.PostRequest, error) {
	rs, err := p.db.Exec(p.q.CreatePost(r))
	if err != nil {
		return nil, err
	}

	id, _ := rs.LastInsertId()
	return &pb.PostRequest{uint32(id)}, nil
}

func (q *sqlQuery) DeletePost(r *pb.PostRequest) string {
	return fmt.Sprintf("DELETE FROM posts WHERE id=%d", r.GetId())
}
func (p *provider) DeletePost(ctx context.Context, r *pb.PostRequest) (*empty.Empty, error) {
	_, err := p.db.Exec(p.q.DeletePost(r))

	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (q *sqlQuery) GetPostComments(r *pb.PostRequest) string {
	return fmt.Sprintf("SELECT id, content, created, last_edited, user_id, post_id FROM comments WHERE post_id=%d", r.GetId())
}
func (p *provider) GetPostComments(r *pb.PostRequest, stream pb.Cms_GetPostCommentsServer) error {
	cs, err := p.db.Query(p.q.GetPostComments(r))

	if err != nil {
		return err
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
		return status.Error(codes.Unknown, "Ouch!")
	}

	return nil
}

func (q *sqlQuery) CreateComment(r *pb.CreateCommentRequest) string {
	return fmt.Sprintf("INSERT INTO comments SET content=%q, user_id=%q, post_id=%d", r.GetContent(), r.GetUserId(), r.GetPostId())
}
func (p *provider) CreateComment(ctx context.Context, r *pb.CreateCommentRequest) (*pb.CommentRequest, error) {
	rs, err := p.db.Exec(p.q.CreateComment(r))
	if err != nil {
		return nil, err
	}

	id, _ := rs.LastInsertId()
	return &pb.CommentRequest{uint32(id)}, nil
}

func (q *sqlQuery) CreateUser(r *pb.CreateUserRequest) string {
	return fmt.Sprintf("INSERT INTO users SET id=%q, email=%q, password=%q, role=%d", r.GetId(), r.GetEmail(), r.GetPassword(), defaultRole)
}
func (p *provider) CreateUser(ctx context.Context, r *pb.CreateUserRequest) (*pb.UserRequest, error) {
	_, err := p.db.Exec(p.q.CreateUser(r))

	if err != nil {
		return nil, err
	}

	return &pb.UserRequest{r.GetId()}, nil
}

func (q *sqlQuery) DeleteComment(r *pb.CommentRequest) string {
	return fmt.Sprintf("DELETE FROM comments WHERE id=%d", r.GetId())
}
func (p *provider) DeleteComment(ctx context.Context, r *pb.CommentRequest) (*empty.Empty, error) {
	_, err := p.db.Exec(p.q.DeleteComment(r))

	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (q *sqlQuery) DeleteUser(r *pb.UserRequest) string {
	return fmt.Sprintf("DELETE FROM users WHERE id=%q", r.GetId())
}
func (p *provider) DeleteUser(ctx context.Context, r *pb.UserRequest) (*empty.Empty, error) {
	_, err := p.db.Exec(p.q.DeleteUser(r))

	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (q *sqlQuery) GetComment(r *pb.CommentRequest) string {
	return fmt.Sprintf("SELECT id, content, created, last_edited, user_id, post_id FROM comments WHERE id=%d", r.GetId())
}
func (p *provider) GetComment(ctx context.Context, r *pb.CommentRequest) (*pb.Comment, error) {
	c := &pb.Comment{}
	err := p.db.QueryRow(p.q.GetComment(r)).Scan(&c.Id, &c.Content, &c.Created, &c.LastEdited, &c.UserId, &c.PostId)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (q *sqlQuery) GetComments() string {
	return "SELECT id, content, created, last_edited, user_id, post_id FROM comments"
}
func (p *provider) GetComments(_ *empty.Empty, stream pb.Cms_GetCommentsServer) error {
	cs, err := p.db.Query(p.q.GetComments())

	if err != nil {
		return err
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
		return status.Error(codes.Unknown, "Ouch!")
	}

	return nil
}

func (q *sqlQuery) GetPosts(_ *empty.Empty) string {
	return "SELECT id, title, content, created, last_edited, published, slug FROM posts"
}
func (p *provider) GetPosts(r *empty.Empty, stream pb.Cms_GetPostsServer) error {
	ps, err := p.db.Query(p.q.GetPosts(r))
	if err != nil {
		return err
	}
	defer ps.Close()

	for ps.Next() {
		p := &pb.Post{}
		err = ps.Scan(&p.Id, &p.Title, &p.Content, &p.Created, &p.LastEdited, &p.Published, &p.Slug)
		if err != nil {
			return err
		}
		err := stream.Send(p)
		if err != nil {
			return err
		}
	}

	if err = ps.Err(); err != nil {
		return err
	}

	return nil
}

func (q *sqlQuery) GetUser(r *pb.UserRequest) string {
	return fmt.Sprintf("SELECT id, email, created, last_active, role FROM users WHERE id=%q", r.GetId())
}
func (p *provider) GetUser(ctx context.Context, r *pb.UserRequest) (*pb.User, error) {
	u := &pb.User{}
	err := p.db.QueryRow(p.q.GetUser(r)).Scan(&u, &u.Id, &u.Email, &u.Created, &u.LastActive, &u.Role)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (q *sqlQuery) GetUserComments(r *pb.UserRequest) string {
	return fmt.Sprintf("SELECT id, content, created, last_edited, user_id, post_id FROM comments WHERE user_id=%q", r.GetId())
}
func (p *provider) GetUserComments(r *pb.UserRequest, stream pb.Cms_GetUserCommentsServer) error {
	cs, err := p.db.Query(p.q.GetUserComments(r))

	if err != nil {
		return err
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
		return status.Error(codes.Unknown, "Ouch!")
	}

	return nil
}

func (q *sqlQuery) PublishPost(r *pb.PostRequest) string {
	return fmt.Sprintf("UPDATE posts SET published=TRUE WHERE id=%d", r.GetId())
}
func (p *provider) PublishPost(ctx context.Context, r *pb.PostRequest) (*empty.Empty, error) {
	_, err := p.db.Exec(p.q.PublishPost(r))

	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (q *sqlQuery) UnPublishPost(r *pb.PostRequest) string {
	return fmt.Sprintf("UPDATE posts SET published=FALSE WHERE id=%d", r.GetId())
}
func (p *provider) UnPublishPost(ctx context.Context, r *pb.PostRequest) (*empty.Empty, error) {
	_, err := p.db.Exec(p.q.UnPublishPost(r))

	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (q *sqlQuery) UpdateComment(r *pb.UpdateCommentRequest) string {
	return fmt.Sprintf("UPDATE comments SET content=%q WHERE id=%d", r.GetContent(), r.GetId())
}
func (p *provider) UpdateComment(ctx context.Context, r *pb.UpdateCommentRequest) (*empty.Empty, error) {
	_, err := p.db.Exec(p.q.UpdateComment(r))

	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (q *sqlQuery) UpdatePost(r *pb.UpdatePostRequest) string {
	return fmt.Sprintf("UPDATE posts SET slug=%q, title=%q, content=%q WHERE id=%d", r.GetSlug(), r.GetTitle(), r.GetContent(), r.GetId())
}
func (p *provider) UpdatePost(ctx context.Context, r *pb.UpdatePostRequest) (*empty.Empty, error) {
	_, err := p.db.Exec(p.q.UpdatePost(r))

	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
