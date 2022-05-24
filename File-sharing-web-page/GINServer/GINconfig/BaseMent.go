package GINconfig

import (
	//本地
	//第三方
	"gorm.io/gorm"
	//目录
)

type UserInfo struct {
	gorm.Model
	Username string ` gorm:"type:varchar(20);not null" validate:"required,min=2,max=12" json:"username"`
	Password string `gorm:"type:varchar(20);not null" validate:"required,min=6,max=20" json:"password"`
}

//func (u *UserInfo)BeforeCreate(_ *gorm.DB) (err error) {
//	u.Password = User.ScryptPW(u.Password)
//	return nil
//}
//
//func (u *UserInfo)BeforeUpdate(_ *gorm.DB) (err error) {
//	u.Password = User.ScryptPW(u.Password)
//	return nil
//}

type MysqlInfo struct {
	Db         string `mapstructure:"db" json:"db"`
	Host       string `mapstructure:"host" json:"host"`
	Port       int    `mapstructure:"port" json:"port"`
	DbUser     string `mapstructure:"dbUser" json:"db_user" gorm:"DEFAULT:root"`
	DbPassword string `mapstructure:"dbPassword" json:"db_password"`
	DbName     string `mapstructure:"dbName" json:"db_name"`
}

type RedisInfo struct {
	DbR         string `mapstructure:"dbr" json:"db_r"`
	HostR       string `mapstructure:"hostr" json:"host_r"`
	PortR       int    `mapstructure:"portr" json:"port_r"`
	DbUserR     string `mapstructure:"dbuserr" json:"dbuser_r" gorm:"DEFAULT:defalut"`
	DbPassWordR string `mapstructure:"dbpasswordr" json:"db_pass_word_r"`
	DBModel     int    `mapstructure:"dbModel" json:"db_model"`
}

type ControlCenter struct {
	HostC string `mapstructure:"HostC" json:"host_c"`
	PortC int    `mapstructure:"PortC" json:"port_c"`
}

type JWT struct {
	JwtKey string `mapstructure:"jwtKey" json:"jwt_key"`
}

type CODE struct {
	FetchCode string `json:"fetch_code" validate:"required,min=6,max=6"`
}
