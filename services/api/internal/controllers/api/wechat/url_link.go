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

func UrlLinkAction(c *gin.Context) {
	services.Handle(c, func() (interface{}, error) {
		var body []byte
		res := gin.H{}
		client := resty.New()
		data := gin.H{}
		if c.Query("path") != "" {
			data["path"] = c.Query("path")
		}
		if c.Query("query") != "" {
			data["query"] = c.Query("query")
		}
		if c.Query("version") != "" {
			data["env_version"] = c.Query("version")
		}

		b, _ := jsoniter.Marshal(data)
		b = bytes.ReplaceAll(b, []byte(`\u0026`), []byte("&"))
		fmt.Println(services.WechatAccessToken())
		resp, err := client.R().SetHeader("Content-Type", "application/json").SetQueryParams(map[string]string{
			"access_token": services.WechatAccessToken(),
		}).SetBody(b).Post("https://api.weixin.qq.com/wxa/generate_urllink")
		if err != nil {
			return nil, err
		}
		body = resp.Body()
		if jsoniter.Get(body, "errcode").ToInt() != 0 {
			return nil, errors.New(jsoniter.Get(body, "errmsg").ToString())
		}
		res["data"] = gin.H{
			"url": jsoniter.Get(body, "url_link").ToString(),
		}
		return res, nil
	})
}
