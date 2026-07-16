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
	var request DeliverRequest
	var deliver bolejiang.Deliver
	err := func() error {
		if err := c.ShouldBindJSON(&request); err != nil {
			return err
		}
		var passage bolejiang.Passage
		ok, err := db.Default().Where("id = ?", request.ID).Get(&passage)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("职位数据不存在")
		}
		accountId := services.AuthGetAccountID(c)
		if accountId == "" {
			return errors.New("用户不存在")
		}
		var account bolejiang.Account
		ok, err = db.Default().Where("id = ?", accountId).Get(&account)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("账号不存在")
		}
		if strings.Trim(account.Mobile, " ") == "" {
			return errors.New("您的账号手机号未填写，请进入个人中心填写手机号")
		}
		//查询手机号是否存在
		ok, err = db.Default().Table(bolejiang.Deliver{}).Where("passage_id = ? and mobile = ? and is_real = 1", passage.Id, account.Mobile).Exist()
		if err != nil {
			return err
		}
		if ok {
			//如果投递存在且是真投递
			if account.Id != 0 {
				return errors.New("您已经投递过该职位，请不要重新投递")
			} else {
				return errors.New("您已经被人投递过该职位，请不要重新投递")
			}
		}

		ok, err = db.Default().Where("passage_id = ? and mobile = ? and is_real = 0", passage.Id, account.Mobile).Get(&deliver)
		if err != nil {
			return err
		}
		if ok {
			deliver.IsReal = 1
			deliver.ProgressEid = 1
			deliver.Progress = "已投递"
			deliver.ResumeUrl = account.ResumeUrl
			deliver.DeliverTime = time.Now().Unix()
			deliver.UpdatedTime = time.Now().Unix()
			_, err := db.Default().Table(bolejiang.Deliver{}).ID(deliver.Id).Cols("is_real", "progress", "resume_url", "deliver_time", "updated_time").Update(deliver)
			if err != nil {
				return err
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
				ok, err := db.Default().Where("id = ?", request.PassageRecommendId).Get(&passageRecommend)
				if err != nil {
					return err
				}
				if !ok {
					return errors.New("推荐职位不存在")
				}
				if passage.Id != passageRecommend.PassageId {
					return errors.New("职位和推荐不一致")
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
			_, err = db.Default().Insert(&deliver)
			if err != nil {
				return err
			}
		}

		if deliver.PassageRecommendId != 0 {
			//services.CountUpdatePassageRecommend(deliver.PassageRecommendId)
			services.CountUpdatePassageRecommendByPath(deliver.GetPassageRecommendFullPath())
			services.CountUpdateAccount(deliver.RecommendAccountId)
		}
		services.CountUpdateAccount(deliver.AccountId)
		return nil
	}()
	if err != nil {
		services.ResponseError(c, -1, err.Error(), nil)
		return
	}
	services.ResponseSuccess(c, deliver)
}
