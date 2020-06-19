package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"goweb/learngin/common"
)

func main() {

	db := common.InitDB()
	defer db.Close()

	r := gin.Default()
	r = CollectRoute(r)

	panic(r.Run()) // 监听并在 0.0.0.0:8080 上启动服务
}
