package UploadFileWeb

import (
	//本地
	"mime/multipart"
	//第三方
	//目录
	"File-sharing-web-page/ALiServer/UpLoad"
)

func UpLoadFile(file *multipart.FileHeader) (FileId string, status int64) {
	FileId = UpLoad.UpLoad(file)
	if FileId == "" {
		status = 5001
	}
	return
}
