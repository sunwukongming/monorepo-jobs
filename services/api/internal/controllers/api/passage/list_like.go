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
	var page *services.Page
	var request ListLikeRequest
	err := func() error {
		if err := c.ShouldBindJSON(&request); err != nil {
			return err
		}
		accountId := services.AuthGetAccountID(c)
		var passages []bolejiang.Passage
		page = services.NewPage(request.Page, request.PageSize)
		query := db.Default().Table(new(bolejiang.Passage)).Select("passage.*").Where("account_id = ?", accountId).
			Join("LEFT", "passage_like", "passage.id = passage_like.passage_id").
			OrderBy("passage_like.id desc")

		if request.SimilarPassageId != 0 {
			var similarPassage bolejiang.Passage
			ok, err := db.Default().Where("id = ?", similarPassage).Get(&similarPassage)
			if err != nil {
				return err
			}
			if ok {
				query.Where("passage.industry_path like ?", similarPassage.IndustryPath+"%")
				query.Where("passage.position_tag_path like ?", similarPassage.PositionTagPath+"%")
			}
		}
		if request.Keyword != "" {
			query.Where("(passage.title like ? or passage.edit_content like ?)", "%"+request.Keyword+"%", "%"+request.Keyword+"%")
		}
		if request.CityId != 0 {
			query.Where("passage.city_id = ?", request.CityId)
		}
		if request.DistrictId != 0 {
			query.Where("passage.district_id = ?", request.DistrictId)
		}
		if request.IndustryPath != "" {
			query.Where("(passage.industry_path like ? or passage.industry_path = ?)", request.IndustryPath+"-%", request.IndustryPath)
		}
		if request.PositionTagPath != "" {
			query.Where("(passage.position_tag_path like ? or passage.position_tag_path = ?)", request.PositionTagPath+"-%", request.PositionTagPath)
		}

		err := page.Execute(query, &passages)
		if err != nil {
			return err
		}
		passageIDs := make([]uint32, 0, len(passages))
		for _, passage := range passages {
			passageIDs = append(passageIDs, uint32(passage.Id))
		}
		passageFulls, err := services.PassageListFullByIDs(passageIDs, uint32(utils.IntVal(accountId)))
		if err != nil {
			return err
		}

		// comapnyAddressIds := make([]int, 0)
		// for _, passage := range passages {
		// 	comapnyAddressIds = append(comapnyAddressIds, passage.CompanyAddressId)
		// }
		// passageCompanies := make([]bolejiang.PassageCompany, 0)
		// err = db.Default().In("address_id", comapnyAddressIds).Find(&passageCompanies)
		// if err != nil {
		// 	return err
		// }

		// rows := []interface{}{}
		// for _, passage := range passages {
		// 	item := services.PassageResponse{
		// 		Passage:         passage,
		// 		CityName:        data.CityMap[passage.CityId].Name,
		// 		DistrictName:    data.DistrictMap[passage.DistrictId].Name,
		// 		IndustryName:    data.IndustryMap[passage.IndustryPath].Name,
		// 		PositionTagName: data.PositionTagMap[passage.PositionTagPath].Name,
		// 	}
		// 	for _, passageCompany := range passageCompanies {
		// 		if item.CompanyAddressId == passageCompany.AddressId {
		// 			item.Address = passageCompany.Address
		// 			item.OutName = passageCompany.OutName
		// 		}
		// 	}
		// 	rows = append(rows, item)
		// }
		page.List = passageFulls
		return nil
	}()
	if err != nil {
		services.ResponseError(c, -1, err.Error(), nil)
		return
	}
	services.ResponseSuccess(c, page)
}
