.PHONY: help install dev-api build-api build-api-linux gen-api lint test tidy

help:
	@echo "常用命令:"
	@echo "  make install         安装 JS/TS 依赖"
	@echo "  make dev-api         启动 Go API（config_dev.yaml）"
	@echo "  make build-api       本地编译 API"
	@echo "  make build-api-linux 交叉编译 Linux amd64"
	@echo "  make dev-admin       启动后台管理"
	@echo "  make gen-api         根据 OpenAPI 生成代码"
	@echo "  make lint            运行 lint"
	@echo "  make test            运行测试"
	@echo "  make tidy            go mod tidy"

install:
	npm install

dev-api:
	cd services/api && GOWORK=off go run -tags=jsoniter ./cmd/app -c config_dev.yaml

build-api:
	cd services/api && GOWORK=off go build -tags=jsoniter -o app ./cmd/app

build-api-linux:
	cd services/api && GOWORK=off CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=jsoniter -o app ./cmd/app

dev-admin:
	npm run dev -w @jobs/admin

gen-api:
	./tooling/scripts/gen-api.sh

lint:
	npm run lint --workspaces --if-present
	cd services/api && GOWORK=off go vet -tags=jsoniter ./...

test:
	npm test --workspaces --if-present
	cd services/api && GOWORK=off go test -tags=jsoniter ./...

tidy:
	cd services/api && GOWORK=off go mod tidy
