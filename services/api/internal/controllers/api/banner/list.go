package banner

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"

	"github.com/gin-gonic/gin"
)

func ListAction(c *gin.Context) {
	services.Handle(c, func() (interface{}, error) {
		banners := make([]bolejiang.Banner, 0)
		if err := db.Default().Find(&banners).Error; err != nil {
			return nil, err
		}
		return gin.H{"list": banners}, nil
	})
}
