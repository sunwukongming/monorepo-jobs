package passage

import (
	"app/internal/services"
	"app/pkg/utils"
	"errors"

	"github.com/gin-gonic/gin"
)

func GetOriginAction(c *gin.Context) {
	type Request struct {
		ID uint32 `json:"id"`
	}
	services.Handle(c, func() (interface{}, error) {
		var request Request
		if err := c.ShouldBindJSON(&request); err != nil {
			return nil, err
		}
		if request.ID == 0 {
			return nil, errors.New("职位id必须填写")
		}
		accountId := services.AuthGetAccountID(c)
		passageResponse, err := services.PassageGetFullByID(request.ID, uint32(utils.IntVal(accountId)))
		if err != nil {
			return nil, err
		}
		return passageResponse, nil
	})
}
