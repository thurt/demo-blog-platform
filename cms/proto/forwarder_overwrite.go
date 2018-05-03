package cms

import (
	"fmt"
	"net/http"

	"context"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

func maybeRespondNotModified(req *http.Request, w http.ResponseWriter, etag string) bool {
	// create an etag
	wetag := fmt.Sprintf("W/%q", etag)
	w.Header().Add("ETag", wetag)

	// check whether the request provides an etag
	inm := req.Header.Get("If-None-Match")
	if inm == wetag {
		w.WriteHeader(http.StatusNotModified)
		return true
	}
	return false
}

func GetPost(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, req *http.Request, resp proto.Message, opts ...func(context.Context, http.ResponseWriter, proto.Message) error) {
	// add cache-control rules for this proxy endpoint
	w.Header().Add("Cache-Control", "no-cache,private")
	if maybeRespondNotModified(req, w, resp.(*Post).GetLastEdited()) == false {
		runtime.ForwardResponseMessage(ctx, mux, marshaler, w, req, resp, opts...)
	}
}

func GetUser(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, req *http.Request, resp proto.Message, opts ...func(context.Context, http.ResponseWriter, proto.Message) error) {
	// add cache-control rules for this proxy endpoint
	w.Header().Add("Cache-Control", "no-cache,public")
	if maybeRespondNotModified(req, w, resp.(*User).GetLastActive()) == false {
		runtime.ForwardResponseMessage(ctx, mux, marshaler, w, req, resp, opts...)
	}
}

func GetComment(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, req *http.Request, resp proto.Message, opts ...func(context.Context, http.ResponseWriter, proto.Message) error) {
	// add cache-control rules for this proxy endpoint
	w.Header().Add("Cache-Control", "public,max-age=60")
	if maybeRespondNotModified(req, w, resp.(*Comment).GetLastEdited()) == false {
		runtime.ForwardResponseMessage(ctx, mux, marshaler, w, req, resp, opts...)
	}
}

func AuthUser(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, req *http.Request, resp proto.Message, opts ...func(context.Context, http.ResponseWriter, proto.Message) error) {
	// add cache-control rules for this proxy endpoint
	w.Header().Add("Cache-Control", "no-store")
	runtime.ForwardResponseMessage(ctx, mux, marshaler, w, req, resp, opts...)
}

func Setup(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, req *http.Request, resp proto.Message, opts ...func(context.Context, http.ResponseWriter, proto.Message) error) {
	// add cache-control rules for this proxy endpoint
	w.Header().Add("Cache-Control", "no-store")
	runtime.ForwardResponseMessage(ctx, mux, marshaler, w, req, resp, opts...)
}

func IsSetup(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, req *http.Request, resp proto.Message, opts ...func(context.Context, http.ResponseWriter, proto.Message) error) {
	// add cache-control rules for this proxy endpoint
	w.Header().Add("Cache-Control", "no-store")
	runtime.ForwardResponseMessage(ctx, mux, marshaler, w, req, resp, opts...)
}

func init() {
	forward_Cms_GetPost_0 = GetPost
	forward_Cms_GetPostBySlug_0 = GetPost
	forward_Cms_GetUser_0 = GetUser
	forward_Cms_GetComment_0 = GetComment

	forward_Cms_AuthUser_0 = AuthUser
	forward_Cms_Setup_0 = Setup
	forward_Cms_IsSetup_0 = IsSetup
}
