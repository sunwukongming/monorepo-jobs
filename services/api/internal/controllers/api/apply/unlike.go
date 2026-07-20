package apply

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"
	"errors"

	"github.com/gin-gonic/gin"
)

func UnlikeAction(c *gin.Context) {
	type UnlikeRequest struct {
		ID int `json:"id"`
	}
	services.Handle(c, func() (interface{}, error) {
		var request UnlikeRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			return nil, err
		}
		accountId := services.AuthGetAccountID(c)
		if accountId == "" {
			return nil, errors.New("用户不存在")
		}
		var apply bolejiang.AccountApply
		ok, err := db.Get(db.Default().Where("id = ?", request.ID), &apply)
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
			err := db.Default().Where("id = ?", applyLike.Id).Delete(new(bolejiang.AccountApplyLike)).Error
			if err != nil {
				return nil, err
			}
		}
		return nil, nil
	})
}
