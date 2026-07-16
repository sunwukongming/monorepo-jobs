/**

Filename: 		time.go
Author: 		gengming - gengming.zb@ccbft.com
Description:	loongrpc time tool logic
Create:			2022-07-02 14:33:10
Last Modified:	2022-07-02 15:01:32

*/

package utils

import (
	"fmt"
	"time"
	_ "time/tzdata"
)

var TimeZone = "Asia/Shanghai"

var TimeLocation *time.Location

func init() {
	TimeLocation, _ = time.LoadLocation(TimeZone)
}

func TimeLoadLocation(name string) {
	TimeZone = name
	TimeLocation, _ = time.LoadLocation(name)
}

// TimeToStr 时间转字符串
func TimeToStr(timestamp int64) string {
	return time.Unix(timestamp, 0).In(TimeLocation).Format("2006-01-02 15:04:05")
}

// TimeToDay 时间转日期字符串
func TimeToDay(timestamp int64) string {
	return time.Unix(timestamp, 0).In(TimeLocation).Format("2006-01-02")
}

func TimeSecondHumanize(second int64) string {
	s := second % 60
	m := (second % 3600) / 60
	h := second / 3600
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}
