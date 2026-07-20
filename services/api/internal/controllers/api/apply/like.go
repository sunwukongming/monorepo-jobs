package apply

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"
	"app/pkg/utils"
	"errors"

	"github.com/gin-gonic/gin"
)

func LikeAction(c *gin.Context) {
	type LikeRequest struct {
		ID int `json:"id"`
	}
	services.Handle(c, func() (interface{}, error) {
		var request LikeRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			return nil, err
		}
		accountId := services.AuthGetAccountID(c)
		if accountId == "" {
			return nil, errors.New("用户不存在")
		}
		var accountApply bolejiang.AccountApply
		ok, err := db.Get(db.Default().Where("id = ?", request.ID), &accountApply)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, errors.New("职位不存在")
		}
		var applyLike bolejiang.AccountApplyLike
		ok, err = db.Get(db.Default().Where("account_id = ? and account_apply_id = ?", accountId, request.ID), &applyLike)
		if err != nil {
			return nil, err
		}
		if ok {
			return nil, errors.New("该求职已被收藏")
		}
		applyLike.AccountId = utils.IntVal(accountId)
		applyLike.AccountApplyId = request.ID
		err = db.Default().Create(&applyLike).Error
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
}
