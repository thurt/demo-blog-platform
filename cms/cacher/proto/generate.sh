#! /bin/bash

SERVICE=./cacher.proto

#generate ${SERVICE}.pb.go
protoc \
    -I/usr/local/include \
    -I. \
    -I$GOPATH/src \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --go_out=plugins=grpc:. \
    ${SERVICE} && \

#generate Cacher mocks
mockgen github.com/thurt/demo-blog-platform/cms/cacher/proto CacherServer > ../mock_proto/mock_proto.go
