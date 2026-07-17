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
	var request LoginWechatRequest
	data := gin.H{}
	err := func() error {
		if err := c.ShouldBindJSON(&request); err != nil {
			return err
		}
		client := resty.New()
		resp, err := client.R().SetQueryParams(map[string]string{
			"appid":      services.WechatAppId,
			"secret":     services.WechatAppsecret,
			"js_code":    request.Code,
			"grant_type": "authorization_code",
		}).Get("https://api.weixin.qq.com/sns/jscode2session")
		if err != nil {
			return err
		}
		body := resp.Body()
		errcode := jsoniter.Get(body, "errcode").ToInt()
		if errcode != 0 {
			errmsg := jsoniter.Get(body, "errmsg").ToString()
			return errors.New(errmsg)
		}
		openId := jsoniter.Get(body, "openid").ToString()
		unionId := jsoniter.Get(body, "unionid").ToString()
		var account bolejiang.Account
		ok, err := db.Get(db.Default().Where("unionid = ?", unionId), &account)
		if err != nil {
			return err
		}
		if !ok {
			account.Openid = openId
			account.Unionid = unionId
			account.CreatedTime = time.Now().Unix()
			account.UpdatedTime = time.Now().Unix()
			err := db.Default().Create(&account).Error
			if err != nil {
				return err
			}
		}
		data["token"], _ = utils.GetToken(account.Id)
		data["openid"] = openId
		data["unionid"] = unionId
		data["sessionKey"] = jsoniter.Get(body, "session_key")
		return nil
	}()
	if err != nil {
		services.ResponseError(c, -1, err.Error(), nil)
		return
	}
	services.ResponseSuccess(c, data)
}
