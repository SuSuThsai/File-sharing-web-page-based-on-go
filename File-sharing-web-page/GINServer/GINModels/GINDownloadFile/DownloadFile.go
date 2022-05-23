package GINDownloadFile

import (
	//本地
	"net/http"
	//第三方
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	//目录
	"File-sharing-web-page/ALiServer/UpLoad"
	"File-sharing-web-page/Global"
)

func DownLoadFile(FetchCode string) (fileId string, code int) {
	cmdR := Global.Global.DBR
	c := context.Background()
	value, err := cmdR.Get(c, FetchCode).Result()
	fileId = value
	if err != nil && err == redis.Nil {
		code = http.StatusConflict
		return
	} else if err != nil {
		code = http.StatusInternalServerError
		return
	} else {
		code = http.StatusOK
		return
	}
}

func GetUrlDownload(fileId string) (URL string) {
	URLInfo := UpLoad.GetUrl(fileId)
	URL = *URLInfo.Url
	return
}
