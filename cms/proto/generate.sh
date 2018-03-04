#! /bin/bash

SERVICE=./cms.proto

#generate ${SERVICE}.pb.go
protoc \
    -I/usr/local/include \
    -I. \
    -I$GOPATH/src \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
    --go_out=plugins=grpc:. \
    ${SERVICE} && \

#generate ${SERVICE}.validator.pb.go
protoc \
    -I/usr/local/include \
    -I. \
    -I$GOPATH/src \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
    --govalidators_out=logtostderr=true:. \
    ${SERVICE} && \

#generate ${SERVICE}.pb.gw.go
protoc \
    -I/usr/local/include \
    -I. \
    -I$GOPATH/src \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
     --grpc-gateway_out=logtostderr=true:. \
     ${SERVICE} && \

#generate ${SERVICE}.swagger.json
protoc \
    -I/usr/local/include \
    -I. \
    -I$GOPATH/src \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
     --swagger_out=logtostderr=true:../ \
     ${SERVICE} && \

#generate Cms mocks
mockgen github.com/thurt/demo-blog-platform/cms/proto CmsServer > ../mock_proto/mock_proto.go && \
mockgen github.com/thurt/demo-blog-platform/cms/proto CmsAuthServer > ../mock_proto/mock_proto_auth.go && \
mockgen github.com/thurt/demo-blog-platform/cms/proto HasherServer > ../mock_proto/mock_hasher.go && \
mockgen github.com/thurt/demo-blog-platform/cms/proto EmailerServer > ../mock_proto/mock_emailer.go
