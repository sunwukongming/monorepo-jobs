package article

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"
	"app/pkg/utils"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetArticleDetailAction(model interface{}) gin.HandlerFunc {
	tableName := db.TableName(model)
	return func(c *gin.Context) {
		type Request struct {
			ID int `json:"id"`
		}
		services.Handle(c, func() (interface{}, error) {
			var instance = map[string]interface{}{}
			data := map[string]any{}
			var request Request
			if err := c.ShouldBindJSON(&request); err != nil {
				return nil, err
			}
			ok, err := db.Get(db.Default().Table(db.TableName(model)).Where("id = ?", request.ID), &instance)
			if err != nil {
				return nil, err
			}
			if !ok {
				return nil, errors.New("数据不存在")
			}
			for k, v := range instance {
				data[k] = v
			}
			passageIds, _ := instance["passage_ids"].(string)
			if passageIds != "" {
				passages, err := services.PassageListFullByIDs(
					utils.Uint32ArrayFromStringArray(strings.Split(passageIds, ",")),
					uint32(utils.IntVal(services.AuthGetAccountID(c))),
				)
				if err != nil {
					return nil, err
				}
				data["passages"] = passages
			}
			accountId := services.AuthGetAccountID(c)
			var articleLike bolejiang.ArticleLike
			ok, err = db.Get(db.Default().Where("article_table = ? and article_id = ? and account_id = ?", tableName, request.ID, accountId), &articleLike)
			if err != nil {
				return nil, err
			}
			data["isLike"] = ok
			err = db.Default().Exec("update "+tableName+" set view_count = view_count + 1 where id = ?", request.ID).Error
			if err != nil {
				return nil, err
			}
			return data, nil
		})
	}
}
