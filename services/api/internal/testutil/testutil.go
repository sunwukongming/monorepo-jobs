// Package testutil 为 API 提供轻量集成测试脚手架：
// 用 go-sqlmock 注入一个 MySQL 方言的 mock 连接，装配真实路由，
// 便于对 handler 的响应外壳（code/message/data 信封）做回归测试。
package testutil

import (
	"testing"

	"app/config"
	mdb "app/internal/db/mysql"
	"app/internal/db/mysql/query"
	"app/internal/middleware"
	"app/internal/router"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// Setup 构建一个使用 mock DB 的测试 API：注入 mock 连接、装配 /api 路由。
// 返回 gin 引擎与 sqlmock 句柄（用于设置查询期望）。
func Setup(t *testing.T) (*gin.Engine, sqlmock.Sqlmock) {
	t.Helper()
	config.Set(&config.Config{Mode: "release", Key: "test-key", TimeZone: "Asia/Shanghai"})

	sqlDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	if err != nil {
		t.Fatalf("sqlmock.New: %v", err)
	}
	mock.MatchExpectationsInOrder(false)
	t.Cleanup(func() { _ = sqlDB.Close() })

	gdb, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{Logger: gormlogger.Default.LogMode(gormlogger.Silent)})
	if err != nil {
		t.Fatalf("gorm.Open: %v", err)
	}
	mdb.SetGorm(gdb)
	query.SetDefault(gdb)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(middleware.Cors())
	r.Use(middleware.Common())
	api := r.Group("/api")
	router.RegisterApi(api)
	return r, mock
}
