package services

import (
	"sync"
	"time"

	"app/config"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
)

var wechatAccessToken string
var wechatAccessTokenCreatedTime int64
var wechatMutex sync.Mutex

// 历史硬编码值，仅作为「配置未提供 wechat 段」时的回退，保证现网不中断。
// 待所有环境的 config 均补齐 wechat 段后，应删除这两个常量、改为强依赖配置。
const fallbackWechatAppID = "wx1cb6a1ca70e8dbfa"
const fallbackWechatAppSecret = "eb48ed536cbbc7bce1ecad2f342a306d"

// WechatAppID 返回小程序 AppID：优先取 config 的 wechat.appId，未配置则回退历史常量。
func WechatAppID() string {
	if v := config.Get().Wechat.AppID; v != "" {
		return v
	}
	return fallbackWechatAppID
}

// WechatAppSecret 返回小程序 AppSecret：优先取 config 的 wechat.appSecret，未配置则回退。
func WechatAppSecret() string {
	if v := config.Get().Wechat.AppSecret; v != "" {
		return v
	}
	return fallbackWechatAppSecret
}

func WechatAccessToken() string {
	if time.Now().Unix()-wechatAccessTokenCreatedTime < 3600 && wechatAccessToken != "" {
		return wechatAccessToken
	}
	wechatMutex.Lock()
	defer wechatMutex.Unlock()
	client := resty.New()
	resp, err := client.R().SetQueryParams(map[string]string{
		"grant_type": "client_credential",
		"appid":      WechatAppID(),
		"secret":     WechatAppSecret(),
	}).Get("https://api.weixin.qq.com/cgi-bin/token")
	if err != nil {
		return ""
	}
	body := resp.Body()
	wechatAccessToken := jsoniter.Get(body, "access_token").ToString()
	wechatAccessTokenCreatedTime = time.Now().Unix()
	return wechatAccessToken
}
