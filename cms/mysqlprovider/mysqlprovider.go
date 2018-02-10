package mysqlprovider

import (
	"database/sql"
	"errors"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/thurt/demo-blog-platform/cms/domain"
	"github.com/thurt/demo-blog-platform/cms/mysqlprovider/query"
	pb "github.com/thurt/demo-blog-platform/cms/proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type provider struct {
	db *sql.DB
	q  *query.Query
}

func New(db *sql.DB) domain.Provider {
	s := &provider{db, &query.Query{}}
	return s
}

func (p *provider) GetUserPassword(ctx context.Context, r *pb.UserRequest) (*pb.UserPassword, error) {
	u := &pb.UserPassword{}
	err := p.db.QueryRow(p.q.GetUserPassword(r)).Scan(&u.Password)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (p *provider) GetPost(ctx context.Context, r *pb.PostRequest) (*pb.Post, error) {
	po := &pb.Post{}
	err := p.db.QueryRow(p.q.GetPost(r)).Scan(&po.Id, &po.Title, &po.Content, &po.Created, &po.LastEdited, &po.Published, &po.Slug)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return po, nil
}

func (p *provider) GetPostBySlug(ctx context.Context, r *pb.PostBySlugRequest) (*pb.Post, error) {
	po := &pb.Post{}
	err := p.db.QueryRow(p.q.GetPostBySlug(r)).Scan(&po.Id, &po.Title, &po.Content, &po.Created, &po.LastEdited, &po.Published, &po.Slug)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return po, nil
}

func (p *provider) CreatePost(ctx context.Context, r *pb.CreatePostWithSlug) (*pb.PostRequest, error) {
	rs, err := p.db.Exec(p.q.CreatePost(r))
	if err != nil {
		return nil, err
	}

	id, _ := rs.LastInsertId()
	return &pb.PostRequest{uint32(id)}, nil
}

func (p *provider) DeletePost(ctx context.Context, r *pb.PostRequest) (*empty.Empty, error) {
	_, err := p.db.Exec(p.q.DeletePost(r))

	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
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

func (p *provider) CreateComment(ctx context.Context, r *pb.CreateCommentRequest) (*pb.CommentRequest, error) {
	rs, err := p.db.Exec(p.q.CreateComment(r))
	if err != nil {
		return nil, err
	}

	id, _ := rs.LastInsertId()
	return &pb.CommentRequest{uint32(id)}, nil
}

func (p *provider) CreateUser(ctx context.Context, r *pb.CreateUserWithRole) (*pb.UserRequest, error) {
	_, err := p.db.Exec(p.q.CreateUser(r))

	if err != nil {
		return nil, err
	}

	return &pb.UserRequest{r.User.GetId()}, nil
}

func (p *provider) DeleteComment(ctx context.Context, r *pb.CommentRequest) (*empty.Empty, error) {
	_, err := p.db.Exec(p.q.DeleteComment(r))

	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (p *provider) DeleteUser(ctx context.Context, r *pb.UserRequest) (*empty.Empty, error) {
	_, err := p.db.Exec(p.q.DeleteUser(r))

	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (p *provider) GetComment(ctx context.Context, r *pb.CommentRequest) (*pb.Comment, error) {
	c := &pb.Comment{}
	err := p.db.QueryRow(p.q.GetComment(r)).Scan(&c.Id, &c.Content, &c.Created, &c.LastEdited, &c.UserId, &c.PostId)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (p *provider) GetComments(r *empty.Empty, stream pb.Cms_GetCommentsServer) error {
	cs, err := p.db.Query(p.q.GetComments(r))

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

func (p *provider) GetPosts(r *pb.GetPostsOptions, stream pb.Cms_GetPostsServer) error {
	var sql string
	if r.GetIncludeUnPublished() == true {
		// (options) requires dispatching the correct sql request when IncludeUnPublished is true
		sql = p.q.GetPosts()
	} else {
		// (options) requires dispatching the correct sql request when IncludeUnPublished is false
		sql = p.q.GetPublishedPosts()
	}

	ps, err := p.db.Query(sql)
	if err != nil {
		return err
	}
	defer ps.Close()

	for ps.Next() {
		po := &pb.Post{}
		err = ps.Scan(&po.Id, &po.Title, &po.Content, &po.Created, &po.LastEdited, &po.Published, &po.Slug)
		if err != nil {
			return err
		}
		err := stream.Send(po)
		if err != nil {
			return err
		}
	}

	if err = ps.Err(); err != nil {
		return err
	}

	return nil
}

// GetUser gets the user from the db. It returns a zero-value struct if the user is not found.
func (p *provider) GetUser(ctx context.Context, r *pb.UserRequest) (*pb.User, error) {
	u := &pb.User{}
	err := p.db.QueryRow(p.q.GetUser(r)).Scan(&u.Id, &u.Email, &u.Created, &u.LastActive, &u.Role)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return u, nil
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

func (p *provider) UpdateComment(ctx context.Context, r *pb.UpdateCommentRequest) (*empty.Empty, error) {
	_, err := p.db.Exec(p.q.UpdateComment(r))

	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (p *provider) UpdatePost(ctx context.Context, r *pb.UpdatePostWithSlug) (*empty.Empty, error) {
	_, err := p.db.Exec(p.q.UpdatePost(r))

	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (p *provider) AdminExists(ctx context.Context, r *empty.Empty) (*wrappers.BoolValue, error) {
	var exists int
	adminExists := &wrappers.BoolValue{}

	err := p.db.QueryRow(p.q.AdminExists(r)).Scan(&exists)

	if err != nil {
		return nil, err
	}

	if exists == 0 {
		adminExists.Value = false
	} else {
		adminExists.Value = true
	}

	return adminExists, nil
}

func (p *provider) UpdateUserLastActive(ctx context.Context, r *pb.UserRequest) (*empty.Empty, error) {
	return nil, errors.New("")
}
