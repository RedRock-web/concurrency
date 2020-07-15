// @program: concurrency
// @author: edte
// @create: 2020-07-15 15:49
package router

import (
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化 router
func InitRouter() {
	r := gin.Default()
	SetRouter(r)
	_ = r.Run("8080")
}

// SetRouter 设置 router
func SetRouter(r *gin.Engine) {
	r.GET("/getGoods", controller.SelectGoods)
	r.POST("/order", controller.MakeOrder)
}
