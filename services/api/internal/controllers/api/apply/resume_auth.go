package apply

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"
	"app/pkg/utils"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

func ResumeAuthAction(c *gin.Context) {
	type Request struct {
		ID int `json:"id"`
	}
	var request Request
	err := func() error {
		if err := c.ShouldBindJSON(&request); err != nil {
			return err
		}
		accountId := services.AuthGetAccountID(c)
		if accountId == "" {
			return errors.New("用户不存在")
		}

		var currentAccount bolejiang.Account
		currentAccountOk, err := db.Default().Where("id = ?", accountId).Get(&currentAccount)
		if err != nil {
			return err
		}

		if currentAccountOk {
			if currentAccount.IsResumeWatcher == 1 {
				return nil
			}
		}

		var accountApply bolejiang.AccountApply
		ok, err := db.Default().Where("id = ?", request.ID).Get(&accountApply)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("职位不存在")
		}
		var accountApplyResumeAuth bolejiang.AccountApplyResumeAuth
		ok, err = db.Default().Where("request_account_id = ? and account_apply_id = ?", accountId, request.ID).Get(&accountApplyResumeAuth)
		if err != nil {
			return err
		}
		if ok {
			if accountApplyResumeAuth.AuthState == 0 {
				return errors.New("您已经请求过查看候选人简历，请耐心等待审核")
			}
			if accountApplyResumeAuth.AuthState == 1 {
				return errors.New("您的请求已通过")
			}
			return errors.New("您的请求被拒绝，请联系管理员")
		}
		accountApplyResumeAuth.AccountApplyId = accountApply.Id
		accountApplyResumeAuth.RequestAccountId = utils.IntVal(accountId)
		accountApplyResumeAuth.CreatedTime = time.Now().Unix()
		accountApplyResumeAuth.UpdatedTime = time.Now().Unix()
		_, err = db.Default().Insert(&accountApplyResumeAuth)
		if err != nil {
			return err
		}
		return nil
	}()
	if err != nil {
		services.ResponseError(c, -1, err.Error(), nil)
		return
	}
	services.ResponseSuccess(c, nil)
}
