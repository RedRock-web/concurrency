// @program: concurrency
// @author: edte
// @create: 2020-07-15 15:49
package router

import (
	"github.com/gin-gonic/gin"

	"concurrency/controller"
)

// InitRouter 初始化 router
func InitRouter() {
	r := gin.Default()
	SetRouter(r)
	_ = r.Run("8080")
}

// SetRouter 设置 router
func SetRouter(r *gin.Engine) {
	goods := r.Group("/goods")
	{
		goods.GET("/get_all", controller.SelectAllGoods)
		goods.GET("/get", controller.SelectGoodsByGid)
		goods.POST("add", controller.AddGoods)
	}

	order := r.Group("/order")
	{
		order.POST("/order", controller.MakeOrder)
	}
}
