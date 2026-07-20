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
	services.Handle(c, func() (interface{}, error) {
		var request Request
		if err := c.ShouldBindJSON(&request); err != nil {
			return nil, err
		}
		accountId := services.AuthGetAccountID(c)
		if accountId == "" {
			return nil, errors.New("用户不存在")
		}

		var currentAccount bolejiang.Account
		currentAccountOk, err := db.Get(db.Default().Where("id = ?", accountId), &currentAccount)
		if err != nil {
			return nil, err
		}

		if currentAccountOk {
			if currentAccount.IsResumeWatcher == 1 {
				return nil, nil
			}
		}

		var accountApply bolejiang.AccountApply
		ok, err := db.Get(db.Default().Where("id = ?", request.ID), &accountApply)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, errors.New("职位不存在")
		}
		var accountApplyResumeAuth bolejiang.AccountApplyResumeAuth
		ok, err = db.Get(db.Default().Where("request_account_id = ? and account_apply_id = ?", accountId, request.ID), &accountApplyResumeAuth)
		if err != nil {
			return nil, err
		}
		if ok {
			if accountApplyResumeAuth.AuthState == 0 {
				return nil, errors.New("您已经请求过查看候选人简历，请耐心等待审核")
			}
			if accountApplyResumeAuth.AuthState == 1 {
				return nil, errors.New("您的请求已通过")
			}
			return nil, errors.New("您的请求被拒绝，请联系管理员")
		}
		accountApplyResumeAuth.AccountApplyId = accountApply.Id
		accountApplyResumeAuth.RequestAccountId = utils.IntVal(accountId)
		accountApplyResumeAuth.CreatedTime = time.Now().Unix()
		accountApplyResumeAuth.UpdatedTime = time.Now().Unix()
		err = db.Default().Create(&accountApplyResumeAuth).Error
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
}
