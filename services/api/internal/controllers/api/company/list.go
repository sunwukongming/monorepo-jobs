package company

import (
	"app/internal/db/mysql/query"
	"app/internal/services"

	"github.com/gin-gonic/gin"
)

type ListRequest struct {
	Keyword      string `json:"keyword"`
	CityId       uint32 `json:"cityId"`
	DistrictId   uint32 `json:"districtId"`
	IndustryPath string `json:"industryPath"`
	Scale        uint32 `json:"scale"`
	Stage        uint32 `json:"stage"`
	Page         int    `json:"page"`
	PageSize     int    `json:"pageSize"`
}

func ListAction(c *gin.Context) {
	services.Handle(c, func() (interface{}, error) {
		var request ListRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			return nil, err
		}
		keyword := "%" + request.Keyword + "%"
		// 注意：gen DAO 的 Where 返回新对象，必须重新赋值 dao，否则条件被丢弃
		dao := query.PassageCompany.Where(
			query.PassageCompany.Where(query.PassageCompany.Name.Like(keyword)).Or(query.PassageCompany.OutName.Like(keyword)),
		).Order(query.PassageCompany.IsTop.Desc(), query.PassageCompany.Sort.Desc(), query.PassageCompany.ID.Desc())
		if request.CityId > 0 {
			dao = dao.Where(query.PassageCompany.CityID.Eq(request.CityId))
		}
		if request.DistrictId > 0 {
			dao = dao.Where(query.PassageCompany.DistrictID.Eq(request.DistrictId))
		}
		if request.IndustryPath != "" {
			dao = dao.Where(query.PassageCompany.Where(
				query.PassageCompany.IndustryPath.Like(request.IndustryPath + "-%"),
			).Or(query.PassageCompany.IndustryPath.Eq(request.IndustryPath)))
		}
		if request.Scale > 0 {
			dao = dao.Where(query.PassageCompany.Scale.Eq(request.Scale))
		}
		if request.Stage > 0 {
			dao = dao.Where(query.PassageCompany.Stage.Eq(request.Stage))
		}
		page := services.NewPage(request.Page, request.PageSize)
		total, err := dao.Count()
		if err != nil {
			return nil, err
		}
		page.SetTotal(int(total))
		companies, err := dao.Offset(page.Offset).Limit(page.PerPage).Find()
		if err != nil {
			return nil, err
		}
		companyFulls, err := services.CompanyFullFind(companies)
		if err != nil {
			return nil, err
		}
		page.List = companyFulls
		return page, nil
	})
}
