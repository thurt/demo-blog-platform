package query

import (
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/thurt/demo-blog-platform/cms/proto"
)

type Query struct{}

func (q *Query) GetUser(r *pb.UserRequest) string {
	return fmt.Sprintf("SELECT id, email, created, last_active, role FROM users WHERE id=%q", r.GetId())
}

func (q *Query) AdminExists(_ *empty.Empty) string {
	return fmt.Sprintf("SELECT EXISTS(SELECT id FROM users WHERE role=%d)", pb.UserRole_ADMIN)
}

func (q *Query) GetUserPassword(r *pb.UserRequest) string {
	return fmt.Sprintf("SELECT password FROM users WHERE id=%q", r.GetId())
}

func (q *Query) UpdatePost(r *pb.UpdatePostWithSlug) string {
	return fmt.Sprintf("UPDATE posts SET slug=%q, title=%q, content=%q WHERE id=%d", r.GetSlug(), r.Post.GetTitle(), r.Post.GetContent(), r.Post.GetId())
}

func (q *Query) UpdateComment(r *pb.UpdateCommentRequest) string {
	return fmt.Sprintf("UPDATE comments SET content=%q WHERE id=%d", r.GetContent(), r.GetId())
}

func (q *Query) UnPublishPost(r *pb.PostRequest) string {
	return fmt.Sprintf("UPDATE posts SET published=FALSE WHERE id=%d", r.GetId())
}

func (q *Query) PublishPost(r *pb.PostRequest) string {
	return fmt.Sprintf("UPDATE posts SET published=TRUE WHERE id=%d", r.GetId())
}

func (q *Query) GetUserComments(r *pb.UserRequest) string {
	return fmt.Sprintf("SELECT id, content, created, last_edited, user_id, post_id FROM comments WHERE user_id=%q", r.GetId())
}

func (q *Query) GetPosts() string {
	return "SELECT id, title, content, created, last_edited, published, slug FROM posts"
}

func (q *Query) GetPublishedPosts() string {
	return "SELECT id, title, content, created, last_edited, published, slug FROM posts WHERE published=1"
}

func (q *Query) GetComments(_ *empty.Empty) string {
	return "SELECT id, content, created, last_edited, user_id, post_id FROM comments"
}

func (q *Query) GetComment(r *pb.CommentRequest) string {
	return fmt.Sprintf("SELECT id, content, created, last_edited, user_id, post_id FROM comments WHERE id=%d", r.GetId())
}

func (q *Query) DeleteUser(r *pb.UserRequest) string {
	return fmt.Sprintf("DELETE FROM users WHERE id=%q", r.GetId())
}

func (q *Query) DeleteComment(r *pb.CommentRequest) string {
	return fmt.Sprintf("DELETE FROM comments WHERE id=%d", r.GetId())
}

func (q *Query) CreateUser(r *pb.CreateUserWithRole) string {
	return fmt.Sprintf("INSERT INTO users SET id=%q, email=%q, password=%q, role=%d", r.User.GetId(), r.User.GetEmail(), r.User.GetPassword(), r.GetRole())
}

func (q *Query) CreateComment(r *pb.CreateCommentRequest) string {
	return fmt.Sprintf("INSERT INTO comments SET content=%q, user_id=%q, post_id=%d", r.GetContent(), r.GetUserId(), r.GetPostId())
}

func (q *Query) GetPostComments(r *pb.PostRequest) string {
	return fmt.Sprintf("SELECT id, content, created, last_edited, user_id, post_id FROM comments WHERE post_id=%d", r.GetId())
}

func (q *Query) DeletePost(r *pb.PostRequest) string {
	return fmt.Sprintf("DELETE FROM posts WHERE id=%d", r.GetId())
}

func (q *Query) CreatePost(r *pb.CreatePostWithSlug) string {
	return fmt.Sprintf("INSERT INTO posts SET slug=%q, title=%q, content=%q", r.GetSlug(), r.Post.GetTitle(), r.Post.GetContent())
}

func (q *Query) GetPost(r *pb.PostRequest) string {
	return fmt.Sprintf("SELECT id, title, content, created, last_edited, published, slug FROM posts WHERE id=%d", r.GetId())
}
