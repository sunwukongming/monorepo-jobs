package wechat

import (
	"app/internal/services"
	"bytes"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
)

func URLSchemaAction(c *gin.Context) {
	services.Handle(c, func() (interface{}, error) {
		var body []byte
		res := gin.H{}
		client := resty.New()
		b, _ := jsoniter.Marshal(map[string]interface{}{
			"jump_wxa": map[string]interface{}{
				"path":        c.Query("path"),
				"query":       c.Query("query"),
				"env_version": c.Query("version"),
			},
		})
		b = bytes.ReplaceAll(b, []byte(`\u0026`), []byte("&"))
		fmt.Println(services.WechatAccessToken())
		resp, err := client.R().SetHeader("Content-Type", "application/json").SetQueryParams(map[string]string{
			"access_token": services.WechatAccessToken(),
		}).SetBody(b).Post("https://api.weixin.qq.com/wxa/generatescheme")
		if err != nil {
			return nil, err
		}
		body = resp.Body()
		if jsoniter.Get(body, "errcode").ToInt() != 0 {
			return nil, errors.New(jsoniter.Get(body, "errmsg").ToString())
		}
		res["data"] = gin.H{
			"url": jsoniter.Get(body, "openlink").ToString(),
		}
		return res, nil
	})
}
