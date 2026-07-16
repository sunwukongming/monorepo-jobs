package dictionary

import (
	"app/data"
	"app/internal/services"

	"github.com/gin-gonic/gin"
)

func DataAction(c *gin.Context) {
	services.ResponseSuccess(c, data.DictionaryMap)
}
