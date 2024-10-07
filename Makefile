.PHONY: build clean test test-race
# 支持参数

VERSION=0.0.1
BIN=myddns
DIR_SRC=.
DOCKER_CMD=docker
# 系统环境参数

GO_ENV=CGO_ENABLED=0
GO_FLAGS=-ldflags="-X main.version=$(VERSION) -X 'main.buildTime=`date`' -extldflags -static"
GO=$(GO_ENV) $(shell which go)
GOROOT=$(shell `which go` env GOROOT)
GOPATH=$(shell `which go` env GOPATH)
# go语言环境参数

build: $(DIR_SRC)/main.go
	@$(GO) build $(GO_FLAGS) -o $(BIN) $(DIR_SRC)
# 编译

build_docker_image: 
	@$(DOCKER_CMD) build -f ./Dockerfile -t ddns-go:$(VERSION) .
# 构建docker镜像

init:
	@go get -u github.com/go-bindata/go-bindata/...
# 初始化

test:
	@$(GO) test ./...
# 自动测试

test-race:
	@$(GO) test -race ./...

bindata:
	@go-bindata -pkg util -o util/staticPages.go static/pages/...
	@go-bindata -pkg asserts -o asserts/html.go -prefix "static/" static/

# 嵌套静态资源到go代码中
dev:
	@go-bindata -debug -pkg util -o util/staticPages.go static/pages/...
# 调试版本静态资源嵌套

# clean all build result
clean:
	@$(GO) clean ./...
	@rm -f $(BIN)
	@rm -rf ./dist/*
# 清理