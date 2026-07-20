package services

import "github.com/gin-gonic/gin"

// Handle 收敛控制器的通用响应样板：运行 fn，成功则以其返回值走 ResponseSuccess，
// 出错则统一走 ResponseError(-1)。替代各控制器里
// `err := func() error {...}()` + 尾部 if/ResponseError/ResponseSuccess 的重复结构。
//
// 用法：
//
//	func XAction(c *gin.Context) {
//	    services.Handle(c, func() (interface{}, error) {
//	        ...
//	        return resp, nil
//	    })
//	}
func Handle(c *gin.Context, fn func() (interface{}, error)) {
	data, err := fn()
	if err != nil {
		ResponseError(c, -1, err.Error(), nil)
		return
	}
	ResponseSuccess(c, data)
}
