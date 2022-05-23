package Mysql

import (
	//本地
	"fmt"
	"log"
	"os"
	"time"
	//第三方
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	//目录
	"File-sharing-web-page/Global"
)

var err error

func InitsMysql() {
	inits := Global.Global.GINMysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		inits.DbUser, inits.DbPassword, inits.Host, inits.Port, inits.DbName)
	log.Println("数据库信息：", dsn)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 禁用彩色打印
		},
	)
	Global.Global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// gorm日志模式
		Logger: newLogger,
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal("连接数据库失败，请检查参数：", err)
	}
	//Global.Global.DB.AutoMigrate(&GINconfig.UserInfo{})
}
