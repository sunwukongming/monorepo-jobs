package passage

import (
	"app/internal/db/mysql/query"
	"app/internal/services"
	"app/pkg/utils"

	"github.com/gin-gonic/gin"
)

type DeliversRequest struct {
	services.ListRequest
	ID uint32 `json:"id"`
}

func DeliversAction(c *gin.Context) {
	var request DeliversRequest
	var page *services.Page
	err := func() error {
		if err := c.ShouldBindJSON(&request); err != nil {
			return err
		}
		page = services.NewPage(request.Page, request.PageSize)
		accountId := services.AuthGetAccountID(c)
		dao := query.Deliver.Where(
			query.Deliver.PassageID.Eq(request.ID),
			query.Deliver.RecommendAccountID.Eq(uint32(utils.IntVal(accountId))),
			query.Deliver.AccountID.Neq(uint32(utils.IntVal(accountId))),
			query.Deliver.IsReal.Eq(1),
		)
		total, err := dao.Count()
		if err != nil {
			return err
		}
		page.SetTotal(int(total))
		page.List, err = dao.Offset(page.Offset).Limit(page.PerPage).Find()
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
