package account

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

type BindWechatMobileRequest struct {
	Code string `json:"code"`
}

func BindWechatMobileAction(c *gin.Context) {
	services.Handle(c, func() (interface{}, error) {
		var request BindWechatMobileRequest
		data := gin.H{}
		if err := c.ShouldBindJSON(&request); err != nil {
			return nil, err
		}
		accountId := services.AuthGetAccountID(c)
		if accountId == "" {
			return nil, errors.New("用户不存在")
		}
		bs, _ := jsoniter.MarshalToString(map[string]string{
			"code": request.Code,
		})
		client := resty.New()
		resp, err := client.R().
			SetBody(bs).
			SetQueryParam("access_token", services.WechatAccessToken()).
			SetHeader("Content-Type", "application/json").Post("https://api.weixin.qq.com/wxa/business/getuserphonenumber")
		if err != nil {
			return nil, err
		}
		body := resp.Body()
		errcode := jsoniter.Get(body, "errcode").ToInt()
		if errcode != 0 {
			errmsg := jsoniter.Get(body, "errmsg").ToString()
			return nil, errors.New(errmsg)
		}
		mobile := jsoniter.Get(body, "phone_info", "purePhoneNumber").ToString()
		logrus.WithFields(logrus.Fields{
			"action": "wechat_mobile",
			"body":   string(body),
			"mobile": mobile,
		}).Info()
		var account bolejiang.Account
		ok, err := db.Get(db.Default().Where("id = ?", accountId), &account)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, errors.New("用户不存在")
		}
		err = services.AccountBindMobile(account, mobile)
		if err != nil {
			return nil, err
		}

		services.AccountUpdateRelevant(account)
		return data, nil
	})
}
