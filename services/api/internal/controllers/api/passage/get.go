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
	var request GetRequest
	var response Response
	//var passage bolejiang.Passage
	//var passageCompanyFull services.PassageCompanyFull
	//var passageCompany bolejiang.PassageCompany
	var passageRecommend bolejiang.PassageRecommend
	//isLike := false
	err := func() error {
		if err := c.ShouldBindJSON(&request); err != nil {
			return err
		}
		if request.ID == 0 {
			return errors.New("职位id必须填写")
		}
		if request.PassageRecommendId != "" {
			ok, err := db.Default().Where("id = ?", request.PassageRecommendId).Get(&passageRecommend)
			if err != nil {
				return err
			}
			if !ok {
				return errors.New("该分享不存在")
			}
			if passageRecommend.PassageId != int(request.ID) {
				return errors.New("分享职位和当前职位不一致")
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
				return errors.New("文章类型错误")
			}
			tablename := utils.SnakeCase(request.ArticleType)
			var article = map[string]string{}
			ok, err := db.Default().Table(tablename).Where("id = ?", request.ArticleId).Get(&article)
			if err != nil {
				return err
			}
			if !ok {
				return errors.New("文章不存在, 请确认文章id传输正确")
			}
			recommendAccountId := utils.IntVal(article["account_id"])
			if recommendAccountId != 0 {
				ok, err := db.Default().Where("account_id = ? and passage_id = ?", recommendAccountId, request.ID).Get(&passageRecommend)
				if err != nil {
					return err
				}
				if !ok {
					passageRecommend.AccountId = recommendAccountId
					passageRecommend.ParentPassageRecommendId = 0
					passageRecommend.PassageId = int(request.ID)
					passageRecommend.CreatedTime = time.Now().Unix()
					passageRecommend.UpdatedTime = time.Now().Unix()
					_, err = db.Default().Insert(&passageRecommend)
					if err != nil {
						return err
					}
				}
			}
		}

		accountId := services.AuthGetAccountID(c)
		passageResponse, err := services.PassageGetFullByID(request.ID, uint32(utils.IntVal(accountId)))
		if err != nil {
			return err
		}
		response.PassageFull = passageResponse
		passage := passageResponse.Passage

		// ok, err := db.Default().Where("id = ?", request.ID).Get(&passage)
		// if err != nil {
		// 	return err
		// }
		// if !ok {
		// 	return errors.New("职位数据不存在")
		// }

		// isLike, err = db.Default().Table(new(bolejiang.PassageLike)).Where("account_id = ? and passage_id = ?", accountId, passage.Id).Exist()
		// if err != nil {
		// 	return err
		// }
		// _, err = db.Default().Where("id = ?", passage.PsgCompany).Get(&passageCompany)
		// if err != nil {
		// 	return err
		// }

		// passageCompany, err := query.PassageCompany.Where(query.PassageCompany.ID).First()
		// if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		// 	return err
		// }
		// if err == nil {
		// 	full, err := services.CompanyFullGet(passageCompany)
		// 	if err != nil {
		// 		return err
		// 	}
		// 	passageCompanyFull = *full
		// }

		//获取当前用户
		var currentAccount bolejiang.Account
		ok, err := db.Default().Where("id = ?", accountId).Get(&currentAccount)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("账号不存在")
		}
		//deliver已经存在则读取deliver中的account, 若不存在，则读取传过来的推荐的account
		var deliver bolejiang.Deliver
		session := db.Default().Where("passage_id = ?", passage.ID)
		if utils.StringTrim(currentAccount.Mobile) == "" {
			session.Where("account_id = ?", currentAccount.Id)
		} else {
			session.Where("mobile = ?", currentAccount.Mobile)
		}
		ok, err = session.Get(&deliver)
		if err != nil {
			return err
		}
		if ok {
			var recommendAccount bolejiang.Account
			if deliver.RecommendAccountId != 0 {
				_, err = db.Default().Where("id = ?", deliver.RecommendAccountId).Get(&recommendAccount)
				if err != nil {
					return err
				}
			}
			if deliver.PassageRecommendId != 0 {
				passageRecommend = bolejiang.PassageRecommend{}
				ok, err := db.Default().Where("passage_id = ? and id = ?", passage.ID, deliver.PassageRecommendId).Get(&passageRecommend)
				if err != nil {
					return err
				}
				if !ok {
					return errors.New("该职位的推荐信息无效")
				}
			}
			response.RecommendAccount = gin.H{
				"id":                 recommendAccount.Id,
				"name":               recommendAccount.Name,
				"mobile":             recommendAccount.Mobile,
				"recommendCount":     passageRecommend.RecommendCount,
				"recommendCountL2":   passageRecommend.RecommendCountL2,
				"shareCount":         passageRecommend.ShareCount,
				"shareCountL2":       passageRecommend.ShareCountL2,
				"passageRecommendId": passageRecommend.Id,
			}
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
				_, err = db.Default().Where("id = ?", passageRecommend.AccountId).Get(&recommendAccount)
				if err != nil {
					return err
				}
				response.RecommendAccount = gin.H{
					"id":                 recommendAccount.Id,
					"name":               recommendAccount.Name,
					"mobile":             recommendAccount.Mobile,
					"recommendCount":     passageRecommend.RecommendCount,
					"recommendCountL2":   passageRecommend.RecommendCountL2,
					"shareCount":         passageRecommend.ShareCount,
					"shareCountL2":       passageRecommend.ShareCountL2,
					"passageRecommendId": passageRecommend.Id,
				}
				if deliver.AccountId != passageRecommend.AccountId {
					deliver.Type = 2
				}
			} else {
				deliver.PassageRecommendPath = ""
			}

			deliver.IsReal = 0
			deliver.CreatedTime = time.Now().Unix()
			deliver.UpdatedTime = time.Now().Unix()
			_, err = db.Default().Insert(&deliver)
			if err != nil {
				return err
			}
			services.CountUpdatePassageRecommendByPath(deliver.GetPassageRecommendFullPath())
		}

		// 获取自己的shareAccount
		var selfPassageRecommend bolejiang.PassageRecommend
		_, err = db.Default().Where("account_id = ? and passage_id = ?", currentAccount.Id, request.ID).Get(&selfPassageRecommend)
		if err != nil {
			return err
		}
		response.SelfAccount = gin.H{
			"id":                 currentAccount.Id,
			"name":               currentAccount.Name,
			"mobile":             currentAccount.Mobile,
			"recommendCount":     selfPassageRecommend.RecommendCount,
			"recommendCountL2":   selfPassageRecommend.RecommendCountL2,
			"shareCount":         selfPassageRecommend.ShareCount,
			"shareCountL2":       selfPassageRecommend.ShareCountL2,
			"passageRecommendId": selfPassageRecommend.Id,
		}
		return nil
	}()
	if err != nil {
		services.ResponseError(c, -1, err.Error(), nil)
		return
	}
	// passageResponse := services.PassageResponse{
	// 	Passage:              passage,
	// 	OutName:              passageCompanyFull.OutName,
	// 	Address:              passageCompanyFull.Address,
	// 	CompanyRemark:        passageCompanyFull.Remark,
	// 	CompanyPassageAmount: passageCompanyFull.PassageAmount,
	// 	CityName:             data.CityMap[passage.CityId].Name,
	// 	DistrictName:         data.DistrictMap[passage.DistrictId].Name,
	// 	IndustryName:         data.IndustryMap[passage.IndustryPath].Name,
	// 	PositionTagName:      data.PositionTagMap[passage.PositionTagPath].Name,
	// 	IsLike:               isLike,
	// }
	//response.PassageResponse = passageResponse
	services.ResponseSuccess(c, response)
}
