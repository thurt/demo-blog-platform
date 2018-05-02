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
	return fmt.Sprintf("UPDATE posts SET slug=%q, title=%q, content=%q, published=%t WHERE id=%d", r.GetSlug(), r.Post.GetTitle(), r.Post.GetContent(), r.Post.GetPublished(), r.Post.GetId())
}

func (q *Query) UpdateComment(r *pb.UpdateCommentRequest) string {
	return fmt.Sprintf("UPDATE comments SET content=%q WHERE id=%d", r.GetContent(), r.GetId())
}

func (q *Query) GetUserComments(r *pb.UserRequest) string {
	return fmt.Sprintf("SELECT id, content, created, last_edited, user_id, post_id FROM comments WHERE user_id=%q ORDER BY created DESC", r.GetId())
}

func (q *Query) GetPosts() string {
	return "SELECT id, title, content, created, last_edited, slug FROM published_posts ORDER BY created DESC"
}

func (q *Query) GetUnpublishedPosts() string {
	return "SELECT id, title, content, created, last_edited, slug FROM posts ORDER BY created DESC"
}

func (q *Query) GetComments(_ *empty.Empty) string {
	return "SELECT id, content, created, last_edited, user_id, post_id FROM comments ORDER BY created DESC"
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
	return fmt.Sprintf("SELECT id, content, created, last_edited, user_id, post_id FROM comments WHERE post_id=%d ORDER BY created DESC", r.GetId())
}

func (q *Query) DeletePost(r *pb.PostRequest) string {
	return fmt.Sprintf("DELETE FROM posts WHERE id=%d", r.GetId())
}

func (q *Query) CreatePost(r *pb.CreatePostWithSlug) string {
	return fmt.Sprintf("INSERT INTO posts SET slug=%q, title=%q, content=%q", r.GetSlug(), r.Post.GetTitle(), r.Post.GetContent())
}

func (q *Query) GetPost(r *pb.PostRequest) string {
	return fmt.Sprintf("SELECT id, title, content, created, last_edited, slug FROM published_posts WHERE id=%d", r.GetId())
}

func (q *Query) GetPostBySlug(r *pb.PostBySlugRequest) string {
	return fmt.Sprintf("SELECT id, title, content, created, last_edited, slug FROM published_posts WHERE slug=%q", r.GetSlug())
}

func (q *Query) GetUnpublishedPost(r *pb.PostRequest) string {
	return fmt.Sprintf("SELECT id, title, content, created, last_edited, slug FROM posts WHERE id=%d", r.GetId())
}

func (q *Query) GetUnpublishedPostBySlug(r *pb.PostBySlugRequest) string {
	return fmt.Sprintf("SELECT id, title, content, created, last_edited, slug FROM posts WHERE slug=%q", r.GetSlug())
}

func (q *Query) UpdateUserLastActive(r *pb.UserRequest) string {
	return fmt.Sprintf("UPDATE users SET last_active=CURRENT_TIMESTAMP WHERE id=%q", r.GetId())
}
