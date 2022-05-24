package Routers

import (
	//本地
	"log"
	"strconv"
	//第三方
	"github.com/gin-gonic/gin"
	//目录
	v1 "File-sharing-web-page/GINServer/Api/v1"
	"File-sharing-web-page/GINServer/MiddleWare"
	"File-sharing-web-page/GINServer/Until"
)

func InitUserRouter() {
	gin.SetMode("debug")
	cod := gin.New()
	cod.Use(MiddleWare.Cors())
	cod.Use(gin.Recovery())
	login := cod.Group("")
	{
		login.POST("/Login", v1.Login)
		login.POST("/adduser", v1.AddUser)
	}
	User := cod.Group("/user", MiddleWare.JwtToken())
	{
		User.POST("/UploadFile", v1.UpLoadWeb)
		User.POST("/Download", v1.Download)
	}
	Visitor := cod.Group("/visitor")
	{
		Visitor.POST("/UploadFile", v1.UpLoadWeb)
		Visitor.POST("/Download", v1.Download)
	}
	port, err := Until.GetFreePort()
	if err != nil {
		log.Panicln("获取端口失败,启用端口5053：", err.Error())
	}
	cod.Run(":" + strconv.Itoa(port))
}
