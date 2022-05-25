package LoadBalance

import (
	"File-sharing-web-page/GINServer/MiddleWare"
	"github.com/gin-gonic/gin"
)

func InitUserRouter() {
	gin.SetMode("debug")
	cod := gin.New()
	cod.Use(MiddleWare.Cors())
	cod.Use(gin.Recovery())
	x := cod.Group("")
	x.GET("")

	//ginRegusture.GinRegusture(port)
	cod.Run(":8999")
}
