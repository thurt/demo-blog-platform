package query

import (
	"os"
	"testing"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/thurt/demo-blog-platform/cms/proto"

	"github.com/google/gofuzz"
	"github.com/xwb1989/sqlparser"
)

func checkSyntax(sql string, t *testing.T) {
	var err error
	_, err = sqlparser.ParseStrictDDL(sql)
	if err != nil {
		t.Error(t.Name(), err, sql)
	}
}

var (
	f *fuzz.Fuzzer
	q *Query
)

func TestMain(m *testing.M) {
	// create the fuzzer
	f = fuzz.New()

	// create the Query struct
	q = &Query{}

	os.Exit(m.Run())
}

func TestCreateComment(t *testing.T) {
	ccr := &pb.CreateCommentRequest{}
	f.Fuzz(ccr)
	checkSyntax(q.CreateComment(ccr), t)
}
func TestCreatePost(t *testing.T) {
	cpr := &pb.CreatePostRequest{}
	f.Fuzz(cpr)
	checkSyntax(q.CreatePost(cpr), t)
}
func TestCreateUser(t *testing.T) {
	cuwr := &pb.CreateUserWithRole{}
	f.Fuzz(cuwr)
	checkSyntax(q.CreateUser(cuwr), t)
}
func TestDeleteComment(t *testing.T) {
	cr := &pb.CommentRequest{}
	f.Fuzz(cr)
	checkSyntax(q.DeleteComment(cr), t)
}
func TestDeletePost(t *testing.T) {
	pr := &pb.PostRequest{}
	f.Fuzz(pr)
	checkSyntax(q.DeletePost(pr), t)
}
func TestDeleteUser(t *testing.T) {
	ur := &pb.UserRequest{}
	f.Fuzz(ur)
	checkSyntax(q.DeleteUser(ur), t)
}
func TestGetComment(t *testing.T) {
	cr := &pb.CommentRequest{}
	f.Fuzz(cr)
	checkSyntax(q.GetComment(cr), t)
}
func TestGetComments(t *testing.T) {
	e := &empty.Empty{}
	f.Fuzz(e)
	checkSyntax(q.GetComments(e), t)
}
func TestGetPost(t *testing.T) {
	pr := &pb.PostRequest{}
	f.Fuzz(pr)
	checkSyntax(q.GetPost(pr), t)
}
func TestGetPostComments(t *testing.T) {
	pr := &pb.PostRequest{}
	f.Fuzz(pr)
	checkSyntax(q.GetPostComments(pr), t)
}
func TestGetUser(t *testing.T) {
	ur := &pb.UserRequest{}
	f.Fuzz(ur)
	checkSyntax(q.GetUser(ur), t)
}
func TestGetUserComments(t *testing.T) {
	ur := &pb.UserRequest{}
	f.Fuzz(ur)
	checkSyntax(q.GetUserComments(ur), t)
}
func TestPublishPost(t *testing.T) {
	pr := &pb.PostRequest{}
	f.Fuzz(pr)
	checkSyntax(q.PublishPost(pr), t)
}
func TestUnPublishPost(t *testing.T) {
	pr := &pb.PostRequest{}
	f.Fuzz(pr)
	checkSyntax(q.UnPublishPost(pr), t)
}
func TestUpdateComment(t *testing.T) {
	ucr := &pb.UpdateCommentRequest{}
	f.Fuzz(ucr)
	checkSyntax(q.UpdateComment(ucr), t)
}
func TestUpdatePost(t *testing.T) {
	upr := &pb.UpdatePostRequest{}
	f.Fuzz(upr)
	checkSyntax(q.UpdatePost(upr), t)
}
func TestAdminExists(t *testing.T) {
	e := &empty.Empty{}
	f.Fuzz(e)
	checkSyntax(q.AdminExists(e), t)
}
func TestGetPosts(t *testing.T) {
	checkSyntax(q.GetPosts(), t)
}
func TestGetPublishedPosts(t *testing.T) {
	checkSyntax(q.GetPublishedPosts(), t)
}
func TestGetUserPassword(t *testing.T) {
	ur := &pb.UserRequest{}
	f.Fuzz(ur)
	checkSyntax(q.GetUserPassword(ur), t)
}
