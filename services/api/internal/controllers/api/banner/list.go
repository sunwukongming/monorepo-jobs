package banner

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"

	"github.com/gin-gonic/gin"
)

func ListAction(c *gin.Context) {
	data := gin.H{}
	err := func() error {
		banners := make([]bolejiang.Banner, 0)
		err := db.Default().Find(&banners).Error
		if err != nil {
			return err
		}
		data["list"] = banners
		return nil
	}()
	if err != nil {
		services.ResponseError(c, -1, err.Error(), nil)
		return
	}
	services.ResponseSuccess(c, data)
}
