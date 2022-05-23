package KEYPassTime

import (
	//本地
	"fmt"
	"log"
	"time"
	"unsafe"
	//第三方
	"github.com/gomodule/redigo/redis"
	//目录
	"File-sharing-web-page/ALiServer/DeleteFile"
	"File-sharing-web-page/GINServer/GINModels/GINStatusFile"
	"File-sharing-web-page/Global"
)

type PSubscribeCallback func(pattern, channel, message string)

type PSubscriber struct {
	client redis.PubSubConn
	cbMap  map[string]PSubscribeCallback
}

func (c *PSubscriber) PConnect() {
	configR := Global.Global.GINRedis
	conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", configR.HostR, configR.PortR), redis.DialUsername(configR.DbUserR), redis.DialPassword(configR.DbPassWordR))
	if err != nil {
		log.Println("redis监听器连接失败:", err)
	}

	c.client = redis.PubSubConn{Conn: conn}
	c.cbMap = make(map[string]PSubscribeCallback)
	go func() {
		for {
			switch res := c.client.Receive().(type) {
			case redis.Message:
				pattern := &res.Pattern
				channel := &res.Channel
				message := (*string)(unsafe.Pointer(&res.Data))
				c.cbMap[*channel](*pattern, *channel, *message)
				fileId := GINStatusFile.StatusFile(*message)
				err = DeleteFile.DeleteFile(fileId)
				if err == nil {
					log.Println("删除成功")
				} else {
					log.Println("删除失败", err.Error())
				}
			case redis.Subscription:
				log.Printf("监控数据库以及形式 : %s: Channel : %s 键值数为 :%d\n", res.Channel, res.Kind, res.Count)
			case error:
				log.Println("error handle...", res)
				continue
			}
		}
	}()
}
func (c *PSubscriber) Psubscribe(channel interface{}, cb PSubscribeCallback) {
	err := c.client.PSubscribe(channel)
	if err != nil {
		log.Println("键值监控出错:", err.Error())
	}

	c.cbMap[channel.(string)] = cb
}

func TestPubCallback(patter, chann, msg string) {
	log.Println("监控数据库以及形式 : "+patter+" channel : ", chann, " 键值为 : ", msg)
}

func SPYOn() {
	var psub PSubscriber
	psub.PConnect()
	psub.Psubscribe("__keyevent@0__:expired", TestPubCallback)
	// 还可以是： `__keyspace@0__:cool`
	for {
		time.Sleep(1 * time.Second)
	}
}
