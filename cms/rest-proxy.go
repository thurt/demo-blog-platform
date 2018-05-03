package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"

	"context"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	pb "github.com/thurt/demo-blog-platform/cms/proto"
	"github.com/thurt/reverseProxyHostRewrite"
)

var GrpcHost string
var STATICHOST_ADDRESS string

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	CORSOptions := []handlers.CORSOption{
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "PUT", "POST", "DELETE"}),
	}

	cms_mux := runtime.NewServeMux()
	cms_mux_wrapper := mux.NewRouter()
	cms_mux_wrapper.Path("/_swagger").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./cms.swagger.json")
	})
	cms_mux_wrapper.PathPrefix("/").Handler(cms_mux)

	opts := []grpc.DialOption{grpc.WithInsecure()}
	GrpcHost = fmt.Sprintf("localhost:%d", PORT)
	err := pb.RegisterCmsHandlerFromEndpoint(ctx, cms_mux, GrpcHost, opts)
	if err != nil {
		return err
	}
	cms_mux_cors := handlers.CORS(CORSOptions...)(cms_mux_wrapper)

	router_mux := mux.NewRouter()

	STATICHOST_ADDRESS = os.Getenv("STATICHOST_ADDRESS")
	statichost, err := url.Parse(STATICHOST_ADDRESS)
	if err != nil {
		panic(err)
	}

	log.Println("rest proxy is using static host:", STATICHOST_ADDRESS)

	proxy_statichost := reverseProxyHostRewrite.New(statichost)

	var mainPagePath string
	// this host does not serve index.html automatically when requesting the root.
	// it acts more like a fileserver rather than a webserver in this respect, so
	// the path to mainPage must be absolute.
	if statichost.Host == "storage.googleapis.com" {
		mainPagePath = "/index.html"
	} else { // assume other hosts act like regular webserver, serving index.html automatically when requesting root
		mainPagePath = "/"
	}

	router_mux.PathPrefix("/api/").Handler(http.StripPrefix("/api", cms_mux_cors))
	router_mux.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// when there is no extension (ex .png, .html) at the end of the request path,
		// reroute it to the mainPagePath. This allows the SPA (located at mainPagePath) to
		// decide how to handle all the path requests at the application level. meanwhile
		// request paths w/ extension will skip this block (assuming these are static files
		// located on statichost)
		if ext := path.Ext(r.URL.Path); ext == "" {
			r.URL.Path = mainPagePath
		}

		proxy_statichost.ServeHTTP(w, r)
	})

	return http.ListenAndServe(":8080", router_mux)
}

func ProxyServe() error {
	return run()
}
