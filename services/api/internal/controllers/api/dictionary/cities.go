package dictionary

import (
	"app/data"
	"app/internal/services"

	"github.com/gin-gonic/gin"
)

func CitiesAction(c *gin.Context) {
	services.ResponseSuccess(c, gin.H{
		"list": data.GetComposedCities(),
	})
}
