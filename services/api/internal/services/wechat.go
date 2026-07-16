package services

import (
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
)

var wechatAccessToken string
var wechatAccessTokenCreatedTime int64
var wechatMutex sync.Mutex

//var WechatAppId = "wxd225a74d6c05cd9e"
//var WechatAppsecret = "5fa4fe648bff26c2bfa1ca8468e879df"

var WechatAppId = "wx1cb6a1ca70e8dbfa"
var WechatAppsecret = "eb48ed536cbbc7bce1ecad2f342a306d"

func WechatAccessToken() string {
	if time.Now().Unix()-wechatAccessTokenCreatedTime < 3600 && wechatAccessToken != "" {
		return wechatAccessToken
	}
	wechatMutex.Lock()
	defer wechatMutex.Unlock()
	client := resty.New()
	resp, err := client.R().SetQueryParams(map[string]string{
		"grant_type": "client_credential",
		"appid":      WechatAppId,
		"secret":     WechatAppsecret,
	}).Get("https://api.weixin.qq.com/cgi-bin/token")
	if err != nil {
		return ""
	}
	body := resp.Body()
	wechatAccessToken := jsoniter.Get(body, "access_token").ToString()
	wechatAccessTokenCreatedTime = time.Now().Unix()
	return wechatAccessToken
}
