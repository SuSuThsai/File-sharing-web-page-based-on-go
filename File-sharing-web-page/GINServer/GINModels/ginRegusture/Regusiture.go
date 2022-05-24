package ginRegusture

import (
	"File-sharing-web-page/Global"
	"log"
	"net/http"
	"strconv"
)

func GinRegusture(port int) {
	resp, err := http.Get("http://" + Global.Global.ControlCenter.HostC + ":" + strconv.Itoa(Global.Global.ControlCenter.PortC) + "/ADD_GINSERVER" + "?host=" + strconv.Itoa(port))
	if err != nil {
		log.Panic("注册出错：", err.Error())
	}
	if resp.StatusCode == http.StatusOK {
		log.Println("服务注册成功")
	} else {
		log.Fatal("服务注册失败")
	}
}
