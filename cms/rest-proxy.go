package main

import (
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/thurt/demo-blog-platform/cms/proto"
)

var GrpcHost string

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	GrpcHost = fmt.Sprintf("localhost:%d", PORT)
	err := pb.RegisterCmsHandlerFromEndpoint(ctx, mux, GrpcHost, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8080", mux)
}

func ProxyServe() error {
	return run()
}
