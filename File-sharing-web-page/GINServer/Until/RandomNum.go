package Until

import (
	//本地
	"fmt"
	"math/rand"
	"strconv"
	"time"
	//第三方
	//目录
)

func GetRandom() (ans string) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 6; i++ {
		x := rand.Intn(36)
		if x < 10 {
			ans += strconv.Itoa(x)
		} else {
			ans += string(byte(x + 55))
		}
	}
	fmt.Println("取件码为：", ans)
	return
}
