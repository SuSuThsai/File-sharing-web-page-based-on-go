package ALiInits

import (
	"File-sharing-web-page/Global"
	//本地
	"fmt"
	"log"
	//第三方
	"github.com/jakeslee/aliyundrive"
	"github.com/spf13/viper"
)

func InitsConfig() {
	base := viper.New()
	//文件路径
	base.SetConfigFile(Global.ConfigFilePath)
	if err := base.ReadInConfig(); err != nil {
		fmt.Println("配置文件路径出错", err)
		log.Panic(err)
	}
	if err := base.Unmarshal(&Global.Global.DriverConfig); err != nil {
		fmt.Println("解析配置文件失败", err.Error())
		log.Panic(err)
	}
	if err := base.Unmarshal(&Global.Global.RefreshTokenConfig); err != nil {
		fmt.Println("获取Token出错", err.Error())
		log.Panic(err)
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
	fmt.Println("阿里云主人ID为：", Global.Global.Credential.UserId)
}

func UploadType(x string) string {
	return "/文件"
}
