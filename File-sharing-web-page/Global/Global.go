package Global

import (
	//本地

	//第三方
	"github.com/jakeslee/aliyundrive"
	//目录
	"File-sharing-web-page/ALiServer/ALiconfig"
	"File-sharing-web-page/GINServer/UserInfo"
)

//全局变量

var ConfigFilePath = "ALiServer/ALiconfig/AliInits.yaml"

var Global struct {
	//阿里端
	Credential         *aliyundrive.Credential
	Driver             *aliyundrive.AliyunDrive
	DriverConfig       ALiconfig.DriverConfig
	RefreshTokenConfig ALiconfig.RefreshTokenConfig
	//Gin端
	User UserInfo.UserInfo
}
