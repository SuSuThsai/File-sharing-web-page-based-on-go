package v1

import (
	//本地
	"net/http"
	//第三方
	"github.com/gin-gonic/gin"
	//目录
	"File-sharing-web-page/GINServer/GINModels/GINSaveFile"
	"File-sharing-web-page/GINServer/GINModels/UploadFileWeb"
	"File-sharing-web-page/GINServer/Until/ERRTell"
)

func UpLoadWeb(c *gin.Context) {
	file, _ := c.FormFile("file")
	FileId, status := UploadFileWeb.UpLoadFile(file)
	if status == 0 {
		key, flag := GINSaveFile.SaveFileInfo(FileId)
		if flag {
			c.JSON(http.StatusOK, gin.H{
				"status":    http.StatusOK,
				"CodeCreat": true,
				"message":   "上传文件成功",
				"FileCode":  key,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":    http.StatusInternalServerError,
				"CodeCreat": nil,
				"message":   "保存出错请重试",
				"FileCode":  key,
			})
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":    status,
			"CodeCreat": false,
			"message":   ERRTell.GetErrMsg(int(status)),
			"FileId":    FileId,
		})
	}
}
