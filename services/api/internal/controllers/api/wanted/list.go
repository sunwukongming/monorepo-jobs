package wanted

import (
	"app/internal/db/mysql/query"
	"app/internal/services"

	"github.com/gin-gonic/gin"
)

func ListAction(c *gin.Context) {
	type Request struct {
		Keyword  string `json:"keyword"`
		Page     int    `json:"page"`
		PageSize int    `json:"pageSize"`
	}
	var page *services.Page
	err := func() error {
		var request Request
		if err := c.ShouldBindJSON(&request); err != nil {
			return err
		}
		dao := query.ProfService.Where(query.ProfService.Status.Eq(1)).Order(query.ProfService.TimeUpdate.Desc(), query.ProfService.ID.Desc())
		//session := db.Default().Table(bolejiang.ProfService{}).Where("status = 1").OrderBy("time_update desc, id desc")
		keyword := "%" + request.Keyword + "%"
		if request.Keyword != "" {
			dao = dao.Where(query.ProfService.Name.Like(keyword)).Where(query.ProfService.Introduction.Like(keyword))
			//session.Where("name like ? or introduction like ?", keyword, keyword)
		}
		count, err := dao.Count()
		if err != nil {
			return err
		}
		page = services.NewPage(request.Page, request.PageSize)
		//var wanteds []bolejiang.ProfService
		wanteds, err := dao.Offset(page.Offset).Limit(page.PerPage).Find()
		if err != nil {
			return err
		}
		page.List = wanteds
		page.Total = int(count)
		return nil
	}()
	if err != nil {
		services.ResponseError(c, -1, err.Error(), nil)
		return
	}
	services.ResponseSuccess(c, page)
}
