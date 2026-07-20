package user

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"
	"errors"

	"github.com/gin-gonic/gin"
)

func DeleteEducationAction(c *gin.Context) {
	services.Handle(c, func() (interface{}, error) {
		type Request struct {
			Id int `json:"id"`
		}
		var request Request
		var accountEducation bolejiang.AccountEducation
		if err := c.ShouldBindJSON(&request); err != nil {
			return nil, err
		}
		id := services.AuthGetAccountID(c)
		var user bolejiang.Account
		ok, err := db.Get(db.Default().Where("id = ?", id), &user)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, errors.New("用户不存在")
		}
		ok, err = db.Get(db.Default().Where("id = ? and account_id = ?", request.Id, user.Id), &accountEducation)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, errors.New("求职目标不存在")
		}
		err = db.Default().Where("id = ?", accountEducation.Id).Delete(&bolejiang.AccountEducation{}).Error
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
}
