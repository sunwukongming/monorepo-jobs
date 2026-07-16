package wanted

import (
	"app/internal/db/mysql"
	"app/internal/db/mysql/model"
	"app/internal/db/mysql/query"
	"app/internal/services"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UnlikeAction(c *gin.Context) {
	type Request struct {
		ID uint32 `json:"id"`
	}
	var wanted *model.ProfService
	err := func() error {
		var request Request
		if err := c.ShouldBindJSON(&request); err != nil {
			return err
		}
		var err error
		wanted, err = query.ProfService.Where(query.ProfService.ID.Eq(request.ID)).First()
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("数据不存在")
			}
			return err
		}
		accountId, _ := strconv.Atoi(services.AuthGetAccountID(c))
		likes, err := query.ArticleLike.Where(query.ArticleLike.ArticleTable.Eq("prof_service")).
			Where(query.ArticleLike.ArticleID.Eq(request.ID)).
			Where(query.ArticleLike.AccountID.Eq(uint32(accountId))).Find()
		if err != nil {
			return err
		}
		if len(likes) > 0 {
			for _, like := range likes {
				query.ArticleLike.Where(query.ArticleLike.ID.Eq(like.ID)).Delete()
			}
			tx := mysql.Gorm().Exec("update prof_service as a set like_count = (select count(*) from article_like where article_table = ? and article_id = ?) where id = ?", "prof_service", request.ID, request.ID)
			if tx.Error != nil {
				return tx.Error
			}
		}
		return nil
	}()
	if err != nil {
		services.ResponseError(c, -1, err.Error(), nil)
		return
	}
	services.ResponseSuccess(c, wanted)
}
