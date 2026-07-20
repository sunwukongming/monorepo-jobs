package apply

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"
	"errors"

	"github.com/gin-gonic/gin"
)

func DetailAction(c *gin.Context) {
	type DetailRequest struct {
		ID int `json:"id"`
	}
	type Resp struct {
		bolejiang.AccountApply
		SimilarApplies []bolejiang.AccountApply
		IsLike         bool                              `json:"isLike"`
		ResumeAuth     *bolejiang.AccountApplyResumeAuth `json:"resumeAuth"`
		ResumeUrl      string
	}
	services.Handle(c, func() (interface{}, error) {
		var request DetailRequest
		var response Resp
		if err := c.ShouldBindJSON(&request); err != nil {
			return nil, err
		}
		accountId := services.AuthGetAccountID(c)

		var currentAccount bolejiang.Account
		currentAccountOk, err := db.Get(db.Default().Where("id = ?", accountId), &currentAccount)
		if err != nil {
			return nil, err
		}

		var apply bolejiang.AccountApply
		ok, err := db.Get(db.Default().Where("id = ?", request.ID), &apply)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, errors.New("求职信息不存在")
		}
		response.AccountApply = apply
		var applies []bolejiang.AccountApply
		//destCity := "%" + apply.DestCity + "%"
		destIndustry := "%" + apply.DestIndustry + "%"
		//destPosition := "%" + apply.DestPosition + "%"
		session := db.Default().Where(" dest_industry like ? and id != ?", destIndustry, apply.Id).Order("updated_time desc")
		err = session.Find(&applies).Error
		if err != nil {
			return nil, err
		}
		response.SimilarApplies = applies
		//response.Contact = "微信 1832293829"

		if currentAccountOk {
			var like bolejiang.AccountApplyLike
			ok, err = db.Get(db.Default().Where("account_id = ? and account_apply_id = ?", accountId, apply.Id), &like)
			if err != nil {
				return nil, err
			}
			response.IsLike = ok
			if currentAccount.IsResumeWatcher == 1 {
				response.ResumeAuth = &bolejiang.AccountApplyResumeAuth{
					AuthState: 1,
				}
				var account bolejiang.Account
				_, err = db.Get(db.Default().Where("id = ?", apply.AccountId), &account)
				if err != nil {
					return nil, err
				}
				if account.ResumeUrl != "" {
					_, err := services.GetOssService().PresignUrl(account.ResumeUrl)
					if err != nil {
						return nil, err
					}
				}
				response.ResumeUrl = account.ResumeUrl
			} else {
				var accountApplyResumeAuth bolejiang.AccountApplyResumeAuth
				ok, err = db.Get(db.Default().Where("request_account_id = ? and account_apply_id = ?", accountId, request.ID), &accountApplyResumeAuth)
				if err != nil {
					return nil, err
				}
				if ok {
					response.ResumeAuth = &accountApplyResumeAuth
					if accountApplyResumeAuth.AuthState == 1 {
						var account bolejiang.Account
						_, err = db.Get(db.Default().Where("id = ?", apply.AccountId), &account)
						if err != nil {
							return nil, err
						}
						if account.ResumeUrl != "" {
							_, err := services.GetOssService().PresignUrl(account.ResumeUrl)
							if err != nil {
								return nil, err
							}
						}
						response.ResumeUrl = account.ResumeUrl
					}
				}
			}

		}
		return response, nil
	})
}
