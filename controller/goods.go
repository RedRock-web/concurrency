// @program: concurrency
// @author: edte
// @create: 2020-07-15 15:45
package controller

import (
	"github.com/gin-gonic/gin"

	"concurrency/router"
	"concurrency/service"
)

// AddGoods 增加商品
func AddGoods(c *gin.Context) {
	var g service.Goods

	if err := c.ShouldBindJSON(&g); err != nil {
		router.FormError(c)
		return
	}

	service.AddGoods(g)
	router.Ok(c)
}

func SelectAllGoods(c *gin.Context) {
	goods := service.SelectAllGoods()
	router.OkWithData(c, goods)
}

func SelectGoodsByGid(c *gin.Context) {

}
