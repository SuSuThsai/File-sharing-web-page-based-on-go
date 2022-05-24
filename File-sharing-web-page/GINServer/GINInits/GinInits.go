package GINInits

import (
	"fmt"
	"log"
	//第三方
	"github.com/spf13/viper"
	//目录
	"File-sharing-web-page/GINServer/MiddleWare"
	"File-sharing-web-page/Global"
)

func InitsConfig() {
	base := viper.New()
	//文件路径
	base.SetConfigFile(Global.GINConfigFilePath)
	if err := base.ReadInConfig(); err != nil {
		log.Panic("配置文件路径出错", err.Error())
	}
	if err := base.Unmarshal(&Global.Global.GINMysql); err != nil {
		log.Panic("解析Mysql配置文件失败", err.Error())
	}
	if err := base.Unmarshal(&Global.Global.GINRedis); err != nil {
		log.Panic("解析Redis配置文件失败", err.Error())
	}
	if err := base.Unmarshal(&Global.Global.GINRedis2); err != nil {
		log.Panic("解析Redis配置文件失败", err.Error())
	}
	if err := base.Unmarshal(&MiddleWare.GetJwtKey); err != nil {
		log.Panic("获取JwtKey失败", err.Error())
	}
	if err := base.Unmarshal(&Global.Global.ControlCenter); err != nil {
		log.Panic("获取控制中心配置失败", err.Error())
	}
	fmt.Println(Global.Global.ControlCenter.PortC)

	fmt.Println("1")
}
