// @program: concurrency
// @author: edte
// @create: 2020-07-15 15:28
package model

import (
	"github.com/jinzhu/gorm"
)

// Order 表示订单
type Order struct {
	gorm.Model
	UserID  string
	GoodsID uint
	Num     int
}

// MakeOder 表示下单
func (order *Order) MakeOder() error {
	return DB.Create(&order).Error
}

// GetOrderByUid 通过 user id 获取相应的订单
func GetOrderByUid(uid string) (orders []Order, err error) {
	err = DB.Table("orders").Where("user_id", uid).Find(&orders).Error
	return orders, err
}
