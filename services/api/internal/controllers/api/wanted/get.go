package wanted

import (
	"app/db"
	"app/internal/db/mysql/model"
	"app/internal/db/mysql/query"
	"app/internal/services"
	"app/models/bolejiang"
	"app/pkg/utils"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAction(c *gin.Context) {
	type Request struct {
		ID uint32 `json:"id"`
	}
	type Response struct {
		model.ProfService
		Passages any `json:"passages"`
		IsLike   any `json:"isLike"`
	}
	var response Response
	var wanted *model.ProfService
	err := func() error {
		var request Request
		if err := c.ShouldBindJSON(&request); err != nil {
			return err
		}
		tableName := "prof_service"
		var err error
		wanted, err = query.ProfService.Where(query.ProfService.ID.Eq(request.ID)).First()
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("数据不存在")
			}
			return err
		}
		response.ProfService = *wanted
		passageIds := wanted.PassageIds
		if passageIds != "" {
			passages, err := services.PassageListFullByIDs(
				utils.Uint32ArrayFromStringArray(strings.Split(passageIds, ",")),
				uint32(utils.IntVal(services.AuthGetAccountID(c))),
			)
			if err != nil {
				return err
			}
			response.Passages = passages
		}
		accountId := services.AuthGetAccountID(c)
		var articleLike bolejiang.ArticleLike
		ok, err := db.Default().Where("article_table = ? and article_id = ? and account_id = ?", tableName, request.ID, accountId).Get(&articleLike)
		if err != nil {
			return err
		}
		response.IsLike = ok
		_, err = query.ProfService.Where(query.ProfService.ID.Eq(request.ID)).UpdateSimple(query.ProfService.ViewCount.Add(1))
		if err != nil {
			return err
		}
		return nil
	}()
	if err != nil {
		services.ResponseError(c, -1, err.Error(), nil)
		return
	}
	services.ResponseSuccess(c, response)
}
