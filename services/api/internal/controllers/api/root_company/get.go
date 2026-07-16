package root_company

import (
	"app/internal/db/mysql/query"
	"app/internal/services"

	"github.com/gin-gonic/gin"
)

type GetRequest struct {
	ID uint32 `json:"id"`
}

func GetAction(c *gin.Context) {
	var request GetRequest
	var companyFull *services.RootCompanyFull
	err := func() error {
		if err := c.ShouldBindJSON(&request); err != nil {
			return err
		}
		rootCompany, err := query.CompanyInfo.Where(query.CompanyInfo.ID.Eq(request.ID)).First()
		if err != nil {
			return err
		}
		companyFull, err = services.RootCompanyFullGet(rootCompany)
		if err != nil {
			return err
		}
		return nil
	}()
	if err != nil {
		services.ResponseError(c, -1, err.Error(), nil)
		return
	}
	services.ResponseSuccess(c, companyFull)
}
