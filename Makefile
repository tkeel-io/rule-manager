GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
API_PROTO_FILES=$(shell find api -name *.proto)
DOCKERTAG?=tkeelio/rule-manager:dev
BINNAME = rule-manager

GOCMD = GO111MODULE=on go

GIT_BRANCH=$(shell git symbolic-ref --short -q HEAD)
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
BUILD_DATE=$(shell date '+%Y-%m-%d-%H:%M:%S')
GOBUILD = $(GOCMD) build
.PHONY: init
# init env
init:
	go get -d -u  github.com/tkeel-io/tkeel-interface/openapi
	go get -d -u  github.com/tkeel-io/kit
	go get -d -u  github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.7.0

	go install  github.com/tkeel-io/tkeel-interface/tool/cmd/artisan@latest
	go install  google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
	go install  google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0
	go install  github.com/tkeel-io/tkeel-interface/protoc-gen-go-http@latest
	go install  github.com/tkeel-io/tkeel-interface/protoc-gen-go-errors@latest
	go install  github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.7.0

.PHONY: api
# generate api proto
api:
	protoc --proto_path=. \
	       --proto_path=./third_party \
	       --proto_path=./api/rule/v1 \
 	       --go_out=paths=source_relative:. \
 	       --go-http_out=paths=source_relative:. \
 	       --go-grpc_out=paths=source_relative:. \
 	       --go-errors_out=paths=source_relative:. \
 	       --openapiv2_out=./api/ \
		   --openapiv2_opt=allow_merge=true \
 	       --openapiv2_opt=logtostderr=true \
 	       --openapiv2_opt=json_names_for_fields=false \
	       $(API_PROTO_FILES)

.PHONY: build
# build
build:
	@rm -rf bin/
	@mkdir bin/
	@echo "---------------------------"
	@echo "-        build...         -"
	@$(GOBUILD)    -o bin/$(BINNAME) cmd/rule-manager/main.go
	@echo "-     build(linux)...     -"
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64  $(GOBUILD) -o bin/linux/$(BINNAME) cmd/rule-manager/main.go
	@echo "-    builds completed!    -"
	@echo "---------------------------"

.PHONY: generate
# generate
generate:
	go generate ./...


.PHONY: all
# generate all
all:
	make api;
	make generate;

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

docker-build:
	docker build -t $(DOCKERTAG) .
docker-push:
	docker push $(DOCKERTAG)
docker-auto:
	docker build -t $(DOCKERTAG) .
	docker push $(DOCKERTAG)
pretest:
	make build
	make docker-auto
