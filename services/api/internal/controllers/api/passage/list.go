package passage

import (
	"app/db"
	mquery "app/internal/db/mysql/query"
	"app/internal/services"
	"app/models/bolejiang"
	"app/pkg/utils"

	"github.com/gin-gonic/gin"
)

type ListRequest struct {
	Keyword          string `json:"keyword"`
	CityId           int    `json:"cityId"`
	DistrictId       int    `json:"districtId"`
	IndustryPath     string `json:"industryPath"`
	PositionTagPath  string `json:"positionTagPath"`
	SimilarPassageId int    `json:"similarPassageId"`
	IsAnonymous      string `json:"isAnonymous"`
	CompanyID        int    `json:"companyId"`
	RootCompanyID    uint32 `json:"rootCompanyId"`
	Page             int    `json:"page"`
	PageSize         int    `json:"pageSize"`
}

func ListAction(c *gin.Context) {
	services.Handle(c, func() (interface{}, error) {
		var request ListRequest
		var page *services.Page
		if err := c.ShouldBindJSON(&request); err != nil {
			return nil, err
		}
		var passages []bolejiang.Passage
		query := db.Default().Table("passage").Select("passage.*").
			Joins("LEFT JOIN passage_company ON passage.psg_company = passage_company.id").
			Where("passage.type = 0 and passage.status = 0").Order("passage.mtime desc, passage.id desc")
		if request.Keyword != "" {
			accountId := services.AuthGetAccountID(c)
			var account bolejiang.Account
			ok, err := db.Get(db.Default().Where("id = ?", accountId), &account)
			if err != nil {
				return nil, err
			}
			if ok && account.IsAllies == 1 {

			} else {
				query = query.Where("passage_company.name != ?", request.Keyword)
			}
		}

		query, err := applySimilarFilter(query, request.SimilarPassageId)
		if err != nil {
			return nil, err
		}
		if request.Keyword != "" {
			query = query.Where("(passage.title like ? or passage.edit_content like ? or passage_company.name like ?)", "%"+request.Keyword+"%", "%"+request.Keyword+"%", "%"+request.Keyword+"%")
		}
		query = applyPassageGeoFilters(query, request.CityId, request.DistrictId, request.IndustryPath, request.PositionTagPath)
		if request.IsAnonymous != "" {
			query = query.Where("(passage.is_anonymous = ?)", request.IsAnonymous)
		}
		if request.CompanyID != 0 {
			query = query.Where("(passage.psg_company = ? and passage.is_anonymous = 0)", request.CompanyID)
		}
		if request.RootCompanyID != 0 {
			childCompanies, err := mquery.PassageCompany.Where(mquery.PassageCompany.CompanyID.Eq(request.RootCompanyID)).Find()
			if err != nil {
				return nil, err
			}
			ids := make([]uint32, 0, len(childCompanies))
			for i := range childCompanies {
				ids = append(ids, childCompanies[i].ID)
			}
			query = query.Where("passage.psg_company IN ?", ids).Where("passage.is_anonymous = 0")
		}
		if request.PageSize == 0 {
			request.PageSize = 10
		}
		page = services.NewPage(request.Page, request.PageSize)
		err = page.Execute(query, &passages)
		if err != nil {
			return nil, err
		}

		passageIDs := make([]uint32, 0, len(passages))
		for _, passage := range passages {
			passageIDs = append(passageIDs, uint32(passage.Id))
		}

		accountId := services.AuthGetAccountID(c)
		passageFulls, err := services.PassageListFullByIDs(passageIDs, uint32(utils.IntVal(accountId)))
		if err != nil {
			return nil, err
		}

		page.List = passageFulls
		return page, nil
	})
}
