# Jobs Monorepo

多语言 monorepo：微信小程序、Go API（原 job-server）、后台管理（Vite + React）。

## 目录

```text
apps/miniapp      微信小程序（原 wechat-mini）
apps/admin        后台管理
services/api      Go 主 API（原 job-server，module: app）
packages/*        JS/TS 共享包
contracts/openapi API 契约
tooling/scripts   仓库脚本
```

## 前置要求

- Node.js >= 20（见 `.nvmrc`）
- 包管理使用 **npm**（workspaces + `package-lock.json`），不使用 pnpm / yarn
- Go >= 1.26
- 微信开发者工具（打开 `apps/miniapp`）

## 快速开始

```bash
make install
make dev-api      # 使用 services/api/config_dev.yaml，默认 :8080
make dev-admin    # :5173
```

小程序：用微信开发者工具导入 `apps/miniapp`。

API 接口说明见 `services/api/README.md`。

## 常用命令

| 命令 | 说明 |
|------|------|
| `make install` | 安装 JS/TS 依赖 |
| `make dev-api` | 启动 API（`-c config_dev.yaml`） |
| `make build-api` | 编译 API 到 `services/api/app` |
| `make build-api-linux` | 交叉编译 Linux amd64 |
| `make dev-admin` | 启动后台 |
| `make gen-api` | OpenAPI 代码生成（占位脚本） |
| `make lint` / `make test` | lint / 测试 |
| `make tidy` | `go mod tidy` |

也可用 npm：`npm run dev:admin`、`npm run gen:api` 等。

## 约定

- JS/TS：**npm** workspaces（统一用 `npm`，`npm install` / `npm run -w <pkg>`；提交 `package-lock.json`）
- Go：`services/api` 使用 Go 1.26 + modules；构建时 `GOWORK=off` 避免与根目录 `go.work` 冲突（`vendor/` 不入库）
- ORM：统一使用 **GORM**（`gorm.io/gorm`）。`db.Default()` 返回 `*gorm.DB`；单条查询用 `db.Get(tx, &dest)` 保留旧的 `(存在?, error)` 语义。模型代码生成见 `cmd/genmodel`
- 跨语言任务：Makefile
- 接口契约：`contracts/openapi/openapi.yaml`
