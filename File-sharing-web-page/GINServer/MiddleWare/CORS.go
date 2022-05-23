package MiddleWare

import (
	//本地
	"time"
	//第三方
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	//目录
)

func Cors() gin.HandlerFunc {
	return cors.New(
		cors.Config{
			AllowOrigins:     []string{"*"}, // 等同于允许所有域名
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"*", "Authorization"},
			ExposeHeaders:    []string{"Content-Length", "text/plain", "Authorization", "Content-Type", "Referer", "Host"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		})
}
