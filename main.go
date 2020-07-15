// @program: concurrency
// @author: edte
// @create: 2020-07-15 15:18
package main

import (
	"concurrency/model"
	"concurrency/router"
	"concurrency/service"
)

func main() {
	model.InitDB()
	service.InitService()
	router.InitRouter()
}
