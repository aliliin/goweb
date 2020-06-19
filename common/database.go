package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"goweb/learngin/model"
)

var DB *gorm.DB

// 开启链接池
func InitDB() *gorm.DB {

	driverName := "mysql"
	host := "127.0.0.1"
	port := "3306"
	database := "goweb"
	username := "root"
	password := ""
	charset := "utf8"

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)

	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	// 自动创建表
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
