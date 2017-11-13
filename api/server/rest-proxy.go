package main

import (
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/thurt/demo-blog-platform/api/proto"
)

const (
	GrpcHost = "localhost:10000"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterCmsHandlerFromEndpoint(ctx, mux, GrpcHost, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8080", mux)
}

func ProxyStartup() {
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
