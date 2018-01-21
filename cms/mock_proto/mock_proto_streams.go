package mock_proto

import (
	pb "github.com/thurt/demo-blog-platform/cms/proto"
	"google.golang.org/grpc"
)

type mockCms_GetPostsServer struct {
	grpc.ServerStream
	Results []*pb.Post
	pos     int
	nextErr map[int]error
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

func NewMockCms_GetPostsServer() *mockCms_GetPostsServer {
	return &mockCms_GetPostsServer{nextErr: make(map[int]error)}
}

type mockCms_GetPostCommentsServer struct {
	grpc.ServerStream
	Results []*pb.Comment
}

func (m *mockCms_GetPostCommentsServer) Send(c *pb.Comment) error {
	m.Results = append(m.Results, c)
	return nil
}

type mockCms_GetCommentsServer struct {
	grpc.ServerStream
	Results []*pb.Comment
}

func (m *mockCms_GetCommentsServer) Send(c *pb.Comment) error {
	m.Results = append(m.Results, c)
	return nil
}

type mockCms_GetUserCommentsServer struct {
	grpc.ServerStream
	Results []*pb.Comment
}

func (m *mockCms_GetUserCommentsServer) Send(c *pb.Comment) error {
	m.Results = append(m.Results, c)
	return nil
}
