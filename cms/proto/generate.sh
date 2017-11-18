#! /bin/bash

SERVICE=./cms.proto

#generate ${SERVICE}.pb.go
protoc \
    -I/usr/local/include \
    -I. \
    -I$GOPATH/src \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --go_out=plugins=grpc:. \
    ${SERVICE}

#generate ${SERVICE}.pb.gw.go
protoc \
    -I/usr/local/include \
    -I. \
    -I$GOPATH/src \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
     --grpc-gateway_out=logtostderr=true:. \
     ${SERVICE}

#generate ${SERVICE}.swagger.json
protoc \
    -I/usr/local/include \
    -I. \
    -I$GOPATH/src \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
     --swagger_out=logtostderr=true:. \
     ${SERVICE}
