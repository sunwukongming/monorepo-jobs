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
	services.Handle(c, func() (interface{}, error) {
		var request GetRequest
		var companyFull *services.RootCompanyFull
		if err := c.ShouldBindJSON(&request); err != nil {
			return nil, err
		}
		rootCompany, err := query.CompanyInfo.Where(query.CompanyInfo.ID.Eq(request.ID)).First()
		if err != nil {
			return nil, err
		}
		companyFull, err = services.RootCompanyFullGet(rootCompany)
		if err != nil {
			return nil, err
		}
		return companyFull, nil
	})
}
