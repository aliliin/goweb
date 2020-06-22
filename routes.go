package main

import (
	"github.com/gin-gonic/gin"
	"goweb/learngin/controller"
	"goweb/learngin/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info) // 用户信息的路由
	return r
}
