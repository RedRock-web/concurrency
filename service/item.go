// @program: concurrency
// @author: edte
// @create: 2020-07-15 16:12
package service

import (
	"log"
	"sync"
	"time"

	"github.com/robfig/cron"

	"concurrency/model"
)

type User struct {
	UserId  string
	GoodsId uint
}

var OrderChan = make(chan User, 1024)

//  上架一次的项目
type Item struct {
	ID        uint   // 商品id
	Name      string // 名字
	Total     int    // 商品总量
	Left      int    // 商品剩余数量
	IsSoldOut bool   // 是否售罄
	leftCh    chan int
	sellCh    chan int
	done      chan struct{}
	Lock      sync.Mutex
}

// Item map
var ItemMap = make(map[uint]*Item)

// initMap() 每天凌晨执行 map 初始化任务
func initMap() {
	c := cron.New()
	if err := c.AddFunc("0 0 1 * * ?", GetAllGoods); err != nil {
		log.Printf("failed to execute sync task:%s", err)
		return
	}
}

// GetAllGoods 把数据库中的数据同步到 map 中
func GetAllGoods() {
	goods, err := model.SelectAllGoods()
	if err != nil {
		log.Printf("failed to get all goods:%s", err)
		return
	}

	for _, good := range goods {
		item := Item{
			ID:        good.ID,
			Name:      good.Name,
			Total:     good.Num,
			Left:      good.Num,
			IsSoldOut: false,
			leftCh:    make(chan int),
			sellCh:    make(chan int),
			done:      make(chan struct{}),
		}

		ItemMap[item.ID] = &item
	}
}

func getItem(itemId uint) *Item {
	return ItemMap[itemId]
}

func order() {
	// 遍历处理订单
	for {
		// 响应给用户，表明订单处理完成，但是这里并没有传商品购买成功与否的信息
		// 需要更改一下或用户处理那里轮讯

		user := <-OrderChan
		// 获取对应商品服务
		item := getItem(user.GoodsId)
		//
		item.SecKilling(user.UserId)
	}
}

func (item *Item) SecKilling(userId string) {

	// 并发加锁
	item.Lock.Lock()
	defer item.Lock.Unlock()
	// 等价
	// var lock = make(chan struct{}, 1}
	// lock <- struct{}{}
	// defer func() {
	// 		<- lock
	// }
	if item.IsSoldOut {
		return
	}

	item.BuyGoods(1)

	MakeOrder(userId, item.ID, 1)
}

// 定时下架
func (item *Item) OffShelve() {
	beginTime := time.Now()
	// 获取第二天时间
	// nextTime := beginTime.Add(time.Hour * 24)
	// 计算次日零点，即商品下架的时间
	// offShelveTime := time.Date(nextTime.Year(), nextTime.Month(), nextTime.Day(), 0, 0, 0, 0, nextTime.Location())
	offShelveTime := beginTime.Add(time.Second * 5)
	timer := time.NewTimer(offShelveTime.Sub(beginTime))

	// 阻塞一直等待到达下架时间
	<-timer.C
	delete(ItemMap, item.ID)
	close(item.done)
}

// 出售商品
func (item *Item) SalesGoods() {
	// 一直遍历出售商品，直到商品售完
	for {
		select {
		// 获取购买的数据
		case num := <-item.sellCh:
			if item.Left -= num; item.Left <= 0 {
				item.IsSoldOut = true
			}

		// 商品剩余的数目小于用户要购买的数目
		case item.leftCh <- item.Left:

		// 出售商品的服务下架,如果下架了，直接结束商品出售服务
		case <-item.Done():
			return
		}
	}
}

// Done 用于得知商品售出服务的情况
func (item *Item) Done() <-chan struct{} {
	return item.done
}

// Monitor
func (item *Item) Monitor() {
	go item.SalesGoods()
}

// 获取剩余库存
func (item *Item) GetLeft() int {
	var left int
	left = <-item.leftCh
	return left
}

// 购买商品
func (item *Item) BuyGoods(num int) {
	item.sellCh <- num
}

// InitService 初始化服务
func InitService() {
	// 初始化商品信息:数据库层加载到内存中
	initMap()

	for _, item := range ItemMap {
		item.Monitor()
		go item.OffShelve()
	}

	for i := 0; i < 10; i++ {
		go order()
	}
}
