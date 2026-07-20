package article

import (
	"app/db"
	"app/internal/services"

	"github.com/gin-gonic/gin"
)

func GetArticlesAction(model interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		type Request struct {
			Keyword  string `json:"keyword"`
			Page     int    `json:"page"`
			PageSize int    `json:"pageSize"`
		}
		services.Handle(c, func() (interface{}, error) {
			var page *services.Page
			var request Request
			if err := c.ShouldBindJSON(&request); err != nil {
				return nil, err
			}
			session := db.Default().Table(db.TableName(model)).Where("status = 0").Order("is_top desc, time_update desc, id desc")
			keyword := "%" + request.Keyword + "%"
			if request.Keyword != "" {
				session = session.Where("title like ? or content like ?", keyword, keyword)
			}
			page = services.NewPage(request.Page, request.PageSize)
			var models []map[string]interface{}
			err := page.Execute(session, &models)
			if err != nil {
				return nil, err
			}
			list := []map[string]any{}
			for i := range models {
				model := map[string]any{}
				for k, v := range models[i] {
					model[k] = v
				}
				list = append(list, model)
			}
			page.List = list
			return page, nil
		})
	}
}
