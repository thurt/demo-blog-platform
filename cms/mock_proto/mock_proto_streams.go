package mock_proto

import (
	"context"

	pb "github.com/thurt/demo-blog-platform/cms/proto"
	"google.golang.org/grpc"
)

type mockCms_GetPostsServer struct {
	grpc.ServerStream
	Results []*pb.Post
	pos     int
	nextErr map[int]error
	ctx     context.Context
}

func (m *mockCms_GetPostsServer) SetSendError(pos int, err error) *mockCms_GetPostsServer {
	m.nextErr[pos] = err
	return m
}

func (m *mockCms_GetPostsServer) Send(p *pb.Post) error {
	if err, ok := m.nextErr[m.pos]; ok {
		return err
	}

	m.Results = append(m.Results, p)

	m.pos++
	return nil
}

func (m *mockCms_GetPostsServer) Context() context.Context {
	return m.ctx
}

func (m *mockCms_GetPostsServer) SetContext(ctx context.Context) {
	m.ctx = ctx
}

func NewMockCms_GetPostsServer() *mockCms_GetPostsServer {
	return &mockCms_GetPostsServer{nextErr: make(map[int]error)}
}

type MockCms_GetPostCommentsServer struct {
	grpc.ServerStream
	Results []*pb.Comment
}

func (m *MockCms_GetPostCommentsServer) Send(c *pb.Comment) error {
	m.Results = append(m.Results, c)
	return nil
}

type MockCms_GetCommentsServer struct {
	grpc.ServerStream
	Results []*pb.Comment
}

func (m *MockCms_GetCommentsServer) Send(c *pb.Comment) error {
	m.Results = append(m.Results, c)
	return nil
}

type MockCms_GetUserCommentsServer struct {
	grpc.ServerStream
	Results []*pb.Comment
}

func (m *MockCms_GetUserCommentsServer) Send(c *pb.Comment) error {
	m.Results = append(m.Results, c)
	return nil
}
