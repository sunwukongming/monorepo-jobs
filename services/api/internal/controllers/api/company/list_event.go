package company

import (
	"app/internal/db/mysql/query"
	"app/internal/services"

	"github.com/gin-gonic/gin"
)

type ListEventRequest struct {
	services.ListRequest
	ID uint32 `json:"id"`
}

func ListEventAction(c *gin.Context) {
	var request ListEventRequest
	var page *services.Page
	err := func() error {
		if err := c.ShouldBindJSON(&request); err != nil {
			return err
		}
		dao := query.PassageCompanyEvent.Where(query.PassageCompanyEvent.PassageCompanyID.Eq(request.ID))
		page = services.NewPage(request.Page, request.PageSize)
		total, err := dao.Count()
		if err != nil {
			return err
		}
		page.SetTotal(int(total))
		events, err := dao.Offset(page.Offset).Limit(page.PerPage).Find()
		if err != nil {
			return err
		}
		page.List = events
		return nil
	}()
	if err != nil {
		services.ResponseError(c, -1, err.Error(), nil)
		return
	}
	services.ResponseSuccess(c, page)
}
