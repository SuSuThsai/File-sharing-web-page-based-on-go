package UpLoad

import (
	"log"
	//第三方
	"github.com/jakeslee/aliyundrive/models"
	//目录
	"File-sharing-web-page/Global"
)

func GetUrl(fileId string) (URLInfo *models.DownloadURLResponse) {
	var err error
	URLInfo, err = Global.Global.Driver.GetDownloadURL(Global.Global.Credential, fileId)
	if err != nil {
		log.Println("获取URL失败:", err.Error())
		return
	}
	return
}
