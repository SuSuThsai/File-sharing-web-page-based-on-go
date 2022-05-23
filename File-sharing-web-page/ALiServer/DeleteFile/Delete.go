package DeleteFile

import (
	//本地
	"log"
	//第三方
	//目录
	"File-sharing-web-page/Global"
)

func DeleteFile(fileId string) (err error) {
	base, err := Global.Global.Driver.RemoveFile(Global.Global.Credential, fileId)
	if err != nil {
		log.Println("删除失败:", err.Error())
		return
	}
	log.Println(base.Code, base.Message)
	return
}
