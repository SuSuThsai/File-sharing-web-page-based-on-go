package Global

import (
	//本地
	//第三方
	"github.com/go-redis/redis/v8"
	"github.com/jakeslee/aliyundrive"
	"gorm.io/gorm"
	//目录
	"File-sharing-web-page/ALiServer/ALiconfig"
	"File-sharing-web-page/GINServer/GINconfig"
)

//全局变量

var AliConfigFilePath = "ALiServer/ALiconfig/AliInits.yaml"
var GINConfigFilePath = "GINServer/GINconfig/GinInits.yaml"

var Global struct {
	//阿里端
	Credential         *aliyundrive.Credential
	Driver             *aliyundrive.AliyunDrive
	DriverConfig       ALiconfig.DriverConfig
	RefreshTokenConfig ALiconfig.RefreshTokenConfig
	//Gin端
	User          GINconfig.UserInfo
	DB            *gorm.DB
	DBR           *redis.Client
	DBR2          *redis.Client
	GINMysql      GINconfig.MysqlInfo
	GINRedis      GINconfig.RedisInfo
	GINRedis2     GINconfig.RedisInfo
	ControlCenter GINconfig.ControlCenter
}
