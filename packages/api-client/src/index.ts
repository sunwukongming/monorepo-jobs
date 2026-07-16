import { API_PREFIX } from "@jobs/shared";

export type HealthResponse = {
  status: string;
};

/**
 * 手写占位客户端。正式流程由 `make gen-api` 根据 OpenAPI 生成后替换本目录。
 */
export async function getHealth(baseUrl = API_PREFIX): Promise<HealthResponse> {
  const res = await fetch(`${baseUrl}/health`);
  if (!res.ok) {
    throw new Error(`health check failed: ${res.status}`);
  }
  return res.json() as Promise<HealthResponse>;
}
