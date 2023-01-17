package main

import (
	"github.com/BurntSushi/toml"
)

var Configer *Config

type MyConfig struct {
	User     string `json:"user"`
	Host     string `json:"host"`
	Password string `json:"password"`
	DbName   string `json:"dbname"`
}

type RedConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}

type Config struct {
	RedMysql MyConfig  `toml:"red_mysql"`
	Redis    RedConfig `toml:"redis"`
}

// Init 初始化配置文件
func init() {
	Configer = &Config{}
	_, err := toml.DecodeFile("config.toml", Configer)
	if err != nil {
		panic("读取配置文件失败!,原因:" + err.Error())
	}
	//redis
	GenRedis()
	//mysql 启动
	//MysqlInit()
}
