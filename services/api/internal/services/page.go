package services

import (
	"gorm.io/gorm"
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

// Execute 执行分页查询。
// tx 为已构建好条件（Where/Joins/Order/Select 等）的 GORM 查询，
// 先统计总数，再按 offset/limit 取出当前页数据写入 rowSlicePtr。
func (page *Page) Execute(tx *gorm.DB, rowSlicePtr interface{}) error {
	var total int64
	// Count 会忽略 Order/Limit/Offset，仅在同一组条件上统计总数
	if err := tx.Session(&gorm.Session{}).Count(&total).Error; err != nil {
		return err
	}
	page.SetTotal(int(total))
	err := tx.Limit(page.PerPage).Offset(page.Offset).Find(rowSlicePtr).Error
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
