package GINStatusFile

import (
	//本地
	"log"
	//第三方
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	//目录
	"File-sharing-web-page/Global"
)

func StatusFile(FetchCode string) (fileId string) {
	cmdR2 := Global.Global.DBR2
	c := context.Background()
	value, err := cmdR2.Get(c, FetchCode).Result()
	_, err2 := cmdR2.Del(c, FetchCode).Result()
	if err2 != nil {
		log.Println("静态库删除出错", err2.Error())
	}
	fileId = value
	if err != nil && err == redis.Nil {
		return
	} else if err != nil {
		log.Println("查询键值出错:", err.Error())
		return
	}
	return
}
