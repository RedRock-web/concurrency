// @program: concurrency
// @author: edte
// @create: 2020-07-15 15:46
package service

import (
	"log"

	"concurrency/model"
)

type Goods struct {
	GID   int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Num   int    `json:"num"`
}

func AddGoods(g Goods) {
	goods := model.Goods{
		Gid:   g.GID,
		Name:  g.Name,
		Price: g.Price,
		Num:   g.Num,
	}
	err := goods.AddGoods()
	if err != nil {
		log.Printf("add goods failed:%s", err)
	}
}

func SelectAllGoods() (goods []Goods) {
	mgoods, err := model.SelectAllGoods()
	if err != nil {
		log.Print(err)
		return
	}

	for _, v := range mgoods {
		good := Goods{
			GID:   v.Gid,
			Name:  v.Name,
			Price: v.Price,
			Num:   v.Num,
		}
		goods = append(goods, good)
	}
	return goods
}

func SelectGoodsByGid() {

}
