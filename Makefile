.PHONY: compile

CMS_SERVICE=./cms/proto/cms.proto
CACHER_SERVICE=./cms/cacher/proto/cacher.proto

authentication:
	#generate mc.Conn interface
	ifacemaker \
		-f ${GOPATH}/src/github.com/memcachier/mc/mc.go \
		-s Conn -i Conn -p authentication \
		-o ./cms/authentication/mc.go
	#generate mc.Conn mock
	mockgen -package=mock_mc -source=./authentication/mc.go >  ./cms/mock_mc/mock_mc.go 

cacher:
	protoc \
		-I/usr/local/include \
		-I./cms/cacher/proto \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:./cms/cacher/proto \
		${CACHER_SERVICE}

cacher_mock: cacher
	mockgen github.com/thurt/demo-blog-platform/cms/cacher/proto CacherServer > ./cms/cacher/mock_proto/mock_proto.go

cacher_all: cacher cacher_mock

cms:
	protoc \
		-I/usr/local/include \
		-I./cms/proto \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
		--go_out=plugins=grpc:./cms/proto \
		${CMS_SERVICE}

cms_validator:
	protoc \
		-I/usr/local/include \
		-I./cms/proto \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
		--govalidators_out=logtostderr=true:./cms/proto \
		${CMS_SERVICE}

cms_gateway:
	protoc \
		-I/usr/local/include \
		-I./cms/proto \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
		 --grpc-gateway_out=logtostderr=true:./cms/proto \
		 ${CMS_SERVICE}

cms_swagger:
	protoc \
		-I/usr/local/include \
		-I./cms/proto \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
		 --swagger_out=logtostderr=true:./cms \
		 ${CMS_SERVICE}

cms_mocks:
	mockgen github.com/thurt/demo-blog-platform/cms/proto CmsServer > cms/mock_proto/mock_proto.go
	mockgen github.com/thurt/demo-blog-platform/cms/proto CmsAuthServer > cms/mock_proto/mock_proto_auth.go
	mockgen github.com/thurt/demo-blog-platform/cms/proto HasherServer > cms/mock_proto/mock_hasher.go
	mockgen github.com/thurt/demo-blog-platform/cms/proto EmailerServer > cms/mock_proto/mock_emailer.go

cms_all: cms cms_gateway cms_validator cms_swagger cms_mocks

domain: cms
	mockgen github.com/thurt/demo-blog-platform/cms/domain Provider > ./cms/mock_domain/mock_domain.go

secrets:
	tar cvf secrets.tar client-secret.json cms/app.yaml
	travis encrypt-file secrets.tar --add
