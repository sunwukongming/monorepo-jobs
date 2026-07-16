package company

import (
	"app/internal/db/mysql/query"
	"app/internal/services"

	"github.com/gin-gonic/gin"
)

type ListMemberRequest struct {
	services.ListRequest
	ID uint32 `json:"id"`
}

func ListMemberAction(c *gin.Context) {
	var request ListMemberRequest
	var page *services.Page
	err := func() error {
		if err := c.ShouldBindJSON(&request); err != nil {
			return err
		}
		dao := query.PassageCompanyMember.Where(query.PassageCompanyMember.PassageCompanyID.Eq(request.ID))
		page = services.NewPage(request.Page, request.PageSize)
		total, err := dao.Count()
		if err != nil {
			return err
		}
		page.SetTotal(int(total))
		members, err := dao.Offset(page.Offset).Limit(page.PerPage).Find()
		if err != nil {
			return err
		}
		page.List = members
		return nil
	}()
	if err != nil {
		services.ResponseError(c, -1, err.Error(), nil)
		return
	}
	services.ResponseSuccess(c, page)
}
