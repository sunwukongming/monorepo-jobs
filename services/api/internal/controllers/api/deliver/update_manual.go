package deliver

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

func UpdateManualAction(c *gin.Context) {
	type Request struct {
		ID               int    `json:"id"`
		RecommendComment string `json:"recommendComment"`
		ResumeUrl        string `json:"resumeUrl"`
	}
	var request Request
	var deliver bolejiang.Deliver
	err := func() error {
		if err := c.ShouldBindJSON(&request); err != nil {
			return err
		}
		if request.RecommendComment == "" {
			return errors.New("推荐评语不可为空")
		}
		if request.ResumeUrl == "" {
			return errors.New("请上传候选人简历")
		}

		//当前账号
		accountId := services.AuthGetAccountID(c)
		if accountId == "" {
			return errors.New("用户不存在")
		}
		var currentAccount bolejiang.Account
		ok, err := db.Default().Where("id = ?", accountId).Get(&currentAccount)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("账号不存在")
		}

		ok, err = db.Default().ID(request.ID).Where("recommend_account_id = ? and is_real = 1", currentAccount.Id).Get(&deliver)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("该手动投递不存在")
		}
		if deliver.AccountId != 0 {
			return errors.New("该投递是候选人自己投递，您无法修改")
		}
		deliver.ResumeUrl = request.ResumeUrl
		deliver.RecommendComment = request.RecommendComment
		deliver.UpdatedTime = time.Now().Unix()
		_, err = db.Default().ID(deliver.Id).Cols("resume_url", "recommend_comment", "updated_time").Update(deliver)
		if err != nil {
			return err
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
