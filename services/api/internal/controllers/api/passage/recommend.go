package passage

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"
	"app/pkg/utils"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func RecommendAction(c *gin.Context) {
	type RecommendRequest struct {
		ID                       int `json:"id"`
		ParentPassageRecommendId int `json:"parentPassageRecommendId"`
	}
	var request RecommendRequest
	var passageRecommend bolejiang.PassageRecommend
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
			return errors.New("当前账号未登录")
		}
		ok, err = db.Default().Where("account_id = ? and passage_id = ?", accountId, request.ID).Get(&passageRecommend)
		if err != nil {
			return err
		}
		if ok {
			//如果存在则不更改任何数据
			passageRecommend.UpdatedTime = time.Now().Unix()
			_, err = db.Default().Table(bolejiang.PassageRecommend{}).ID(passageRecommend.Id).Cols("updated_time").Update(passageRecommend)
			if err != nil {
				return err
			}
		} else {
			//如果不存在则创建一个推荐数据
			if request.ParentPassageRecommendId != 0 {
				var parentRecommend bolejiang.PassageRecommend
				ok, err := db.Default().ID(request.ParentPassageRecommendId).Get(&parentRecommend)
				if err != nil {
					return err
				}
				if !ok {
					return errors.New("推荐信息不存在")
				}
				passageRecommend.Path = fmt.Sprintf("%s-%d", parentRecommend.Path, parentRecommend.Id)
			} else {
				passageRecommend.Path = "0"
			}
			passageRecommend.AccountId = utils.IntVal(accountId)
			passageRecommend.ParentPassageRecommendId = request.ParentPassageRecommendId
			passageRecommend.PassageId = request.ID
			passageRecommend.CreatedTime = time.Now().Unix()
			passageRecommend.UpdatedTime = time.Now().Unix()
			_, err = db.Default().Insert(&passageRecommend)
			if err != nil {
				return err
			}
			passageRecommend.PathFull = passageRecommend.GetFullPath()
			_, err := db.Default().Table(passageRecommend).ID(passageRecommend.Id).Cols("path_full").Update(passageRecommend)
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
	services.ResponseSuccess(c, gin.H{
		"passageRecommendId": passageRecommend.Id,
	})
}
