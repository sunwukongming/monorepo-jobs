package apply

import (
	"app/data"
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"
	"app/pkg/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func ListAction(c *gin.Context) {
	type ListRequest struct {
		Keyword             string `json:"keyword"`
		DestCity            string `json:"destCity"`
		DestIndustry        string `json:"destIndustry"`
		DestPosition        string `json:"destPosition"`
		DestCityId          int    `json:"destCityId"`
		DestIndustryId      int    `json:"destIndustryId"`
		DestPositionTagId   int    `json:"destPositionTagId"`
		DestIndustryPath    string `json:"destIndustryPath"`
		DestPositionTagPath string `json:"destPositionTagPath"`
		IsHelpRewardVisible string `json:"isHelpRewardVisible"`
		Page                int    `json:"page"`
		PageSize            int    `json:"pageSize"`
	}
	var request ListRequest
	var page *services.Page
	err := func() error {
		if err := c.ShouldBindJSON(&request); err != nil {
			return err
		}
		if request.DestCity == "全国" {
			request.DestCity = ""
		}
		var applies []bolejiang.AccountApply
		keyword := "%" + request.Keyword + "%"
		destCity := "%" + request.DestCity + "%"
		destIndustry := "%" + request.DestIndustry + "%"
		destPosition := "%" + request.DestPosition + "%"
		session := db.Default().Where("(dest_city like ? and dest_industry like ? and dest_position like ?) and (description like ? or education like ? or university like ? or dest_company like ? or dest_position like ? or dest_position_tag like ?)", destCity, destIndustry, destPosition, keyword, keyword, keyword, keyword, keyword, keyword)
		if request.IsHelpRewardVisible != "" {
			if request.IsHelpRewardVisible == "0" {
				session.Where("help_reward = 0")
			} else {
				session.Where("help_reward > 0")
			}
		}
		session.Where("is_public = 1 and status = 0")
		if request.DestCityId != 0 {
			for _, v := range data.Cities {
				if v.Id == request.DestCityId {
					session.Where("dest_city like ?", "%"+v.Name+"%")
					break
				}
			}
		}
		if request.DestIndustryId != 0 {
			for _, v := range data.Industries {
				if v.Id == request.DestIndustryId {
					names := make([]interface{}, 0)
					names = append(names, "%"+v.Name+"%")
					queries := make([]string, 0)
					queries = append(queries, "dest_industry like ?")
					for _, item := range data.Industries {
						if utils.StringStartsWith(item.Path, v.Path+"-") {
							names = append(names, "%"+item.Name+"%")
							queries = append(queries, "dest_industry like ?")
						}
					}
					session.And(strings.Join(queries, " or "), names...)
					break
				}
			}
		}
		if request.DestPositionTagId != 0 {
			for _, v := range data.PositionTags {
				if request.DestPositionTagId == v.Id {
					names := make([]interface{}, 0)
					names = append(names, "%"+v.Name+"%")
					queries := make([]string, 0)
					queries = append(queries, "dest_position_tag like ?")
					for _, item := range data.PositionTags {
						if utils.StringStartsWith(item.Path, v.Path+"-") {
							names = append(names, "%"+item.Name+"%")
							queries = append(queries, "dest_position_tag like ?")
						}
					}
					session.And(strings.Join(queries, " or "), names...)
					break
				}
			}
		}
		if request.DestIndustryPath != "" {
			names := make([]interface{}, 0)
			queries := make([]string, 0)
			for _, item := range data.Industries {
				if utils.StringStartsWith(item.Path, request.DestIndustryPath+"-") || item.Path == request.DestIndustryPath {
					names = append(names, "%"+item.Name+"%")
					queries = append(queries, "dest_industry like ?")
				}
			}
			if len(names) > 0 {
				session.And(strings.Join(queries, " or "), names...)
			}
		}

		if request.DestPositionTagPath != "" {
			names := make([]interface{}, 0)
			queries := make([]string, 0)
			for _, item := range data.PositionTags {
				if utils.StringStartsWith(item.Path, request.DestPositionTagPath+"-") || item.Path == request.DestPositionTagPath {
					names = append(names, "%"+item.Name+"%")
					queries = append(queries, "dest_position_tag like ?")
				}
			}
			if len(names) > 0 {
				session.And(strings.Join(queries, " or "), names...)
			}
		}

		page = services.NewPage(request.Page, request.PageSize)
		err := page.Execute(session.OrderBy("updated_time desc"), &applies)
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
