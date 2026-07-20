package router_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"app/internal/testutil"

	"github.com/DATA-DOG/go-sqlmock"
)

// assertEnvelope 校验统一响应外壳（code/message/data/timestamp）—— 这是 handler 重构
// （如 ③ handler 包装器）必须保持不变的契约。返回解析后的 body。
func assertEnvelope(t *testing.T, w *httptest.ResponseRecorder, wantCode float64) map[string]interface{} {
	t.Helper()
	if w.Code != http.StatusOK {
		t.Fatalf("HTTP status = %d, want 200; body=%s", w.Code, w.Body.String())
	}
	var body map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatalf("响应非合法 JSON: %v; body=%s", err, w.Body.String())
	}
	for _, k := range []string{"code", "message", "data", "timestamp"} {
		if _, ok := body[k]; !ok {
			t.Errorf("响应缺少字段 %q: %s", k, w.Body.String())
		}
	}
	if code, _ := body["code"].(float64); code != wantCode {
		t.Errorf("code = %v, want %v; body=%s", body["code"], wantCode, w.Body.String())
	}
	return body
}

// dictionary/data 纯内存、无 DB，验证路由装配 + 中间件 + 响应外壳。
func TestSmoke_DictionaryData(t *testing.T) {
	r, _ := testutil.Setup(t)
	req := httptest.NewRequest(http.MethodGet, "/api/dictionary/data", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assertEnvelope(t, w, 0)
}

// banner/list 走一次 DB 查询，用 sqlmock 返回空结果集，验证 DB handler 的成功外壳。
func TestSmoke_BannerList(t *testing.T) {
	r, mock := testutil.Setup(t)
	mock.ExpectQuery("banner").WillReturnRows(sqlmock.NewRows([]string{"id", "title"}))

	req := httptest.NewRequest(http.MethodGet, "/api/banner/list", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assertEnvelope(t, w, 0)
}
