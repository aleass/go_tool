package main

import (
	"database/sql"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var (
	Mssql *gorm.DB
	Db    *sql.DB
)

func init() {
	getMysql()
}

/**
 * 获取 mysql 客户端
 */
func getMysql() {
	var err error
	config, err := os.ReadFile("dsn") //数据库 dsn  server=xxxxx;user id=xx;password=xx;port=1433;database=xx;encrypt=disable
	dsn := string(config)
	w := log.New(os.Stdout, "\r\n", log.LstdFlags) // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
	newLogger := logger.New(
		w,
		logger.Config{
			LogLevel:                  logger.Silent, // 日志级别
			IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,         // 禁用彩色打印
		},
	)
	Mssql, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("mysql 启动失败!,原因:" + err.Error())
	}
	Db, err = Mssql.DB()
	if err != nil {
		panic("mysql 启动失败!,原因:" + err.Error())
	}
	err = Db.Ping()
	if err != nil {
		panic("mysql 启动失败!,原因:" + err.Error())
	}
}
