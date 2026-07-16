package root_company

import (
	"app/internal/db/mysql"
	"app/internal/db/mysql/query"
	"app/internal/services"
	"app/pkg/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type ListArticleRequest struct {
	services.ListRequest
	ID uint32 `json:"id"`
}

func ListArticleAction(c *gin.Context) {
	var request ListArticleRequest
	data := gin.H{}
	err := func() error {
		if err := c.ShouldBindJSON(&request); err != nil {
			return err
		}
		rootCompany, err := query.CompanyInfo.Where(query.CompanyInfo.ID.Eq(request.ID)).First()
		if err != nil {
			return err
		}
		list := []any{}
		items := strings.Split(rootCompany.Articles, ",")
		for i := range items {
			s := strings.Trim(items[i], " ")
			if s == "" {
				continue
			}
			a := strings.Split(s, "-")
			if len(a) == 2 {
				table := a[0]
				id := a[1]
				row := map[string]any{}
				tx := mysql.Gorm().Table(table).Where("id = ?", id).Take(&row)
				if tx.Error != nil {
					continue
				}
				row["type"] = utils.CamelCase(table)
				list = append(list, row)
			}
		}
		data["list"] = list
		return nil
	}()
	if err != nil {
		services.ResponseError(c, -1, err.Error(), nil)
		return
	}
	services.ResponseSuccess(c, data)
}
