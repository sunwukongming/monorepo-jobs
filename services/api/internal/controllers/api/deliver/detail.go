package deliver

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"
	"errors"

	"github.com/gin-gonic/gin"
)

func DetailAction(c *gin.Context) {
	type Request struct {
		ID string `json:"id"`
	}
	var request Request
	var deliverPassage DeliverPassage
	err := func() error {
		if err := c.ShouldBindJSON(&request); err != nil {
			return err
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

		ok, err = db.Default().Table(bolejiang.Deliver{}).
			Join("LEFT", "passage", "deliver.passage_id = passage.id").
			Join("LEFT", "account", "deliver.account_id = account.id").
			Where("deliver.id = ? and deliver.recommend_account_id = ? and deliver.is_real = 1", request.ID, currentAccount.Id).Get(&deliverPassage)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("该投递不存在")
		}
		return nil
	}()
	if err != nil {
		services.ResponseError(c, -1, err.Error(), nil)
		return
	}
	services.ResponseSuccess(c, gin.H{
		"deliver": deliverPassage.Deliver,
		"passage": deliverPassage.Passage,
	})
}
