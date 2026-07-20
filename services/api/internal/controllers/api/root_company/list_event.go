package root_company

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
	services.Handle(c, func() (interface{}, error) {
		var request ListEventRequest
		var page *services.Page
		if err := c.ShouldBindJSON(&request); err != nil {
			return nil, err
		}
		dao := query.PassageCompanyEvent.
			Where(query.PassageCompanyEvent.RootCompanyID.Eq(request.ID)).
			Order(query.PassageCompanyEvent.EventDate.Desc(), query.PassageCompanyEvent.ID.Desc())
		page = services.NewPage(request.Page, request.PageSize)
		total, err := dao.Count()
		if err != nil {
			return nil, err
		}
		page.SetTotal(int(total))
		events, err := dao.Offset(page.Offset).Limit(page.PerPage).Find()
		if err != nil {
			return nil, err
		}
		page.List = events
		return page, nil
	})
}
