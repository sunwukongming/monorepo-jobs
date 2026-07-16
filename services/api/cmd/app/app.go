package main

import (
	"net/http"

	"app/config"
	_ "app/init"
	"app/internal/db/mysql"
	"app/internal/db/mysql/query"
	"app/internal/middleware"
	"app/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// var configPath string
	// flag.StringVar(&configPath, "c", path.Join(utils.FileExecuteDir(), "config.yaml"), "配置文件")
	// flag.Parse()
	// configString, err := os.ReadFile(configPath)
	// if err != nil {
	// 	logrus.Fatalf("配置文件错误 %s %v", configPath, err)
	// }
	// loader.LoadConfig(string(configString))
	// bootstrap.Bootstrap(config.Get())

	query.SetDefault(mysql.Gorm())
	// 1.创建路由
	r := gin.Default()
	r.Use(middleware.Cors())
	r.Use(middleware.Common())
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})
	r.GET("/api/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 0,
			"data": gin.H{
				"check": true,
			},
		})
	})
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	apiRouter := r.Group("/api")
	router.RegisterApi(apiRouter)
	r.Run(":" + config.Get().Listen)
}
