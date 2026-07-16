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

func IndustryAssociationDetailAction(c *gin.Context) {
	type Request struct {
		ID int `json:"id"`
	}
	type Response struct {
		bolejiang.IndustryAssociation
		Passages []services.PassageFull `json:"passages"`
	}
	var response Response
	var industryAssociation bolejiang.IndustryAssociation
	err := func() error {
		var request Request
		if err := c.ShouldBindJSON(&request); err != nil {
			return err
		}

		ok, err := db.Default().Table(bolejiang.IndustryAssociation{}).ID(request.ID).Get(&industryAssociation)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("数据不存在")
		}

		response.IndustryAssociation = industryAssociation
		if industryAssociation.PassageIds != "" {
			passages, err := services.PassageListFullByIDs(
				utils.Uint32ArrayFromStringArray(strings.Split(industryAssociation.PassageIds, ",")),
				uint32(utils.IntVal(services.AuthGetAccountID(c))),
			)
			if err != nil {
				return err
			}
			response.Passages = passages
		}
		db.Default().Exec("update industry_association set view_count = view_count + 1 where id = ?", request.ID)
		return nil
	}()
	if err != nil {
		services.ResponseError(c, -1, err.Error(), nil)
		return
	}
	services.ResponseSuccess(c, response)
}
