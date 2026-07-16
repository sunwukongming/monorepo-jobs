package article

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"
	"errors"

	"github.com/gin-gonic/gin"
)

func GetArticleLikeAction(model interface{}) gin.HandlerFunc {
	tableName := db.Default().TableName(model)
	return func(c *gin.Context) {
		type Request struct {
			ID int `json:"id"`
		}
		var instance = map[string]interface{}{}
		err := func() error {
			var request Request
			if err := c.ShouldBindJSON(&request); err != nil {
				return err
			}
			ok, err := db.Default().Table(model).ID(request.ID).Get(&instance)
			if err != nil {
				return err
			}
			if !ok {
				return errors.New("文章数据不存在")
			}
			accountId := services.AuthGetAccountID(c)
			var account bolejiang.Account
			ok, err = db.Default().Where("id = ?", accountId).Get(&account)
			if err != nil {
				return err
			}
			if !ok {
				return errors.New("账号不存在")
			}

			var articleLike bolejiang.ArticleLike
			ok, err = db.Default().Where("article_table = ? and article_id = ? and account_id = ?", tableName, request.ID, account.Id).Get(&articleLike)
			if err != nil {
				return err
			}
			if ok {
				return errors.New("你已经点过赞了")
			} else {
				articleLike.ArticleTable = tableName
				articleLike.ArticleId = request.ID
				articleLike.AccountId = account.Id
				_, err = db.Default().Insert(&articleLike)
				if err != nil {
					return err
				}
				_, err = db.Default().Exec("update "+tableName+" as a set like_count = (select count(*) from article_like where article_table = ? and article_id = ?) where id = ?", tableName, request.ID, request.ID)
				if err != nil {
					return err
				}

			}
			return nil
		}()
		if err != nil {
			services.ResponseError(c, -1, err.Error(), nil)
			return
		}
		services.ResponseSuccess(c, nil)
	}
}
