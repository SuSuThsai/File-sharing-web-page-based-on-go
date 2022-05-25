package HeathyCheck

import (
	"File-sharing-web-page/ControlCenter/CGlobal"
	"fmt"
	"log"
	"net/http"
)

func HeathyCheck() {
	c := http.Client{}

	for s, _ := range CGlobal.ServerIP {
		r, err := http.NewRequest("GET", "http://"+s+"/HeathyCeak", nil)
		if err != nil {
			CGlobal.ServerIP[s]++
		}
		re, err := c.Do(r)
		if err != nil || re.StatusCode != http.StatusOK {
			CGlobal.ServerIP[s]++
		}
		if re.StatusCode == http.StatusOK {
			log.Println("服务器ip：", s, "存活")
		}
		if CGlobal.ServerIP[s] >= 10 {
			fmt.Println("服务器挂掉了IP：", s)
			delete(CGlobal.ServerIP, s)
		}
	}
}
