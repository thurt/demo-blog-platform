.PHONY: secrets run

VENDOR=$(GOPATH)/src/github.com/thurt/demo-blog-platform/cms/vendor

#build tool for mocks
MOCKGEN=$(GOPATH)/bin/mockgen
$(MOCKGEN):
	go get github.com/golang/mock/gomock
	go install github.com/golang/mock/mockgen

#build tool for protobuf
PROTOC=/usr/local/bin/protoc
$(PROTOC):
	echo "You must first install protoc. Refer to https://github.com/google/protobuf/releases"
	exit 1

PROTOC_GEN_GO=$(GOPATH)/bin/protoc-gen-go
$(PROTOC_GEN_GO):
	go get github.com/golang/protobuf/protoc-gen-go

#runtime required for some things
DOCKER=/usr/bin/docker
$(DOCKER):
	echo "You must first install docker. Refer to https://docs.docker.com/install/"
	exit 1

##############################
###### RUN
##############################
run: build | $(DOCKER_COMPOSE)
	docker-compose up 

##############################
###### BUILD
##############################
DOCKER_COMPOSE=/usr/local/bin/docker-compose
$(DOCKER_COMPOSE):
	echo "You must first install docker-compose. Refer to https://docs.docker.com/compose/install/"
	exit 1

build: $(shell find -path "./cms/*" -name "*.go" -not -path "./cms/vendor/*") authentication cms cacher domain test | $(DOCKER_COMPOSE) $(DOCKER)
	docker-compose build
	@touch build

##############################
###### TEST
##############################
PACKAGES=$(shell go list ./... | grep -v /vendor/) # see @rsc final answer https://github.com/golang/go/issues/11659

test: authentication cms cacher domain
	go test \
		-v \
		$(PACKAGES)
	@touch test

test-integration: authentication cms cacher domain | $(DOCKER)
	go test \
		-v \
		-tags=integration \
		$(PACKAGES)
	@touch test-integration

##############################
###### AUTHENICATION
##############################

#build tool dependency
IFACEMAKER=$(GOPATH)/bin/ifacemaker
$(IFACEMAKER): 
	go get github.com/vburenin/ifacemaker

#memcachier
MC_GO=$(VENDOR)/github.com/memcachier/mc/mc.go
MC_IFACE=./cms/authentication/mc.go
MC_MOCK=./cms/mock_mc/mock_mc.go

$(MC_IFACE): $(MC_GO) | $(IFACEMAKER)
	#generate mc.Conn interface -- must use ifacemaker b/c mc.Conn is a struct, not an interface
	ifacemaker \
		-f $(MC_GO) \
		-s Conn -i Conn -p authentication \
		-o $@

$(MC_MOCK): $(MC_IFACE) | $(MOCKGEN)
	#generate mc.Conn mock
	mockgen -package=mock_mc -source=$(MC_IFACE) >  $@

authentication: $(MC_MOCK) $(MC_IFACE)
	@touch authentication


##############################
###### CACHER
##############################

CACHER_PROTO_PACKAGE=github.com/thurt/demo-blog-platform/cms/cacher/proto
CACHER_PROTO=./cms/cacher/proto/cacher.proto
CACHER_GO=./cms/cacher/proto/cacher.pb.go
CACHER_MOCK=./cms/cacher/mock_proto/mock_proto.go

$(CACHER_GO): $(CACHER_PROTO) | $(PROTOC) $(PROTOC_GEN_GO)
	protoc \
		-I/usr/local/include \
		-I./cms/cacher/proto \
		-I${GOPATH}/src \
		-I$(GOOGLEAPIS_DIR) \
		--go_out=plugins=grpc:./cms/cacher/proto \
		${CACHER_PROTO}

$(CACHER_MOCK): $(CACHER_GO) | $(MOCKGEN)
	mockgen $(CACHER_PROTO_PACKAGE) CacherServer > $@

cacher: $(CACHER_GO) $(CACHER_MOCK)
	@touch cacher


##############################
###### CMS 
##############################

GRPCGATEWAY_DIR=$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway
GOOGLEAPIS_DIR=$(GRPCGATEWAY_DIR)/third_party/googleapis

CMS_PROTO_PACKAGE=github.com/thurt/demo-blog-platform/cms/proto
CMS_PROTO=./cms/proto/cms.proto
CMS_GO=./cms/proto/cms.pb.go
CMS_GATEWAY=./cms/proto/cms.pb.gw.go
CMS_VALIDATOR=./cms/proto/cms.validator.pb.go
CMS_SWAGGER=./cms/cms.swagger.json
CMS_MOCK_AUTH=./cms/mock_proto/mock_proto_auth.go
CMS_MOCK_HASHER=./cms/mock_proto/mock_hasher.go
CMS_MOCK_EMAILER=./cms/mock_proto/mock_emailer.go
CMS_MOCK=./cms/mock_proto/mock_proto.go

