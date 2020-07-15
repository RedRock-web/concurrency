// @program: concurrency
// @author: edte
// @create: 2020-07-15 15:46
package service

import (
	"log"

	"concurrency/model"
)

func MakeOrder(userId string, goodsId uint, num int) {
	order := model.Order{
		UserID:  userId,
		GoodsID: goodsId,
		Num:     num,
	}

	err := order.MakeOder()
	if err != nil {
		log.Printf("Error make an order. Error: %s", err)
	}
}
