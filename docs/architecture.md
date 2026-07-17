# 架构概览

本仓库是多语言 monorepo：

- `apps/`：前端应用（微信小程序、后台管理）
- `services/`：后端服务（Go）
- `packages/`：JS/TS 共享包
- `contracts/openapi/`：跨语言 API 契约

## 依赖方向

- 小程序 / 后台 → `packages/shared`、`packages/api-client`
- 小程序 / 后台 → `services/api`（HTTP）
- `contracts/openapi` → 生成 Go / TS 客户端（`make gen-api`）

## 技术栈

- 前端 JS/TS：npm workspaces（`apps/*`、`packages/*`）
- 后端 `services/api`：Go 1.26 + Gin + **GORM**（`gorm.io/gorm`）。数据访问统一走 `db.Default()`（`*gorm.DB`），已移除历史 xorm 依赖。

## 本地端口

| 服务 | 默认地址 |
|------|----------|
| API | `http://127.0.0.1:8080`（`services/api`，原 job-server） |
| Admin | `http://127.0.0.1:5173`（`/api` 代理到 API） |

API 启动：`make dev-api`（读取 `services/api/config_dev.yaml`）。
