package CGlobal1

import (
	"File-sharing-web-page/ControlCenter/CConfig/Temp"
	//本地
	//第三方
	"github.com/go-redis/redis/v8"
	"github.com/jakeslee/aliyundrive"
	"gorm.io/gorm"
)

//全局变量

var CAliConfigFilePath = "ControlCenter/CConfig/Aliyun.yaml.yaml"
var CGINConfigFilePath = "ControlCenter/CConfig/GinConfig.yaml"
var ServerIP map[string]int
var WebIP map[string]int

var CGlobal struct {
	//阿里端
	Credential         *aliyundrive.Credential
	Driver             *aliyundrive.AliyunDrive
	DriverConfig       Temp.DriverConfig
	RefreshTokenConfig Temp.RefreshTokenConfig
	//Gin端
	DB        *gorm.DB
	DBR       *redis.Client
	DBR2      *redis.Client
	GINMysql  Temp.MysqlInfo
	GINRedis  Temp.RedisInfo
	GetJwtKey *Temp.JWT
}
