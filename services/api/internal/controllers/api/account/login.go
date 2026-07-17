package account

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"
	"app/pkg/utils"
	"errors"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Mobile string `json:"mobile"`
}

func LoginAction(c *gin.Context) {
	var request LoginRequest
	data := gin.H{}
	err := func() error {
		if err := c.ShouldBindJSON(&request); err != nil {
			return err
		}
		var user bolejiang.Account
		ok, err := db.Get(db.Default().Where("mobile = ?", request.Mobile), &user)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("用户未注册")
		}
		data["token"], _ = utils.GetToken(user.Id)
		return nil
	}()
	if err != nil {
		services.ResponseError(c, -1, err.Error(), nil)
		return
	}
	services.ResponseSuccess(c, data)
}
