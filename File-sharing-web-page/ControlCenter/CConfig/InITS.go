package CConfig

import (
	"File-sharing-web-page/ControlCenter/CConfig/Temp/CGlobal1"
	"File-sharing-web-page/ControlCenter/CGlobal"
	"encoding/json"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
)

func CInits() {
	base := viper.New()
	//文件路径
	base.SetConfigFile(CGlobal.FilePath)
	if err := base.ReadInConfig(); err != nil {
		log.Panic("解析配置文件失败", err.Error())
	}
	if err := base.Unmarshal(&CGlobal.GinPot); err != nil {
		log.Panic("解析配置文件失败", err.Error())
	}
	CGlobal.ServerIP = map[string]int{}
	CreatIpfile()
	//c.JSON(http.StatusOK,gin.H{
	//	"DriverConfig": CGlobal1.CGlobal1.DriverConfig,
	//	"RefreshTokenConfig":CGlobal1.CGlobal1.RefreshTokenConfig,
	//	"GINMysql":&CGlobal1.CGlobal1.GINMysql,
	//	"CGlobal1.GINRedis":CGlobal1.CGlobal1.GINRedis,
	//	"GetJwtKey":CGlobal1.CGlobal1.GetJwtKey,
	//})
}

func CreatIpfile() {
	CGlobal1.ServerIP = map[string]int{}
	Ipfile, err := os.Create("IP+Host.txt")
	if err != nil {
		log.Fatal("创建储存Ip文件失败：", err.Error())
	}
	body, _ := ioutil.ReadAll(Ipfile)
	json.Unmarshal(body, &CGlobal1.ServerIP)
}

func SaveIpfile() {
	Ipfile, err := os.Create("IP+Host.txt")
	x, err := json.Marshal(&CGlobal1.ServerIP)
	if err != nil {
		log.Fatal("储存Ip失败：", err.Error())
	}
	Ipfile.Write(x)
	defer Ipfile.Close()
}
