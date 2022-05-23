package GINSaveFile

import (
	//本地
	"log"
	"time"
	//第三方
	"golang.org/x/net/context"
	//目录
	"File-sharing-web-page/GINServer/Until"
	"File-sharing-web-page/Global"
)

func SaveFileInfo(fileId string) (Code string, Result bool) {
	c := context.Background()
	key := Until.GetRandom()
	dbR := Global.Global.DBR
	dbR2 := Global.Global.DBR2
	if dbR.Get(c, key).Err() != nil {
		status := dbR.Set(c, key, fileId, 24*time.Hour)
		status2 := dbR2.Set(c, key, fileId, 0)
		if status.Err() == nil || status2.Err() == nil {
			Result = true
			Code = key
		} else {
			Code = key
			log.Println("生成取件码出错：", status.Err(), status2.Err())
		}
	} else {
		Code = "-1"
	}
	return
}
