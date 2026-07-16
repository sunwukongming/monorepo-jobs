#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
SPEC="$ROOT_DIR/contracts/openapi/openapi.yaml"

echo "OpenAPI spec: $SPEC"
echo "骨架阶段：此处预留 codegen（如 oapi-codegen / openapi-typescript）。"
echo "生成目标建议："
echo "  - Go:  services/api/internal/api (或 gen/)"
echo "  - TS:  packages/api-client/src"
echo "请在接入具体生成器后完善本脚本。"
