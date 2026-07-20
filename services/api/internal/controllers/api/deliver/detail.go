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
	services.Handle(c, func() (interface{}, error) {
		var request Request
		var deliver bolejiang.Deliver
		var passage bolejiang.Passage
		if err := c.ShouldBindJSON(&request); err != nil {
			return nil, err
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

		// 原 xorm `extends` 多表嵌入扫描 GORM 不支持，改为先取 deliver 再按 passage_id 取职位。
		ok, err := db.Get(db.Default().
			Where("id = ? and recommend_account_id = ? and is_real = 1", request.ID, currentAccount.Id), &deliver)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, errors.New("该投递不存在")
		}
		_, err = db.Get(db.Default().Where("id = ?", deliver.PassageId), &passage)
		if err != nil {
			return nil, err
		}
		return gin.H{
			"deliver": deliver,
			"passage": passage,
		}, nil
	})
}
