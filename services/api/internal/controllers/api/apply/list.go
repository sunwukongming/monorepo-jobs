package apply

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"

	"github.com/gin-gonic/gin"
)

func ListAction(c *gin.Context) {
	services.Handle(c, func() (interface{}, error) {
		var request ListRequest
		var page *services.Page
		if err := c.ShouldBindJSON(&request); err != nil {
			return nil, err
		}
		if request.DestCity == "全国" {
			request.DestCity = ""
		}
		var applies []bolejiang.AccountApply
		keyword := "%" + request.Keyword + "%"
		destCity := "%" + request.DestCity + "%"
		destIndustry := "%" + request.DestIndustry + "%"
		destPosition := "%" + request.DestPosition + "%"
		session := db.Default().Model(&bolejiang.AccountApply{}).Where("(dest_city like ? and dest_industry like ? and dest_position like ?) and (description like ? or education like ? or university like ? or dest_company like ? or dest_position like ? or dest_position_tag like ?)", destCity, destIndustry, destPosition, keyword, keyword, keyword, keyword, keyword, keyword)
		if request.IsHelpRewardVisible != "" {
			if request.IsHelpRewardVisible == "0" {
				session = session.Where("help_reward = 0")
			} else {
				session = session.Where("help_reward > 0")
			}
		}
		session = session.Where("is_public = 1 and status = 0")
		session = applyDestFilters(session, request)

		page = services.NewPage(request.Page, request.PageSize)
		err := page.Execute(session.Order("updated_time desc"), &applies)
		if err != nil {
			return nil, err
		}
		return page, nil
	})
}
