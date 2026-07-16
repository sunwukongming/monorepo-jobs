package help

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

func ApplyAction(c *gin.Context) {
	type Request struct {
		AccountApplyID int    `json:"accountApplyId" binding:"required"`
		Company        string `json:"company"`
		Position       string `json:"position"`
		HelpPlan       string `json:"helpPlan"`
	}
	var request Request
	data := gin.H{}
	err := func() error {
		if err := c.ShouldBindJSON(&request); err != nil {
			return err
		}
		/*
			if request.Company == "" {
				return errors.New("公司不可为空")
			}
			if request.Position == "" {
				return errors.New("职位不可为空")
			}
		*/
		if request.HelpPlan == "" {
			return errors.New("求职计划不可为空")
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

		/*
			helperAccount.Company = request.Company
			helperAccount.Position = request.Position
			_, err = db.Default().ID(helperAccount.Id).Cols("company", "position").Update(helperAccount)
			if err != nil {
				return err
			}
		*/

		var accountHelp bolejiang.AccountHelp
		ok, err = db.Default().Where("account_apply_id = ? and helper_account_id = ?", request.AccountApplyID, currentAccountId).Get(&accountHelp)
		if err != nil {
			return err
		}
		if ok {
			accountHelp.Company = request.Company
			accountHelp.Position = request.Position
			accountHelp.HelpPlan = request.HelpPlan
			accountHelp.UpdatedTime = time.Now().Unix()
			_, err = db.Default().ID(accountHelp.Id).Cols("company", "position", "help_plan", "updated_time").Update(accountHelp)
			if err != nil {
				return err
			}
		} else {
			accountHelp.AccountApplyId = request.AccountApplyID
			accountHelp.HelperAccountId = helperAccount.Id
			accountHelp.Company = request.Company
			accountHelp.Position = request.Position
			accountHelp.HelpPlan = request.HelpPlan
			accountHelp.CreatedTime = time.Now().Unix()
			accountHelp.UpdatedTime = time.Now().Unix()
			_, err := db.Default().Insert(&accountHelp)
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
	services.ResponseSuccess(c, data)
}
