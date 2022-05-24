package CRouters

import (
	"File-sharing-web-page/ControlCenter/CConfig"
	"File-sharing-web-page/ControlCenter/CDiscovery"
	"File-sharing-web-page/ControlCenter/CGlobal"
	"File-sharing-web-page/ControlCenter/HeathyCheck"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func IniteRouter() {
	gin.SetMode("debug")
	cod := gin.New()
	cod.Use(gin.Recovery())

	go func() {
		for {
			CConfig.SaveIpfile()
			time.Sleep(time.Second * 60)
		}
	}()

	go func() {
		for {
			HeathyCheck.HeathyCheck()
			time.Sleep(time.Second * 30)
		}
	}()

	cod.GET("/ADD_GINSERVER", CDiscovery.CDiscovery)
	cod.Run(":" + strconv.Itoa(CGlobal.GinPot.Port))
}
