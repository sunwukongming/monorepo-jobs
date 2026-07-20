package passage

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func DeliverAction(c *gin.Context) {
	type DeliverRequest struct {
		ID                 int `json:"id"`
		PassageRecommendId int `json:"passageRecommendId"`
	}
	services.Handle(c, func() (interface{}, error) {
		var request DeliverRequest
		var deliver bolejiang.Deliver
		if err := c.ShouldBindJSON(&request); err != nil {
			return nil, err
		}
		var passage bolejiang.Passage
		ok, err := db.Get(db.Default().Where("id = ?", request.ID), &passage)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, errors.New("职位数据不存在")
		}
		accountId := services.AuthGetAccountID(c)
		if accountId == "" {
			return nil, errors.New("用户不存在")
		}
		// Common 中间件已校验并写入 context，直接复用避免重复查询
		accountPtr, err := services.AuthGetAccountOrError(c)
		if err != nil {
			return nil, err
		}
		account := *accountPtr
		if strings.Trim(account.Mobile, " ") == "" {
			return nil, errors.New("您的账号手机号未填写，请进入个人中心填写手机号")
		}
		//查询手机号是否存在
		var deliverCount int64
		err = db.Default().Model(&bolejiang.Deliver{}).Where("passage_id = ? and mobile = ? and is_real = 1", passage.Id, account.Mobile).Count(&deliverCount).Error
		if err != nil {
			return nil, err
		}
		ok = deliverCount > 0
		if ok {
			//如果投递存在且是真投递
			if account.Id != 0 {
				return nil, errors.New("您已经投递过该职位，请不要重新投递")
			} else {
				return nil, errors.New("您已经被人投递过该职位，请不要重新投递")
			}
		}

		ok, err = db.Get(db.Default().Where("passage_id = ? and mobile = ? and is_real = 0", passage.Id, account.Mobile), &deliver)
		if err != nil {
			return nil, err
		}
		if ok {
			deliver.IsReal = 1
			deliver.ProgressEid = 1
			deliver.Progress = "已投递"
			deliver.ResumeUrl = account.ResumeUrl
			deliver.DeliverTime = time.Now().Unix()
			deliver.UpdatedTime = time.Now().Unix()
			err := db.Default().Model(&deliver).Where("id = ?", deliver.Id).Select("is_real", "progress", "resume_url", "deliver_time", "updated_time").Updates(deliver).Error
			if err != nil {
				return nil, err
			}
		} else {
			deliver.PassageId = passage.Id
			deliver.AccountId = account.Id
			deliver.AccountName = account.Name
			deliver.AccountMobile = account.Mobile
			deliver.Name = account.Name
			deliver.Mobile = account.Mobile
			deliver.Email = account.Email
			deliver.ResumeUrl = account.ResumeUrl
			var passageRecommend bolejiang.PassageRecommend
			if request.PassageRecommendId != 0 {
				ok, err := db.Get(db.Default().Where("id = ?", request.PassageRecommendId), &passageRecommend)
				if err != nil {
					return nil, err
				}
				if !ok {
					return nil, errors.New("推荐职位不存在")
				}
				if passage.Id != passageRecommend.PassageId {
					return nil, errors.New("职位和推荐不一致")
				}
				deliver.PassageRecommendId = passageRecommend.Id
				deliver.PassageRecommendPath = passageRecommend.Path
				deliver.PassageRecommendPathFull = passageRecommend.PathFull
				deliver.RecommendAccountId = passageRecommend.AccountId
			} else {
				deliver.PassageRecommendPath = ""
			}
			deliver.Type = 1

			deliver.ProgressEid = 1
			deliver.Progress = "已投递"
			deliver.IsReal = 1
			deliver.DeliverTime = time.Now().Unix()
			deliver.CreatedTime = time.Now().Unix()
			deliver.UpdatedTime = time.Now().Unix()
			err = db.Default().Create(&deliver).Error
			if err != nil {
				return nil, err
			}
		}

		if deliver.PassageRecommendId != 0 {
			//services.CountUpdatePassageRecommend(deliver.PassageRecommendId)
			services.CountUpdatePassageRecommendByPath(deliver.GetPassageRecommendFullPath())
			services.CountUpdateAccount(deliver.RecommendAccountId)
		}
		services.CountUpdateAccount(deliver.AccountId)
		return deliver, nil
	})
}
