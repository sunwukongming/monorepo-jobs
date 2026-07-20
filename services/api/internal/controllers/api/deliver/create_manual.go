package deliver

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"
	"app/pkg/utils"
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateManualAction(c *gin.Context) {
	type Request struct {
		PassageId          int    `json:"passageId"`
		PassageRecommendId int    `json:"passageRecommendId"`
		Name               string `json:"Name"`
		Mobile             string `json:"Mobile"`
		Email              string `json:"Email"`
		RecommendComment   string `json:"recommendComment"`
		ResumeUrl          string `json:"resumeUrl"`
	}
	services.Handle(c, func() (interface{}, error) {
		var request Request
		var deliver bolejiang.Deliver
		if err := c.ShouldBindJSON(&request); err != nil {
			return nil, err
		}
		if strings.Trim(request.Name, " ") == "" {
			return nil, errors.New("候选人姓名不可为空")
		}
		if !utils.ValidateIsMobile(request.Mobile) {
			return nil, errors.New("候选人手机号格式不正确")
		}
		if strings.Trim(request.Email, " ") != "" && !utils.ValidateIsEmail(request.Email) {
			return nil, errors.New("候选人邮箱格式不正确")
		}
		if strings.Trim(request.RecommendComment, " ") == "" {
			return nil, errors.New("推荐评语不可为空")
		}
		if strings.Trim(request.ResumeUrl, " ") == "" {
			return nil, errors.New("请上传候选人简历")
		}

		//当前账号
		accountId := services.AuthGetAccountID(c)
		if accountId == "" {
			return nil, errors.New("用户不存在")
		}
		// Common 中间件已校验并写入 context，直接复用避免重复查询
		accountPtr, err := services.AuthGetAccountOrError(c)
		if err != nil {
			return nil, err
		}
		currentAccount := *accountPtr
		if currentAccount.Mobile == request.Mobile {
			return nil, errors.New("无法推荐自己")
		}

		var passage bolejiang.Passage
		ok, err := db.Get(db.Default().Where("id = ?", request.PassageId), &passage)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, errors.New("职位数据不存在")
		}

		// 查询候选人是否应聘该职位
		deliver = bolejiang.Deliver{}
		ok, err = db.Get(db.Default().Where("passage_id = ? and mobile = ? and is_real = 1", request.PassageId, request.Mobile), &deliver)
		if err != nil {
			return nil, err
		}
		if ok {
			return nil, errors.New("人选已被推荐或已应聘该职位，暂不能推荐")
		}

		deliver.PassageId = passage.Id
		deliver.AccountId = 0
		deliver.Name = request.Name
		deliver.Mobile = request.Mobile
		deliver.Email = request.Email
		deliver.ResumeUrl = request.ResumeUrl
		deliver.RecommendAccountId = currentAccount.Id
		deliver.RecommendComment = request.RecommendComment
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
		} else {
			deliver.PassageRecommendPath = ""
		}
		deliver.Type = 3
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
		if deliver.PassageRecommendId != 0 {
			//services.CountUpdatePassageRecommend(deliver.PassageRecommendId)
			services.CountUpdatePassageRecommendByPath(deliver.GetPassageRecommendFullPath())
			services.CountUpdateAccount(deliver.RecommendAccountId)
		}
		services.CountUpdateAccount(deliver.AccountId)
		return deliver, nil
	})
}
