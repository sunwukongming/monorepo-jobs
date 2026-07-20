package services

import (
	"app/config"
	"app/pkg/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// ResponseSuccess 返回成功
func ResponseSuccess(c *gin.Context, data interface{}) {
	respond(c, 0, "ok", data)
}

// ResponseError 返回失败（code 由调用方显式传入）
func ResponseError(c *gin.Context, code int, message string, data interface{}) {
	respond(c, code, message, data)
}

// respond 统一构造并输出响应体：数据转 camelCase、写入 code/message 供中间件日志读取、
// 计算耗时、debug 模式回填 response、以 HTTP 200 输出。
func respond(c *gin.Context, code int, message string, data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	d, _ := utils.InterfaceToCamelCase(data)
	c.Set("code", code)
	c.Set("message", message)

	rs := gin.H{
		"code":      code,
		"message":   message,
		"data":      d,
		"timestamp": time.Now().Unix(),
		"latency":   elapsedSeconds(c),
	}
	if config.Get().Mode == "debug" {
		c.Set("response", rs)
	}
	c.JSON(http.StatusOK, rs)
}

// elapsedSeconds 依据 context 中的 startTime 计算已用秒数（缺失则返回 0）。
func elapsedSeconds(c *gin.Context) float64 {
	value, ok := c.Get("startTime")
	if !ok {
		return 0
	}
	startTime, ok := value.(time.Time)
	if !ok {
		return 0
	}
	return time.Since(startTime).Seconds()
}
