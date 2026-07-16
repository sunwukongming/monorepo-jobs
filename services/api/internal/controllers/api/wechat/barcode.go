package wechat

import (
	"app/internal/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
)

func BarcodeAction(c *gin.Context) {
	var body []byte
	err := func() error {
		client := resty.New()
		b, _ := jsoniter.Marshal(map[string]interface{}{
			"page":       c.Query("page"),
			"scene":      c.Query("scene"),
			"check_path": false,
		})
		fmt.Println(services.WechatAccessToken())
		resp, err := client.R().SetHeader("Content-Type", "application/json").SetQueryParams(map[string]string{
			"access_token": services.WechatAccessToken(),
		}).SetBody(b).Post("https://api.weixin.qq.com/wxa/getwxacodeunlimit")
		if err != nil {
			return err
		}
		body = resp.Body()
		return nil
	}()
	if err != nil {
		services.ResponseError(c, -1, err.Error(), nil)
		return
	}
	c.Data(http.StatusOK, "image/png", body)
}
