package ALiInits

import (
	//本地
	"log"
	//第三方
	"github.com/jakeslee/aliyundrive"
	"github.com/spf13/viper"
	//目录
	"File-sharing-web-page/Global"
)

func InitsConfig() {
	base := viper.New()
	//文件路径
	base.SetConfigFile(Global.AliConfigFilePath)
	if err := base.ReadInConfig(); err != nil {
		log.Panic("解析配置文件失败", err.Error())
	}
	if err := base.Unmarshal(&Global.Global.DriverConfig); err != nil {
		log.Panic("解析配置文件失败", err.Error())
	}
	if err := base.Unmarshal(&Global.Global.RefreshTokenConfig); err != nil {
		log.Panic("获取Token出错", err.Error())
	}
}

func InitsDriver() {
	Global.Global.Driver = aliyundrive.NewClient(&aliyundrive.Options{
		AutoRefresh:     Global.Global.DriverConfig.AutoRefresh,
		UploadRate:      Global.Global.DriverConfig.UploadRate * 1024 * 1024, // 限速 8MBps
		RefreshDuration: Global.Global.DriverConfig.RefreshDuration,
	})
	var err error
	Global.Global.Credential, err = Global.Global.Driver.AddCredential(aliyundrive.NewCredential(&aliyundrive.Credential{
		RefreshToken: Global.Global.RefreshTokenConfig.RefreshToken,
	}))

	if err != nil {
		log.Panic("获取个人信息失败", err.Error())
	}
	log.Println("阿里云主人ID为：", Global.Global.Credential.UserId)
}
