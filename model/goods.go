// @program: concurrency
// @author: edte
// @create: 2020-07-15 15:27
package model

import (
	"github.com/jinzhu/gorm"
)

// Goods 表示商品
type Goods struct {
	gorm.Model
	Gid   int
	Name  string
	Price int
	Num   int
}

// AddGoods 增加商品
func (goods Goods) AddGoods() error {
	return DB.Create(goods).Error
}

// SelectGoodsById 通过商品 id 查找商品
func SelectGoodsById(id uint) (goods Goods, err error) {
	err = DB.Table("goods").Where("id = ?", id).First(&goods).Error
	return goods, err
}

// SelectAllGoods 查找所有商品
func SelectAllGoods() (goods []Goods, err error) {
	err = DB.Table("goods").Find(&goods).Error
	return goods, err
}
