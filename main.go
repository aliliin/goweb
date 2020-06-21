package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"goweb/learngin/common"
	"os"
)

func main() {
	InitConfig()
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()
	r = CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	// 监听并在 0.0.0.0:8080 上启动服务
	panic(r.Run())
}

// 初始化配置
func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("env")               // 读取的文件名字
	viper.SetConfigType("yml")               // 读取的文件类型
	viper.AddConfigPath(workDir + "/config") // 文件路径
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
