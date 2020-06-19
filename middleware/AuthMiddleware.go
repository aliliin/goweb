package middleware

import (
	"github.com/gin-gonic/gin"
	"goweb/learngin/common"
	"goweb/learngin/model"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 获取authoriztion header
		tokenString := context.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			context.JSON(http.StatusUnauthorized, gin.H{"code ": 401, "msg": "Authorization use Bearer Token"})
			context.Abort() // 将这次请求抛弃掉
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{"code ": 401, "msg": "Authorization validation failed"})
			context.Abort() // 将这次请求抛弃掉
			return
		}
		// 通过验证 获取 claim 中的 用户 ID

		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		if user.ID == 0 {
			context.JSON(http.StatusUnauthorized, gin.H{"code ": 401, "msg": "权限不足"})
			context.Abort() // 将这次请求抛弃掉
			return
		}

		// 用户存在 将 user 信息写入上下文
		context.Set("user", user)
		context.Next()
	}
}
