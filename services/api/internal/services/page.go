package services

import (
	"xorm.io/xorm"
)

// Page 分页对象
type Page struct {
	List        interface{} `json:"list"`
	Total       int         `json:"total"`
	CurrentPage int         `json:"currentPage"`
	LastPage    int         `json:"lastPage"`
	PerPage     int         `json:"perPage"`
	Offset      int         `json:"-"`
}

// Execute 执行分页查询
func (page *Page) Execute(session *xorm.Session, rowSlicePtr interface{}) error {
	total, err := session.Limit(page.PerPage, page.Offset).FindAndCount(rowSlicePtr)
	page.Total = int(total)
	page.LastPage = (page.Total-1)/page.PerPage + 1
	page.List = rowSlicePtr
	return err
}

func (page *Page) SetTotal(total int) {
	page.Total = total
	page.LastPage = (page.Total-1)/page.PerPage + 1
}

// NewPage 创建新分页
func NewPage(page, pageSize int) *Page {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}
	return &Page{
		List:        nil,
		CurrentPage: page,
		PerPage:     pageSize,
		Offset:      (page - 1) * pageSize,
	}
}
