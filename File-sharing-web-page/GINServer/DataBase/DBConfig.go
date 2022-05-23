package DataBase

import (
	"File-sharing-web-page/GINServer/DataBase/Mysql"
	"File-sharing-web-page/GINServer/DataBase/Redis"
)

func Inits() {
	Mysql.InitsMysql()
	Redis.InitsRedis()
}
