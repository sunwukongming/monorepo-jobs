package article

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"
	"app/pkg/utils"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func CooperationDetailAction(c *gin.Context) {
	type Request struct {
		ID int `json:"id"`
	}
	type Response struct {
		bolejiang.Cooperation
		Passages []services.PassageFull `json:"passages"`
	}
	var response Response
	var cooperation bolejiang.Cooperation
	err := func() error {
		var request Request
		if err := c.ShouldBindJSON(&request); err != nil {
			return err
		}

		ok, err := db.Default().Table(bolejiang.Cooperation{}).ID(request.ID).Get(&cooperation)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("数据不存在")
		}

		response.Cooperation = cooperation
		if cooperation.PassageIds != "" {
			passages, err := services.PassageListFullByIDs(
				utils.Uint32ArrayFromStringArray(strings.Split(cooperation.PassageIds, ",")),
				uint32(utils.IntVal(services.AuthGetAccountID(c))),
			)
			if err != nil {
				return err
			}
			response.Passages = passages
		}
		db.Default().Exec("update cooperation set view_count = view_count + 1 where id = ?", request.ID)
		return nil
	}()
	if err != nil {
		services.ResponseError(c, -1, err.Error(), nil)
		return
	}
	services.ResponseSuccess(c, response)
}
