package services

import (
	"fmt"
	"testing"
)

func TestWechatAccessToken(t *testing.T) {
	fmt.Println(WechatAccessToken())
	t.Fail()
}
