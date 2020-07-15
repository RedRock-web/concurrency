// @program: concurrency
// @author: edte
// @create: 2020-07-15 15:45
package controller

import (
	"github.com/gin-gonic/gin"

	"concurrency/router"
	"concurrency/service"
)

func MakeOrder(c *gin.Context) {
	var u service.User

	if err := c.ShouldBindJSON(&u); err != nil {
		router.FormError(c)
		return
	}

	service.OrderChan <- u

	router.Ok(c)
}
