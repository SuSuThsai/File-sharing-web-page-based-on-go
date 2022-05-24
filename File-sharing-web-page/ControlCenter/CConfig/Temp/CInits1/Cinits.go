package CInits1

import (
	"File-sharing-web-page/ControlCenter/CConfig/Temp/CGlobal1"
	"encoding/json"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
)

func CInits() {
	base := viper.New()
	base1 := viper.New()
	//文件路径
	base.SetConfigFile(CGlobal1.CGINConfigFilePath)
	base1.SetConfigFile(CGlobal1.CAliConfigFilePath)
	if err := base.ReadInConfig(); err != nil {
		log.Panic("解析配置文件失败", err.Error())
	}
	if err := base1.ReadInConfig(); err != nil {
		log.Panic("解析配置文件失败", err.Error())
	}
	if err := base1.Unmarshal(&CGlobal1.CGlobal.DriverConfig); err != nil {
		log.Panic("解析配置文件失败", err.Error())
	}
	if err := base1.Unmarshal(&CGlobal1.CGlobal.RefreshTokenConfig); err != nil {
		log.Panic("获取Token出错", err.Error())
	}
	if err := base.Unmarshal(&CGlobal1.CGlobal.GINMysql); err != nil {
		log.Panic("解析Mysql配置文件失败", err.Error())
	}
	if err := base.Unmarshal(&CGlobal1.CGlobal.GINRedis); err != nil {
		log.Panic("解析Redis配置文件失败", err.Error())
	}
	if err := base.Unmarshal(&CGlobal1.CGlobal.GetJwtKey); err != nil {
		log.Panic("获取JwtKey失败", err.Error())
	}
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
