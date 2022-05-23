package Redis

import (
	//本地
	"fmt"
	"log"

	//第三方
	"github.com/go-redis/redis/v8"
	//目录
	"File-sharing-web-page/Global"
)

func InitsRedis() {
	configR := Global.Global.GINRedis
	//fmt.Println(configR,configR2)
	Global.Global.DBR = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", configR.HostR, configR.PortR),
		Password: configR.DbPassWordR, // no password set
		DB:       configR.DBModel,     // use default DB
	})
	Global.Global.DBR2 = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", configR.HostR, configR.PortR),
		Password: configR.DbPassWordR, // no password set
		DB:       configR.DBModel + 1, // use default DB
	})
	log.Println("数据库REDIS信息：", Global.Global.DBR, Global.Global.DBR2)
}
