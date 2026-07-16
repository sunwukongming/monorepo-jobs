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
	var request UnlikeRequest
	err := func() error {
		if err := c.ShouldBindJSON(&request); err != nil {
			return err
		}
		accountId := services.AuthGetAccountID(c)
		if accountId == "" {
			return errors.New("用户不存在")
		}
		var apply bolejiang.AccountApply
		ok, err := db.Default().Where("id = ?", request.ID).Get(&apply)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("职位不存在")
		}
		var applyLike bolejiang.AccountApplyLike
		ok, err = db.Default().Where("account_id = ? and account_apply_id = ?", accountId, request.ID).Get(&applyLike)
		if err != nil {
			return err
		}
		if ok {
			_, err := db.Default().Table(new(bolejiang.AccountApplyLike)).Where("id = ?", applyLike.Id).Delete()
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
