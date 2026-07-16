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

func FinancingDemandDetailAction(c *gin.Context) {
	type Request struct {
		ID int `json:"id"`
	}
	type Response struct {
		bolejiang.FinancingDemand
		Passages []services.PassageFull `json:"passages"`
	}
	var response Response
	var financingDemand bolejiang.FinancingDemand
	err := func() error {
		var request Request
		if err := c.ShouldBindJSON(&request); err != nil {
			return err
		}

		ok, err := db.Default().Table(bolejiang.FinancingDemand{}).ID(request.ID).Get(&financingDemand)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("数据不存在")
		}

		response.FinancingDemand = financingDemand
		if financingDemand.PassageIds != "" {
			passages, err := services.PassageListFullByIDs(
				utils.Uint32ArrayFromStringArray(strings.Split(financingDemand.PassageIds, ",")),
				uint32(utils.IntVal(services.AuthGetAccountID(c))),
			)
			if err != nil {
				return err
			}
			response.Passages = passages
		}
		db.Default().Exec("update financing_demand set view_count = view_count + 1 where id = ?", request.ID)
		return nil
	}()
	if err != nil {
		services.ResponseError(c, -1, err.Error(), nil)
		return
	}
	services.ResponseSuccess(c, response)
}
