package apply

import (
	"reflect"
	"testing"

	"app/data"
	"app/models/bolejiang"
)

// 验证「结点 + 其所有子级」的名称模式选取逻辑：
// 这是 list.go 与 list_like.go 合并到 applyDestFilters 后共用的核心逻辑，
// 也覆盖了「按 industryId（结点自身）」与「按 industryPath」两种入口现在等价这一前提。
func TestIndustryNamePatterns(t *testing.T) {
	orig := data.Industries
	defer func() { data.Industries = orig }()
	data.Industries = []bolejiang.DataIndustry{
		{Id: 1, Name: "互联网", Path: "1"},
		{Id: 2, Name: "电商", Path: "1-2"},
		{Id: 3, Name: "物流", Path: "1-2-3"},
		{Id: 4, Name: "金融", Path: "4"},
	}
	cases := []struct {
		path string
		want []interface{}
	}{
		{"1", []interface{}{"%互联网%", "%电商%", "%物流%"}}, // 结点 + 全部子级
		{"1-2", []interface{}{"%电商%", "%物流%"}},
		{"4", []interface{}{"%金融%"}},
		{"999", []interface{}{}}, // 无匹配
	}
	for _, tc := range cases {
		if got := industryNamePatterns(tc.path); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("industryNamePatterns(%q) = %v, want %v", tc.path, got, tc.want)
		}
	}
}

func TestPositionTagNamePatterns(t *testing.T) {
	orig := data.PositionTags
	defer func() { data.PositionTags = orig }()
	data.PositionTags = []bolejiang.DataPositionTag{
		{Id: 1, Name: "技术", Path: "1"},
		{Id: 2, Name: "后端", Path: "1-2"},
		{Id: 3, Name: "市场", Path: "3"},
	}
	got := positionTagNamePatterns("1")
	want := []interface{}{"%技术%", "%后端%"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("positionTagNamePatterns(%q) = %v, want %v", "1", got, want)
	}
}
