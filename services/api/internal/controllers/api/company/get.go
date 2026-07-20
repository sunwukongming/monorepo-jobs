package company

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
		if err := c.ShouldBindJSON(&request); err != nil {
			return nil, err
		}
		passageCompany, err := query.PassageCompany.Where(query.PassageCompany.ID.Eq(request.ID)).First()
		if err != nil {
			return nil, err
		}
		return services.CompanyFullGet(passageCompany)
	})
}
