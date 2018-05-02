.PHONY: 

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

CACHER_PROTO=./cms/cacher/proto/cacher.proto
CACHER_GO=./cms/cacher/proto/cacher.pb.go
CACHER_MOCK=./cms/cacher/mock_proto/mock_proto.go

DOMAIN_GO=./cms/domain/domain.go
DOMAIN_MOCK=./cms/mock_domain/mock_domain.go

#for deployment
SECRETS=./secrets.tar.enc
CLIENT_SECRET=./client-secret.json
APP_YAML=./cms/app.yaml

#memcachier
MC_GO=$(GOPATH)/src/github.com/thurt/demo-blog-platform/vendor/memcachier/mc/mc.go
MC_IFACE=./cms/authentication/mc.go
MC_MOCK=./cms/mock_mc/mock_mc.go

$(MC_IFACE): $(MC_GO)
	#generate mc.Conn interface
	ifacemaker \
		-f $(MC_GO) \
		-s Conn -i Conn -p authentication \
		-o $(MC_IFACE) 

$(MC_MOCK): $(MC_IFACE)
	#generate mc.Conn mock
	mockgen -package=mock_mc -source=$(MC_IFACE) >  $(MC_MOCK) 

$(CACHER_GO): $(CACHER_PROTO)
	protoc \
		-I/usr/local/include \
		-I./cms/cacher/proto \
		-I${GOPATH}/src \
		-I$(GOOGLEAPIS_DIR) \
		--go_out=plugins=grpc:./cms/cacher/proto \
		${CACHER_PROTO}

$(CACHER_MOCK): $(CACHER_GO) 
	mockgen github.com/thurt/demo-blog-platform/cms/cacher/proto CacherServer > $(CACHER_MOCK) 

cacher: $(CACHER_GO) $(CACHER_MOCK)

$(CMS_GO): $(CMS_PROTO)
	protoc \
		-I/usr/local/include \
		-I./cms/proto \
		-I${GOPATH}/src \
		-I$(GOOGLEAPIS_DIR) \
		-I$(GRPCGATEWAY_DIR) \
		--go_out=plugins=grpc:./cms/proto \
		${CMS_PROTO}

$(CMS_VALIDATOR): $(CMS_PROTO)
	protoc \
		-I/usr/local/include \
		-I./cms/proto \
		-I${GOPATH}/src \
		-I$(GOOGLEAPIS_DIR) \
		-I$(GRPCGATEWAY_DIR) \
		--govalidators_out=logtostderr=true:./cms/proto \
		${CMS_PROTO}

$(CMS_GATEWAY): $(CMS_PROTO)
	protoc \
		-I/usr/local/include \
		-I./cms/proto \
		-I${GOPATH}/src \
		-I$(GOOGLEAPIS_DIR) \
		-I$(GRPCGATEWAY_DIR) \
		 --grpc-gateway_out=logtostderr=true:./cms/proto \
		 ${CMS_PROTO}

$(CMS_SWAGGER): $(CMS_PROTO) 
	protoc \
		-I/usr/local/include \
		-I./cms/proto \
		-I${GOPATH}/src \
		-I$(GOOGLEAPIS_DIR) \
		-I$(GRPCGATEWAY_DIR) \
		 --swagger_out=logtostderr=true:./cms \
		 ${CMS_PROTO}

$(CMS_MOCK_HASHER): $(CMS_GO)
	mockgen $(CMS_PROTO_PACKAGE) HasherServer > $(CMS_MOCK_HASHER) 

$(CMS_MOCK_EMAILER): $(CMS_GO)
	mockgen $(CMS_PROTO_PACKAGE) EmailerServer > $(CMS_MOCK_EMAILER)

$(CMS_MOCK_AUTH): $(CMS_GO)
	mockgen $(CMS_PROTO_PACKAGE) CmsAuthServer > $(CMS_MOCK_AUTH)

$(CMS_MOCK): $(CMS_GO)
	mockgen $(CMS_PROTO_PACKAGE) CmsServer > $(CMS_MOCK)

cms: $(CMS_PROTO) $(CMS_GO) $(CMS_VALIDATOR) $(CMS_GATEWAY) $(CMS_SWAGGER) $(CMS_MOCK) $(CMS_MOCK_HASHER) $(CMS_MOCK_EMAILER) $(CMS_MOCK_AUTH)

$(DOMAIN_MOCK): $(DOMAIN_GO) 
	mockgen github.com/thurt/demo-blog-platform/cms/domain Provider > $(DOMAIN_MOCK)

domain: $(DOMAIN_MOCK)

$(SECRETS): $(CLIENT_SECRET) $(APP_YAML)
	tar cvf secrets.tar $(CLIENT_SECRET) $(APP_YAML) 
	travis encrypt-file secrets.tar --add
