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
	services.Handle(c, func() (interface{}, error) {
		var request Request
		var response Response
		var wanted *model.ProfService
		if err := c.ShouldBindJSON(&request); err != nil {
			return nil, err
		}
		tableName := "prof_service"
		var err error
		wanted, err = query.ProfService.Where(query.ProfService.ID.Eq(request.ID)).First()
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("数据不存在")
			}
			return nil, err
		}
		response.ProfService = *wanted
		passageIds := wanted.PassageIds
		if passageIds != "" {
			passages, err := services.PassageListFullByIDs(
				utils.Uint32ArrayFromStringArray(strings.Split(passageIds, ",")),
				uint32(utils.IntVal(services.AuthGetAccountID(c))),
			)
			if err != nil {
				return nil, err
			}
			response.Passages = passages
		}
		accountId := services.AuthGetAccountID(c)
		var articleLike bolejiang.ArticleLike
		ok, err := db.Get(db.Default().Where("article_table = ? and article_id = ? and account_id = ?", tableName, request.ID, accountId), &articleLike)
		if err != nil {
			return nil, err
		}
		response.IsLike = ok
		_, err = query.ProfService.Where(query.ProfService.ID.Eq(request.ID)).UpdateSimple(query.ProfService.ViewCount.Add(1))
		if err != nil {
			return nil, err
		}
		return response, nil
	})
}
