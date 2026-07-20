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

func LikeAction(c *gin.Context) {
	type Request struct {
		ID uint32 `json:"id"`
	}
	services.Handle(c, func() (interface{}, error) {
		var request Request
		var wanted *model.ProfService
		if err := c.ShouldBindJSON(&request); err != nil {
			return nil, err
		}
		var err error
		wanted, err = query.ProfService.Where(query.ProfService.ID.Eq(request.ID)).First()
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("数据不存在")
			}
			return nil, err
		}
		accountId, _ := strconv.Atoi(services.AuthGetAccountID(c))
		likes, err := query.ArticleLike.Where(query.ArticleLike.ArticleTable.Eq("prof_service")).
			Where(query.ArticleLike.ArticleID.Eq(request.ID)).
			Where(query.ArticleLike.AccountID.Eq(uint32(accountId))).Find()
		if err != nil {
			return nil, err
		}
		if len(likes) == 0 {
			var like model.ArticleLike
			like.ArticleTable = "prof_service"
			like.ArticleID = request.ID
			like.AccountID = uint32(accountId)
			err := query.ArticleLike.Create(&like)
			if err != nil {
				return nil, err
			}
			tx := mysql.Gorm().Exec("update prof_service as a set like_count = (select count(*) from article_like where article_table = ? and article_id = ?) where id = ?", "prof_service", request.ID, request.ID)
			if tx.Error != nil {
				return nil, tx.Error
			}
		}
		return wanted, nil
	})
}
