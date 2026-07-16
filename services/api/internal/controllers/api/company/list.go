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
	var request ListRequest
	var page *services.Page
	err := func() error {
		if err := c.ShouldBindJSON(&request); err != nil {
			return err
		}
		keyword := "%" + request.Keyword + "%"
		dao := query.PassageCompany.Where(
			query.PassageCompany.Where(query.PassageCompany.Name.Like(keyword)).Or(query.PassageCompany.OutName.Like(keyword)),
		).Order(query.PassageCompany.IsTop.Desc(), query.PassageCompany.Sort.Desc(), query.PassageCompany.ID.Desc())
		if request.CityId > 0 {
			dao.Where(query.PassageCompany.CityID.Eq(request.CityId))
		}
		if request.DistrictId > 0 {
			dao.Where(query.PassageCompany.DistrictID.Eq(request.DistrictId))
		}
		if request.IndustryPath != "" {
			dao.Where(query.PassageCompany.Where(
				query.PassageCompany.IndustryPath.Like(request.IndustryPath + "-%"),
			).Or(query.PassageCompany.IndustryPath.Eq(request.IndustryPath)))
		}
		if request.Scale > 0 {
			dao.Where(query.PassageCompany.Scale.Eq(request.Scale))
		}
		if request.Stage > 0 {
			dao.Where(query.PassageCompany.Stage.Eq(request.Stage))
		}
		page = services.NewPage(request.Page, request.PageSize)
		total, err := dao.Count()
		if err != nil {
			return err
		}
		page.SetTotal(int(total))
		companies, err := dao.Offset(page.Offset).Limit(page.PerPage).Find()
		if err != nil {
			return err
		}
		companyFulls, err := services.CompanyFullFind(companies)
		if err != nil {
			return err
		}
		page.List = companyFulls
		return nil
	}()
	if err != nil {
		services.ResponseError(c, -1, err.Error(), nil)
		return
	}
	services.ResponseSuccess(c, page)
}
