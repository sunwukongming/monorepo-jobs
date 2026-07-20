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
	services.Handle(c, func() (interface{}, error) {
		var request LoginRequest
		data := gin.H{}
		if err := c.ShouldBindJSON(&request); err != nil {
			return nil, err
		}
		var user bolejiang.Account
		ok, err := db.Get(db.Default().Where("mobile = ?", request.Mobile), &user)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, errors.New("用户未注册")
		}
		data["token"], _ = utils.GetToken(user.Id)
		return data, nil
	})
}
