#! /bin/bash

SERVICE=./cms.proto

#generate ${SERVICE}.pb.go
protoc \
    -I/usr/local/include \
    -I. \
    -I$GOPATH/src \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --go_out=plugins=grpc:. \
    ${SERVICE} && \

#generate ${SERVICE}.validator.pb.go
protoc \
    -I/usr/local/include \
    -I. \
    -I$GOPATH/src \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --govalidators_out=logtostderr=true:. \
    ${SERVICE} && \

#generate ${SERVICE}.pb.gw.go
protoc \
    -I/usr/local/include \
    -I. \
    -I$GOPATH/src \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
     --grpc-gateway_out=logtostderr=true:. \
     ${SERVICE} && \

#generate ${SERVICE}.swagger.json
protoc \
    -I/usr/local/include \
    -I. \
    -I$GOPATH/src \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
     --swagger_out=logtostderr=true:. \
     ${SERVICE} && \

#generate Cms mocks
mockgen github.com/thurt/demo-blog-platform/cms/proto CmsServer > ../mock_proto/mock_proto.go && \
mockgen github.com/thurt/demo-blog-platform/cms/proto CmsAuthServer > ../mock_proto/mock_proto_auth.go
