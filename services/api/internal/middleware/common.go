package middleware

import (
	"app/internal/services"
	"bytes"
	"io"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

func Common() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Set("startTime", startTime)
		contentType := c.GetHeader("Content-Type")
		var bd interface{}
		if strings.Contains(contentType, "application/json") || contentType == "" {
			body, _ := c.GetRawData()
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			_ = jsoniter.Unmarshal(body, &bd)
		} else if strings.Contains(contentType, "application/x-www-form-urlencoded") {
			body, _ := c.GetRawData()
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			bd = string(body)
		} else {
			form, err := c.MultipartForm()
			if err == nil {
				bd = form.Value
			}
		}

		accountId := services.AuthGetAccountID(c)
		if accountId != "" {
			account, err := services.AuthGetAccountOrError(c)
			if err != nil {
				c.Abort()
				services.ResponseError(c, 401, err.Error(), nil)
				return
			}
			c.Set("account", account)
		}

		c.Next()
		endTime := time.Now()
		latency := endTime.Sub(startTime)
		code, _ := c.Get("code")
		message, _ := c.Get("message")
		_, isIgnore := c.Get("ignore")
		response, _ := c.Get("response")
		if !isIgnore {
			l := logrus.WithFields(logrus.Fields{
				"action":    "request",
				"url":       c.Request.URL.Path,
				"query":     c.Request.URL.RawQuery,
				"post":      bd,
				"code":      code,
				"response":  response,
				"ua":        c.GetHeader("User-Agent"),
				"useTime":   float32(latency.Microseconds()) / 1000000,
				"timestamp": time.Now().Unix(),
				"accountId": services.AuthGetAccountID(c),
			})
			if code == 0 {
				l.Info(message)
			} else {
				l.Warning(message)
			}
		}
	}
}
