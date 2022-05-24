package CDiscovery

import (
	"File-sharing-web-page/ControlCenter/CGlobal"
	"github.com/gin-gonic/gin"
	"log"
)

func CDiscovery(c *gin.Context) {
	ip := c.ClientIP()
	host := c.Query("host")
	ip = ip + ":" + host
	if host == "" {
		c.JSON(400, gin.H{"msg": "缺少host"})
		c.Abort()
		return
	}
	CGlobal.ServerIP[ip] = 1
	log.Println("已发现", ip)
	c.JSON(200, gin.H{"IP": ip, "host": host})
}
