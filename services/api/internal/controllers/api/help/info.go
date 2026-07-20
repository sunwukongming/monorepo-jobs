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
	services.Handle(c, func() (interface{}, error) {
		var request Request
		data := gin.H{}
		if err := c.ShouldBindJSON(&request); err != nil {
			return nil, err
		}
		currentAccountId := services.AuthGetAccountID(c)
		var helperAccount bolejiang.Account
		ok, err := db.Get(db.Default().Where("id = ?", currentAccountId), &helperAccount)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, errors.New("当前用户不存在")
		}
		var accountApply bolejiang.AccountApply
		ok, err = db.Get(db.Default().Where("id = ?", request.AccountApplyID), &accountApply)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, errors.New("求职者不存在")
		}

		var accountHelp bolejiang.AccountHelp
		_, err = db.Get(db.Default().Where("account_apply_id = ? and helper_account_id = ?", request.AccountApplyID, currentAccountId), &accountHelp)
		if err != nil {
			return nil, err
		}
		data["helpPlan"] = accountHelp.HelpPlan
		data["company"] = accountHelp.Company
		data["position"] = accountHelp.Position
		data["accountApply"] = accountApply
		return data, nil
	})
}
