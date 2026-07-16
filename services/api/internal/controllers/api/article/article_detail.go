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
	tableName := db.Default().TableName(model)
	return func(c *gin.Context) {
		type Request struct {
			ID int `json:"id"`
		}
		var instance = map[string]string{}
		data := map[string]any{}
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
				return errors.New("数据不存在")
			}
			for k, v := range instance {
				data[k] = v
			}
			passageIds := instance["passage_ids"]
			if passageIds != "" {
				passages, err := services.PassageListFullByIDs(
					utils.Uint32ArrayFromStringArray(strings.Split(passageIds, ",")),
					uint32(utils.IntVal(services.AuthGetAccountID(c))),
				)
				if err != nil {
					return err
				}
				data["passages"] = passages
			}
			accountId := services.AuthGetAccountID(c)
			var articleLike bolejiang.ArticleLike
			ok, err = db.Default().Where("article_table = ? and article_id = ? and account_id = ?", tableName, request.ID, accountId).Get(&articleLike)
			if err != nil {
				return err
			}
			data["isLike"] = ok
			_, err = db.Default().Exec("update "+tableName+" set view_count = view_count + 1 where id = ?", request.ID)
			if err != nil {
				return err
			}
			return nil
		}()
		if err != nil {
			services.ResponseError(c, -1, err.Error(), nil)
			return
		}
		services.ResponseSuccess(c, data)
	}
}