#build tools for grpc-ecosystem
PROTOC_GEN_GOVALIDATORS=$(GOPATH)/bin/protoc-gen-govalidators
$(PROTOC_GEN_GOVALIDATORS):
	go get github.com/mwitkow/go-proto-validators
PROTOC_GEN_SWAGGER=$(GOPATH)/bin/protoc-gen-swagger
$(PROTOC_GEN_SWAGGER):
	go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger	
PROTOC_GEN_GRPC_GATEWAY=$(GOPATH)/bin/protoc-gen-grpc-gateway
$(PROTOC_GEN_GRPC_GATEWAY):
	go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

$(CMS_GO): $(CMS_PROTO) | $(PROTOC) $(PROTOC_GEN_GO)
	protoc \
		-I/usr/local/include \
		-I./cms/proto \
		-I${GOPATH}/src \
		-I$(GOOGLEAPIS_DIR) \
		-I$(GRPCGATEWAY_DIR) \
		--go_out=plugins=grpc:./cms/proto \
		${CMS_PROTO}

$(CMS_VALIDATOR): $(CMS_PROTO) | $(PROTOC) $(PROTOC_GEN_GOVALIDATORS)
	protoc \
		-I/usr/local/include \
		-I./cms/proto \
		-I${GOPATH}/src \
		-I$(GOOGLEAPIS_DIR) \
		-I$(GRPCGATEWAY_DIR) \
		--govalidators_out=logtostderr=true:./cms/proto \
		${CMS_PROTO}

$(CMS_GATEWAY): $(CMS_PROTO) | $(PROTOC) $(PROTOC_GEN_GRPC_GATEWAY)
	protoc \
		-I/usr/local/include \
		-I./cms/proto \
		-I${GOPATH}/src \
		-I$(GOOGLEAPIS_DIR) \
		-I$(GRPCGATEWAY_DIR) \
		 --grpc-gateway_out=logtostderr=true:./cms/proto \
		 ${CMS_PROTO}

$(CMS_SWAGGER): $(CMS_PROTO) | $(PROTOC) $(PROTOC_GEN_SWAGGER)
	protoc \
		-I/usr/local/include \
		-I./cms/proto \
		-I${GOPATH}/src \
		-I$(GOOGLEAPIS_DIR) \
		-I$(GRPCGATEWAY_DIR) \
		 --swagger_out=logtostderr=true:./cms \
		 ${CMS_PROTO}

$(CMS_MOCK_HASHER): $(CMS_GO) | $(MOCKGEN)
	mockgen $(CMS_PROTO_PACKAGE) HasherServer > $@

$(CMS_MOCK_EMAILER): $(CMS_GO) | $(MOCKGEN)
	mockgen $(CMS_PROTO_PACKAGE) EmailerServer > $@

$(CMS_MOCK_AUTH): $(CMS_GO) | $(MOCKGEN)
	mockgen $(CMS_PROTO_PACKAGE) CmsAuthServer > $@

$(CMS_MOCK): $(CMS_GO) | $(MOCKGEN)
	mockgen $(CMS_PROTO_PACKAGE) CmsServer > $@

cms: $(CMS_PROTO) $(CMS_GO) $(CMS_VALIDATOR) $(CMS_GATEWAY) $(CMS_SWAGGER) $(CMS_MOCK) $(CMS_MOCK_HASHER) $(CMS_MOCK_EMAILER) $(CMS_MOCK_AUTH)
	@touch cms


##############################
###### DOMAIN 
##############################

DOMAIN_GO=./cms/domain/domain.go
DOMAIN_MOCK=./cms/mock_domain/mock_domain.go

$(DOMAIN_MOCK): $(DOMAIN_GO) | $(MOCKGEN)
	mockgen github.com/thurt/demo-blog-platform/cms/domain Provider > $@

domain: $(DOMAIN_MOCK)
	@touch domain


##############################
###### SECRETS
##############################

RUBY=/usr/bin/ruby
TRAVIS=/usr/local/bin/travis

SECRETS=./secrets.tar.enc
CLIENT_SECRET=./client-secret.json
APP_YAML=./cms/app.yaml

$(RUBY):
	echo "You must first install ruby. Refer to https://www.ruby-lang.org/en/downloads/"
	exit 1

$(TRAVIS):
	gem install travis

$(SECRETS): $(CLIENT_SECRET) $(APP_YAML) | $(RUBY) $(TRAVIS)
	tar cvf secrets.tar $(CLIENT_SECRET) $(APP_YAML) 
	travis encrypt-file secrets.tar --add

secrets: $(SECRETS)

