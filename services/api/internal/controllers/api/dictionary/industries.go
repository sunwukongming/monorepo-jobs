package dictionary

import (
	"app/data"
	"app/internal/services"

	"github.com/gin-gonic/gin"
)

func IndustriesAction(c *gin.Context) {
	services.ResponseSuccess(c, gin.H{
		"list": data.GetComposedIndustries(),
	})
}
