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
	var data any

	err := func() error {
		var request Request
		if err := c.ShouldBindJSON(&request); err != nil {
			return err
		}
		if request.ID == 0 {
			return errors.New("职位id必须填写")
		}
		accountId := services.AuthGetAccountID(c)
		passageResponse, err := services.PassageGetFullByID(request.ID, uint32(utils.IntVal(accountId)))
		if err != nil {
			return err
		}
		data = passageResponse
		return nil
	}()
	if err != nil {
		services.ResponseError(c, -1, err.Error(), nil)
		return
	}
	services.ResponseSuccess(c, data)
}
