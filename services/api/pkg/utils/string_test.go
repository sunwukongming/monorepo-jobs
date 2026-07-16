package utils

import (
	"encoding/base64"
	"fmt"
	"path"
	"regexp"
	"strings"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func TestCamelCase(t *testing.T) {
	for i := 0; i < 10000000; i++ {
		s := CamelCase("good_bye")
		fmt.Println(s)
		if s != "goodBye" {
			t.Fail()
		}
		s = CamelCase("Good_bye")
		if s != "GoodBye" {
			t.Fail()
		}
	}
}

func TestSnakeCase(t *testing.T) {
	s := SnakeCase("realName")
	if s != "real_name" {
		t.Fail()
	}
}

func TestUniqueId(t *testing.T) {
	if UniqueID() == UniqueID() {
		t.Fail()
	}
}

func TestStringSub(t *testing.T) {
	s := "abcdef"
	if StringSub(s, 0, 3) != "abc" {
		t.Fail()
	}
	s = "你好，我是成功人士"
	if StringSub(s, 2, 3) != "，我是" {
		t.Fail()
	}
	if StringSub(s, -1, 3) != "士" {
		t.Fail()
	}
	if StringSub(s, 9, 3) != "" {
		t.Fail()
	}
	if StringSub(s, 0, 100) != s {
		t.Fail()
	}
}

func TestBase64(t *testing.T) {
	s := "ZW52X2lkPTE3ODcwMjgzNDUzMzY5Njg4OTc4Jm1zZ19pZD0yMDIwMTIwOTE0MTExMS4xMzM2NTkzNS5tY0BtYWlsLmhpbmFwb3dlci5jbiZhY2NvdW50PW5vcmVwbHlAbWFpbC5oaW5hcG93ZXIuY24mZnJvbT1ub3JlcGx5QG1haWwuaGluYXBvd2VyLmNuJnJjcHQ9MjkwOTQwODEyQHFxLmNvbSZyZWN2X3RpbWU9MjAyMC0xMi0wOSAxNDoxMToxMSZlbmRfdGltZT0yMDIwLTEyLTA5IDE0OjExOjEzJnN0YXR1cz0wJmV2ZW50PWRlbGl2ZXImcmVnaW9uPWNuLXNoYW5naGFpJmVycl9jb2RlPTI1MCZlcnJfbXNnPTI1MCBTZW5kIE1haWwgT0smZmFpbGVkX3R5cGU9U2VuZE9r"
	d, err := base64.StdEncoding.DecodeString(s)
	t.Log(string(d))
	t.Log(err)
	t.Fail()
}

func TestStringRandom(t *testing.T) {
	fmt.Println(StringRandom(StringAlnumPassword, 100))
}

func TestStringConcat(t *testing.T) {
	s := StringConcat("/", "path/", "to/", "/name")
	if s != "path/to/name" {
		t.Log(s)
		t.Fail()
	}
}

func TestReplace(t *testing.T) {
	replacer := strings.NewReplacer(
		"#name#", "名字",
		"#候选人姓名#", "名字",
		"#position#", "职位",
		"#职位名称#", "职位",
	)

	emailContent := replacer.Replace("<p>候选人姓名:#name#</p>↵↵<p>职位名:#position##position##position#</p>↵↵<p>公司名:#company#</p>↵↵<p>面试时间:")
	fmt.Println(emailContent)
}

func TestJsonSize(t *testing.T) {
	a := `{
        "errcode":0,
        "auth_user_info":{
                "userId":"manager975"
        },
        "auth_corp_info":{
                "corp_type":0,
                "corpid":"dingbef1744b54241496ee0f45d8xxxx",
                "auth_level":0,
                "auth_channel":"",
                "industry":"",
                "full_corp_name":"小程序体验HTTP",
                "corp_name":"小程序体验HTTP",
                "invite_url":"https://wx.dingtalk.com/invite-page/index.html?bizSource=____source____&corpId=dingbef1744b54241496ee0f45d8e4f7c288&inviterUid=DFAD06727FD38CD894460A2FDF52346D",
                "auth_channel_type":"",
                "invite_code":"",
                "is_authenticated":false,
                "license_code":"",
                "corp_logo_url":""
        },
        "errmsg":"ok",
        "channel_auth_info":{
                "channelAgent":[]
        },
        "auth_info":{
                "agent":[
                        {
                                "agentid":852381775,
                                "agent_name":"小程序DEMO",
                                "logo_url":"https://static-legacy.dingtalk.com/media/lADPDefRxxj6mvVUVg_86_84.jpg",
                                "appid":53642,
                                "admin_list":[
                                        "manager975"
                                ]
                        },
						{
							"agentid":852381775,
							"agent_name":"小程序DEMO",
							"logo_url":"https://static-legacy.dingtalk.com/media/lADPDefRxxj6mvVUVg_86_84.jpg",
							"appid":53642,
							"admin_list":[
									"manager975"
							]
					}
                ]
        },
        "auth_market_info":{}
}`
	bs := []byte(a)
	agent := jsoniter.Get(bs, "auth_info", "agent")
	size := agent.Size()
	for i := 0; i < size; i++ {
		fmt.Println(agent.Get(i, "agentid").ToInt())
		agent.Get(i, "agent_name").ToString()
		agent.Get(i, "logo_url").ToString()
		agent.Get(i, "appid").ToInt()
		fmt.Println(agent.Get(i, "admin_list").ToString())
	}
	t.Fail()
}

func TestReg2(t *testing.T) {
	s := " 一         .         基本信息                       姓名     :      耿明             性别     :      男             毕业院校     :      武汉大学              所在地: 北京"
	nameRegexp := regexp.MustCompile("[^\u4e00-\u9fa5]{1}[\u4e00-\u9fa5]{2,4}\\s")
	a := nameRegexp.FindAllString(s, -1)
	for _, item := range a {
		fmt.Println(item)
	}

	t.Fail()
}

func TestExt(t *testing.T) {
	fmt.Println(path.Ext("adfa.PDF"))
	t.Fail()
}
