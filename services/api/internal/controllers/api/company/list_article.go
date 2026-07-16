package company

import (
	"app/internal/db/mysql/query"
	"app/internal/services"

	"github.com/gin-gonic/gin"
)

type ListArticleRequest struct {
	services.ListRequest
	ID uint32 `json:"id"`
}

func ListArticleAction(c *gin.Context) {
	var request ListArticleRequest
	var page *services.Page
	err := func() error {
		if err := c.ShouldBindJSON(&request); err != nil {
			return err
		}
		dao := query.Cooperation.Where(query.Cooperation.PassageCompanyID.Eq(request.ID))
		page = services.NewPage(request.Page, request.PageSize)
		total, err := dao.Count()
		if err != nil {
			return err
		}
		page.SetTotal(int(total))
		articles, err := dao.Offset(page.Offset).Limit(page.PerPage).Find()
		if err != nil {
			return err
		}
		page.List = articles
		return nil
	}()
	if err != nil {
		services.ResponseError(c, -1, err.Error(), nil)
		return
	}
	services.ResponseSuccess(c, page)
}
