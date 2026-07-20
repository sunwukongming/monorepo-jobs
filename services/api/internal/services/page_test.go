package services

import "testing"

func TestNewPage(t *testing.T) {
	// 非法入参回退到默认值
	if p := NewPage(0, 0); p.CurrentPage != 1 || p.PerPage != 20 || p.Offset != 0 {
		t.Errorf("NewPage(0,0) = %+v, want page=1 perPage=20 offset=0", p)
	}
	// offset = (page-1) * pageSize
	if p := NewPage(3, 10); p.CurrentPage != 3 || p.PerPage != 10 || p.Offset != 20 {
		t.Errorf("NewPage(3,10) = %+v, want page=3 perPage=10 offset=20", p)
	}
}

func TestPageSetTotal(t *testing.T) {
	cases := []struct {
		total    int
		lastPage int
	}{
		{10, 1},
		{11, 2},
		{20, 2},
		{21, 3},
		{25, 3},
	}
	for _, tc := range cases {
		p := NewPage(1, 10)
		p.SetTotal(tc.total)
		if p.Total != tc.total || p.LastPage != tc.lastPage {
			t.Errorf("SetTotal(%d): total=%d lastPage=%d, want lastPage=%d", tc.total, p.Total, p.LastPage, tc.lastPage)
		}
	}
}
