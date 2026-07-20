package passage

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"
	"app/pkg/utils"

	"github.com/gin-gonic/gin"
)

type ListLikeRequest struct {
	Keyword          string `json:"keyword"`
	CityId           int    `json:"cityId"`
	DistrictId       int    `json:"districtId"`
	IndustryPath     string `json:"industryPath"`
	PositionTagPath  string `json:"positionTagPath"`
	SimilarPassageId int    `json:"similarPassageId"`
	Page             int    `json:"page"`
	PageSize         int    `json:"pageSize"`
}

func ListLikeAction(c *gin.Context) {
	services.Handle(c, func() (interface{}, error) {
		var page *services.Page
		var request ListLikeRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			return nil, err
		}
		accountId := services.AuthGetAccountID(c)
		var passages []bolejiang.Passage
		page = services.NewPage(request.Page, request.PageSize)
		query := db.Default().Table("passage").Select("passage.*").Where("account_id = ?", accountId).
			Joins("LEFT JOIN passage_like ON passage.id = passage_like.passage_id").
			Order("passage_like.id desc")

		query, err := applySimilarFilter(query, request.SimilarPassageId)
		if err != nil {
			return nil, err
		}
		if request.Keyword != "" {
			query = query.Where("(passage.title like ? or passage.edit_content like ?)", "%"+request.Keyword+"%", "%"+request.Keyword+"%")
		}
		query = applyPassageGeoFilters(query, request.CityId, request.DistrictId, request.IndustryPath, request.PositionTagPath)

		err = page.Execute(query, &passages)
		if err != nil {
			return nil, err
		}
		passageIDs := make([]uint32, 0, len(passages))
		for _, passage := range passages {
			passageIDs = append(passageIDs, uint32(passage.Id))
		}
		passageFulls, err := services.PassageListFullByIDs(passageIDs, uint32(utils.IntVal(accountId)))
		if err != nil {
			return nil, err
		}
		page.List = passageFulls
		return page, nil
	})
}
