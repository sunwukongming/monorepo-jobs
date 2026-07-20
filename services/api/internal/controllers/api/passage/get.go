package passage

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"
	"app/pkg/utils"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAction(c *gin.Context) {
	type GetRequest struct {
		ID                 uint32 `json:"id"`
		PassageRecommendId string `json:"passageRecommendId"`
		ArticleType        string `json:"articleType"`
		ArticleId          string `json:"articleId"`
	}
	type Response struct {
		services.PassageFull
		RecommendAccount interface{} `json:"recommendAccount"`
		SelfAccount      interface{} `json:"selfAccount"`
	}
	services.Handle(c, func() (interface{}, error) {
		var request GetRequest
		var response Response
		var passageRecommend bolejiang.PassageRecommend
		if err := c.ShouldBindJSON(&request); err != nil {
			return nil, err
		}
		if request.ID == 0 {
			return nil, errors.New("职位id必须填写")
		}
		if request.PassageRecommendId != "" {
			ok, err := db.Get(db.Default().Where("id = ?", request.PassageRecommendId), &passageRecommend)
			if err != nil {
				return nil, err
			}
			if !ok {
				return nil, errors.New("该分享不存在")
			}
			if passageRecommend.PassageId != int(request.ID) {
				return nil, errors.New("分享职位和当前职位不一致")
			}
		} else if request.ArticleType != "" {
			switch request.ArticleType {
			case "meeting":
			case "financingDemand":
			case "industryAssociation":
			case "industryInfo":
			case "investmentDemand":
			case "topicQa":
			case "cooperation":
			case "profService":
				break
			default:
				return nil, errors.New("文章类型错误")
			}
			tablename := utils.SnakeCase(request.ArticleType)
			var article = map[string]interface{}{}
			ok, err := db.Get(db.Default().Table(tablename).Where("id = ?", request.ArticleId), &article)
			if err != nil {
				return nil, err
			}
			if !ok {
				return nil, errors.New("文章不存在, 请确认文章id传输正确")
			}
			recommendAccountId := utils.IntVal(article["account_id"])
			if recommendAccountId != 0 {
				ok, err := db.Get(db.Default().Where("account_id = ? and passage_id = ?", recommendAccountId, request.ID), &passageRecommend)
				if err != nil {
					return nil, err
				}
				if !ok {
					passageRecommend.AccountId = recommendAccountId
					passageRecommend.ParentPassageRecommendId = 0
					passageRecommend.PassageId = int(request.ID)
					passageRecommend.CreatedTime = time.Now().Unix()
					passageRecommend.UpdatedTime = time.Now().Unix()
					err = db.Default().Create(&passageRecommend).Error
					if err != nil {
						return nil, err
					}
				}
			}
		}

		accountId := services.AuthGetAccountID(c)
		passageResponse, err := services.PassageGetFullByID(request.ID, uint32(utils.IntVal(accountId)))
		if err != nil {
			return nil, err
		}
		response.PassageFull = passageResponse
		passage := passageResponse.Passage

		//获取当前用户（Common 中间件已校验并写入 context，直接复用避免重复查询）
		accountPtr, err := services.AuthGetAccountOrError(c)
		if err != nil {
			return nil, err
		}
		currentAccount := *accountPtr
		//deliver已经存在则读取deliver中的account, 若不存在，则读取传过来的推荐的account
		var deliver bolejiang.Deliver
		session := db.Default().Where("passage_id = ?", passage.ID)
		if utils.StringTrim(currentAccount.Mobile) == "" {
			session = session.Where("account_id = ?", currentAccount.Id)
		} else {
			session = session.Where("mobile = ?", currentAccount.Mobile)
		}
		ok, err := db.Get(session, &deliver)
		if err != nil {
			return nil, err
		}
		if ok {
			var recommendAccount bolejiang.Account
			if deliver.RecommendAccountId != 0 {
				_, err = db.Get(db.Default().Where("id = ?", deliver.RecommendAccountId), &recommendAccount)
				if err != nil {
					return nil, err
				}
			}
			if deliver.PassageRecommendId != 0 {
				passageRecommend = bolejiang.PassageRecommend{}
				ok, err := db.Get(db.Default().Where("passage_id = ? and id = ?", passage.ID, deliver.PassageRecommendId), &passageRecommend)
				if err != nil {
					return nil, err
				}
				if !ok {
					return nil, errors.New("该职位的推荐信息无效")
				}
			}
			response.RecommendAccount = accountShareInfo(recommendAccount, passageRecommend)
		} else {
			deliver.PassageId = int(passage.ID)
			deliver.AccountId = currentAccount.Id
			deliver.AccountName = currentAccount.Name
			deliver.AccountMobile = currentAccount.Mobile
			deliver.Name = currentAccount.Name
			deliver.Mobile = currentAccount.Mobile
			deliver.Email = currentAccount.Email
			deliver.Type = 1
			if passageRecommend.Id != 0 {
				deliver.PassageRecommendId = passageRecommend.Id
				deliver.PassageRecommendPath = passageRecommend.Path
				deliver.PassageRecommendPathFull = passageRecommend.PathFull
				deliver.RecommendAccountId = passageRecommend.AccountId
				var recommendAccount bolejiang.Account
				_, err = db.Get(db.Default().Where("id = ?", passageRecommend.AccountId), &recommendAccount)
				if err != nil {
					return nil, err
				}
				response.RecommendAccount = accountShareInfo(recommendAccount, passageRecommend)
				if deliver.AccountId != passageRecommend.AccountId {
					deliver.Type = 2
				}
			} else {
				deliver.PassageRecommendPath = ""
			}

			deliver.IsReal = 0
			deliver.CreatedTime = time.Now().Unix()
			deliver.UpdatedTime = time.Now().Unix()
			err = db.Default().Create(&deliver).Error
			if err != nil {
				return nil, err
			}
			services.CountUpdatePassageRecommendByPath(deliver.GetPassageRecommendFullPath())
		}

		// 获取自己的 shareAccount
		var selfPassageRecommend bolejiang.PassageRecommend
		_, err = db.Get(db.Default().Where("account_id = ? and passage_id = ?", currentAccount.Id, request.ID), &selfPassageRecommend)
		if err != nil {
			return nil, err
		}
		response.SelfAccount = accountShareInfo(currentAccount, selfPassageRecommend)
		return response, nil
	})
}
