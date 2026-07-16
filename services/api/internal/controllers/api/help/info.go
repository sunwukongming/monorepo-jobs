package help

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"
	"errors"

	"github.com/gin-gonic/gin"
)

func InfoAction(c *gin.Context) {
	type Request struct {
		AccountApplyID int `json:"accountApplyId" binding:"required"`
	}
	var request Request
	data := gin.H{}
	err := func() error {
		if err := c.ShouldBindJSON(&request); err != nil {
			return err
		}
		currentAccountId := services.AuthGetAccountID(c)
		var helperAccount bolejiang.Account
		ok, err := db.Default().Where("id = ?", currentAccountId).Get(&helperAccount)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("当前用户不存在")
		}
		var accountApply bolejiang.AccountApply
		ok, err = db.Default().Where("id = ?", request.AccountApplyID).Get(&accountApply)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("求职者不存在")
		}

		var accountHelp bolejiang.AccountHelp
		_, err = db.Default().Where("account_apply_id = ? and helper_account_id = ?", request.AccountApplyID, currentAccountId).Get(&accountHelp)
		if err != nil {
			return err
		}
		data["helpPlan"] = accountHelp.HelpPlan
		data["company"] = accountHelp.Company
		data["position"] = accountHelp.Position
		data["accountApply"] = accountApply
		return nil
	}()
	if err != nil {
		services.ResponseError(c, -1, err.Error(), nil)
		return
	}
	services.ResponseSuccess(c, data)
}
