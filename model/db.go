// @program: concurrency
// @author: edte
// @create: 2020-07-15 15:20
package model

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

// InitDB 初始化 database
func InitDB() {
	db, err := gorm.Open("mysql", "mysql", "root:mima@tcp(127.0.0."+
		"1:3306)/concurrency?parseTime=true&charset=utf8&loc=Local")
	if err != nil {
		log.Panicf("Panic while connecting the gorm. Error: %s", err)
	}

	DB = db

	if DB.HasTable(&Goods{}) {
		DB.AutoMigrate(&Goods{})
	} else {
		DB.CreateTable(&Goods{})
	}

	if DB.HasTable(&Order{}) {
		DB.AutoMigrate(&Order{})
	} else {
		DB.CreateTable(&Order{})
	}
}
