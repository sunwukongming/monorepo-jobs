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

func InvestmentDemandDetailAction(c *gin.Context) {
	type Request struct {
		ID int `json:"id"`
	}
	type Response struct {
		bolejiang.InvestmentDemand
		Passages []services.PassageFull `json:"passages"`
	}
	var response Response
	var investmentDemand bolejiang.InvestmentDemand
	err := func() error {
		var request Request
		if err := c.ShouldBindJSON(&request); err != nil {
			return err
		}

		ok, err := db.Default().Table(bolejiang.InvestmentDemand{}).ID(request.ID).Get(&investmentDemand)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("数据不存在")
		}

		response.InvestmentDemand = investmentDemand
		if investmentDemand.PassageIds != "" {
			passages, err := services.PassageListFullByIDs(
				utils.Uint32ArrayFromStringArray(strings.Split(investmentDemand.PassageIds, ",")),
				uint32(utils.IntVal(services.AuthGetAccountID(c))),
			)
			if err != nil {
				return err
			}
			response.Passages = passages
		}

		db.Default().Exec("update investment_demand set view_count = view_count + 1 where id = ?", request.ID)
		return nil
	}()
	if err != nil {
		services.ResponseError(c, -1, err.Error(), nil)
		return
	}
	services.ResponseSuccess(c, response)
}
