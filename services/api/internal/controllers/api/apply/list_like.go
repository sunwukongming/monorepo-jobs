package apply

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"

	"github.com/gin-gonic/gin"
)

func ListLikeAction(c *gin.Context) {
	var request ListRequest
	var page *services.Page
	err := func() error {
		if err := c.ShouldBindJSON(&request); err != nil {
			return err
		}
		if request.DestCity == "全国" {
			request.DestCity = ""
		}
		accountId := services.AuthGetAccountID(c)
		var applies []bolejiang.AccountApply
		keyword := "%" + request.Keyword + "%"
		destCity := "%" + request.DestCity + "%"
		destIndustry := "%" + request.DestIndustry + "%"
		destPosition := "%" + request.DestPosition + "%"
		page = services.NewPage(request.Page, request.PageSize)
		session := db.Default().Table("account_apply").Select("account_apply.*").Where("account_apply_like.account_id = ?", accountId).
			Joins("INNER JOIN account_apply_like ON account_apply.id = account_apply_like.account_apply_id").
			Where("(account_apply.dest_city like ? and account_apply.dest_industry like ? and account_apply.dest_position like ?) and (account_apply.description like ? or account_apply.education like ? or account_apply.university like ? or dest_company like ? or dest_position like ? or dest_position_tag like ?)", destCity, destIndustry, destPosition, keyword, keyword, keyword, keyword, keyword, keyword)
		if request.IsHelpRewardVisible != "" {
			if request.IsHelpRewardVisible == "0" {
				session = session.Where("account_apply.help_reward = 0")
			} else {
				session = session.Where("account_apply.help_reward > 0")
			}
		}
		session = session.Where("account_apply.is_public = 1 and account_apply.status = 0")
		session = applyDestFilters(session, request)

		err := page.Execute(session.Order("account_apply.updated_time desc"), &applies)
		if err != nil {
			return err
		}
		return nil
	}()
	if err != nil {
		services.ResponseError(c, -1, err.Error(), nil)
		return
	}
	services.ResponseSuccess(c, page)
}
