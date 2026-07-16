package article

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"

	"github.com/gin-gonic/gin"
)

func TopicQasAction(c *gin.Context) {
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
		session := db.Default().Table(bolejiang.TopicQa{}).Where("status = 0").OrderBy("is_top desc, time_update desc, id desc")
		keyword := "%" + request.Keyword + "%"
		if request.Keyword != "" {
			session.Where("title like ? or content like ?", keyword, keyword)
		}
		page = services.NewPage(request.Page, request.PageSize)
		var topicQas []bolejiang.TopicQa
		err := page.Execute(session, &topicQas)
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
