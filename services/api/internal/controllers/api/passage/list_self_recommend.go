package passage

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"
	"app/pkg/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func ListSelfRecommendAction(c *gin.Context) {
	type Request struct {
		Keyword  string `json:"keyword"`
		Page     int    `json:"page"`
		PageSize int    `json:"pageSize"`
	}

	type DeliverPassage struct {
		Deliver bolejiang.Deliver `xorm:"extends"`
		Passage bolejiang.Passage `xorm:"extends"`
		Account bolejiang.Account `xorm:"extends"`
	}

	var page *services.Page
	var request Request
	err := func() error {
		if err := c.ShouldBindJSON(&request); err != nil {
			return err
		}
		accountId := services.AuthGetAccountID(c)
		fmt.Println("accountID:", accountId)
		var deliverPassages []DeliverPassage
		page = services.NewPage(request.Page, request.PageSize)
		query := db.Default().Table(bolejiang.Deliver{}).
			Join("LEFT", "passage", "deliver.passage_id = passage.id").
			Where("deliver.account_id = ? and deliver.is_real > 0", accountId).
			OrderBy("deliver.deliver_time desc, deliver.created_time desc")
		if request.Keyword != "" {
			query.Where("(passage.title like ? or passage.edit_content like ? or account.name like ? or account.mobile like ?)",
				"%"+request.Keyword+"%", "%"+request.Keyword+"%", "%"+request.Keyword+"%", "%"+request.Keyword+"%")
		}
		err := page.Execute(query, &deliverPassages)
		if err != nil {
			return err
		}
		comapnyAddressIds := make([]int, 0)
		passageIds := make([]int, 0)
		//passageRecommendIds := make([]int, 0)
		for _, deliverPassage := range deliverPassages {
			passageIds = append(passageIds, deliverPassage.Passage.Id)
			//passageRecommendIds = append(passageRecommendIds, passage.PassageRecommendId)
			comapnyAddressIds = append(comapnyAddressIds, deliverPassage.Passage.CompanyAddressId)
		}
		passageCompanies := make([]bolejiang.PassageCompany, 0)
		err = db.Default().In("address_id", comapnyAddressIds).Find(&passageCompanies)
		if err != nil {
			return err
		}
		var delivers []bolejiang.Deliver
		err = db.Default().Where("recommend_account_id = ?", accountId).In("passage_id", passageIds).OrderBy("updated_time desc, id desc").Find(&delivers)
		if err != nil {
			return err
		}

		passageIDs := make([]uint32, 0, len(deliverPassages))
		for _, deliverPassage := range deliverPassages {
			passageIDs = append(passageIDs, uint32(deliverPassage.Passage.Id))
		}
		passageFulls, err := services.PassageListFullByIDs(passageIDs, uint32(utils.IntVal(accountId)))
		if err != nil {
			return err
		}

		rows := []interface{}{}
		for _, passage := range deliverPassages {
			row := map[string]any{}
			for _, passageFull := range passageFulls {
				if passageFull.Passage.ID == uint32(passage.Deliver.PassageId) {
					row["passage"] = passageFull
				}
			}
			row["deliver"] = passage.Deliver
			rows = append(rows, row)
			// item := services.PassageResponse{
			// 	Passage:         passage.Passage,
			// 	CityName:        data.CityMap[passage.Passage.CityId].Name,
			// 	DistrictName:    data.DistrictMap[passage.Passage.DistrictId].Name,
			// 	IndustryName:    data.IndustryMap[passage.Passage.IndustryPath].Name,
			// 	PositionTagName: data.PositionTagMap[passage.Passage.PositionTagPath].Name,
			// }
			// for _, passageCompany := range passageCompanies {
			// 	if item.PsgCompany == passageCompany.Id {
			// 		item.Address = passageCompany.Address
			// 		item.OutName = passageCompany.OutName
			// 		item.CompanyRemark = passageCompany.Remark
			// 	}
			// }
			// passage.Deliver.AccountName = passage.Account.Name
			// passage.Deliver.AccountMobile = passage.Account.Mobile
			// if passage.Deliver.AccountName == "" {
			// 	//passage.Deliver.AccountName = passage.Deliver.AccountMobile
			// }
			// rows = append(rows, map[string]interface{}{
			// 	"deliver": passage.Deliver,
			// 	"passage": item,
			// })
		}
		page.List = rows
		return nil
	}()
	if err != nil {
		services.ResponseError(c, -1, err.Error(), nil)
		return
	}
	services.ResponseSuccess(c, page)
}
