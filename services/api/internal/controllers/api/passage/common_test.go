package passage

import (
	"testing"

	"app/models/bolejiang"
)

// 验证 GetAction 中多处复用的响应拼装 helper（推荐人 / 当前用户）字段映射正确。
func TestAccountShareInfo(t *testing.T) {
	account := bolejiang.Account{Id: 7, Name: "张三", Mobile: "13800000000"}
	recommend := bolejiang.PassageRecommend{
		Id:               9,
		RecommendCount:   2,
		RecommendCountL2: 3,
		ShareCount:       4,
		ShareCountL2:     5,
	}
	got := accountShareInfo(account, recommend)
	want := map[string]interface{}{
		"id":                 7,
		"name":               "张三",
		"mobile":             "13800000000",
		"recommendCount":     2,
		"recommendCountL2":   3,
		"shareCount":         4,
		"shareCountL2":       5,
		"passageRecommendId": 9,
	}
	if len(got) != len(want) {
		t.Fatalf("accountShareInfo 返回 %d 个字段, want %d: %v", len(got), len(want), got)
	}
	for k, w := range want {
		if got[k] != w {
			t.Errorf("accountShareInfo[%q] = %v, want %v", k, got[k], w)
		}
	}
}
