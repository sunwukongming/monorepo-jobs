package services

import (
	"app/config"
	"app/pkg/utils"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// ResponseSuccess 返回成功
func ResponseSuccess(c *gin.Context, data interface{}) {
	c.Set("code", 0)
	c.Set("message", "ok")
	if data == nil {
		data = gin.H{}
	}
	d, _ := utils.InterfaceToCamelCase(data)
	startTime, ok := c.Get("startTime")
	var useTime float64
	if ok {
		s, ok := startTime.(time.Time)
		if ok {
			e := time.Now()
			latency := e.Sub(s)
			useTime = latency.Seconds()
		}
	}
	rs := gin.H{
		"code":      0,
		"message":   "ok",
		"data":      d,
		"timestamp": time.Now().Unix(),
		"latency":   useTime,
	}

	if config.Get().Mode == "debug" {
		c.Set("response", rs)
	}

	c.JSON(http.StatusOK, rs)
}

// ResponseError 返回失败
func ResponseError(c *gin.Context, code int, message string, data interface{}) {
	fmt.Println(message)
	a := strings.Split(message, "|")
	if len(a) >= 2 {
		c, err := strconv.Atoi(a[0])
		if err == nil {
			code = c
			message = a[1]
		}
	}
	if data == nil {
		data = gin.H{}
	}
	d, _ := utils.InterfaceToCamelCase(data)
	c.Set("code", code)
	c.Set("message", message)
	startTime, ok := c.Get("startTime")
	fmt.Println("startTime", startTime, ok)
	var useTime float64
	if ok {
		s, ok := startTime.(time.Time)
		if ok {
			e := time.Now()
			latency := e.Sub(s)
			useTime = latency.Seconds()
		}
	}

	rs := gin.H{
		"code":      code,
		"message":   message,
		"data":      d,
		"timestamp": time.Now().Unix(),
		"latency":   useTime,
	}
	if config.Get().Mode == "debug" {
		c.Set("response", rs)
	}

	c.JSON(http.StatusOK, rs)

}
