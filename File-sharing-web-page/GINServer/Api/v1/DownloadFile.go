package v1

import (
	//本地
	"log"
	"net/http"
	//第三方
	"github.com/gin-gonic/gin"
	//目录
	"File-sharing-web-page/GINServer/GINModels/GINDownloadFile"
	"File-sharing-web-page/GINServer/GINconfig"
)

func Download(c *gin.Context) {
	var Code GINconfig.CODE
	err := c.ShouldBind(&Code)
	if err != nil {
		log.Fatal(err.Error())
	}
	fildId, code := GINDownloadFile.DownLoadFile(Code.FetchCode)
	if code == http.StatusConflict {
		c.JSON(http.StatusConflict, gin.H{
			"status":  http.StatusConflict,
			"FileGet": false,
			"message": "无效的取件码",
		})
		c.Abort()
		return
	} else if code == http.StatusInternalServerError {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"FileGet": false,
			"message": "验证取件码出错",
		})
		c.Abort()
		return
	}
	URL := GINDownloadFile.GetUrlDownload(fildId)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"FileGet": true,
		"message": "获取成功",
		"FileUrl": URL,
	})
}
