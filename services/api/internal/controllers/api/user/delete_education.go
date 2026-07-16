package user

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"
	"errors"

	"github.com/gin-gonic/gin"
)

func DeleteEducationAction(c *gin.Context) {
	type Request struct {
		Id int `json:"id"`
	}
	var request Request
	var accountEducation bolejiang.AccountEducation
	err := func() error {
		if err := c.ShouldBindJSON(&request); err != nil {
			return err
		}
		id := services.AuthGetAccountID(c)
		var user bolejiang.Account
		ok, err := db.Default().Where("id = ?", id).Get(&user)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("用户不存在")
		}
		ok, err = db.Default().Where("id = ? and account_id = ?", request.Id, user.Id).Get(&accountEducation)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("求职目标不存在")
		}
		_, err = db.Default().Table(bolejiang.AccountEducation{}).Where("id = ?", accountEducation.Id).Delete()
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
