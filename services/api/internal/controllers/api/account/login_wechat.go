package account

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"
	"app/pkg/utils"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
)

type LoginWechatRequest struct {
	Code string `json:"code" binding:"required"`
}

func LoginWechatAction(c *gin.Context) {
	services.Handle(c, func() (interface{}, error) {
		var request LoginWechatRequest
		data := gin.H{}
		if err := c.ShouldBindJSON(&request); err != nil {
			return nil, err
		}
		client := resty.New()
		resp, err := client.R().SetQueryParams(map[string]string{
			"appid":      services.WechatAppID(),
			"secret":     services.WechatAppSecret(),
			"js_code":    request.Code,
			"grant_type": "authorization_code",
		}).Get("https://api.weixin.qq.com/sns/jscode2session")
		if err != nil {
			return nil, err
		}
		body := resp.Body()
		errcode := jsoniter.Get(body, "errcode").ToInt()
		if errcode != 0 {
			errmsg := jsoniter.Get(body, "errmsg").ToString()
			return nil, errors.New(errmsg)
		}
		openId := jsoniter.Get(body, "openid").ToString()
		unionId := jsoniter.Get(body, "unionid").ToString()
		var account bolejiang.Account
		ok, err := db.Get(db.Default().Where("unionid = ?", unionId), &account)
		if err != nil {
			return nil, err
		}
		if !ok {
			account.Openid = openId
			account.Unionid = unionId
			account.CreatedTime = time.Now().Unix()
			account.UpdatedTime = time.Now().Unix()
			err := db.Default().Create(&account).Error
			if err != nil {
				return nil, err
			}
		}
		data["token"], _ = utils.GetToken(account.Id)
		data["openid"] = openId
		data["unionid"] = unionId
		data["sessionKey"] = jsoniter.Get(body, "session_key")
		return data, nil
	})
}
