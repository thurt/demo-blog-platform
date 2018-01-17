package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/thurt/demo-blog-platform/cms/proto"
)

var GrpcHost string
var STATICHOST_ADDRESS string

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	CORSOptions := []handlers.CORSOption{
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
		handlers.AllowedMethods([]string{"GET", "PUT", "POST", "DELETE"}),
	}

	cms_mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	GrpcHost = fmt.Sprintf("localhost:%d", PORT)
	err := pb.RegisterCmsHandlerFromEndpoint(ctx, cms_mux, GrpcHost, opts)
	if err != nil {
		return err
	}
	cms_mux_cors := handlers.CORS(CORSOptions...)(cms_mux)

	router_mux := mux.NewRouter()

	STATICHOST_ADDRESS = os.Getenv("STATICHOST_ADDRESS")
	statichost, err := url.Parse(STATICHOST_ADDRESS)
	if err != nil {
		panic(err)
	}

	proxy_statichost := httputil.NewSingleHostReverseProxy(statichost)
	router_mux.PathPrefix("/api/").Handler(http.StripPrefix("/api/", cms_mux_cors))
	router_mux.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if ext := path.Ext(r.URL.Path); ext == "" {
			r.URL.Path = "/index.html"
		}

		proxy_statichost.ServeHTTP(w, r)
	})

	return http.ListenAndServe(":8080", router_mux)
}

func ProxyServe() error {
	return run()
}
