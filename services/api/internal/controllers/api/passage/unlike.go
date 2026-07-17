package passage

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"
	"errors"

	"github.com/gin-gonic/gin"
)

type UnlikeRequest struct {
	ID int `json:"id"`
}

func UnlikeAction(c *gin.Context) {
	var request UnlikeRequest
	err := func() error {
		if err := c.ShouldBindJSON(&request); err != nil {
			return err
		}
		accountId := services.AuthGetAccountID(c)
		if accountId == "" {
			return errors.New("用户不存在")
		}
		var passage bolejiang.Passage
		ok, err := db.Get(db.Default().Where("id = ?", request.ID), &passage)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("职位不存在")
		}
		var passageLike bolejiang.PassageLike
		ok, err = db.Get(db.Default().Where("account_id = ? and passage_id = ?", accountId, request.ID), &passageLike)
		if err != nil {
			return err
		}
		if ok {
			err := db.Default().Where("id = ?", passageLike.Id).Delete(&bolejiang.PassageLike{}).Error
			if err != nil {
				return err
			}
		}

		//更新用户收藏
		err = db.Default().Exec("update account as a set collect = (select count(*) from passage_like where account_id = a.id) where id = ?", accountId).Error
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
