package initialize

import (
	"go_chat/config"
	"go_chat/global"
	"gopkg.in/ini.v1"
)

// InitConfig 初始化配置
func InitConfig() {
	// 获取配置文件
	cfg, err := ini.Load("config.ini")
	if err != nil {
		panic("Fail to read config")
	}
	// 创建服务器配置
	serviceCfg := &config.ServiceConfig{}
	serviceCfg.Port = cfg.Section("").Key("port").String()
	serviceCfg.Mysql = initMysql(cfg)
	serviceCfg.Redis = initRedis(cfg)
	global.ServiceConfig = serviceCfg
}

// 初始化mysql配置
func initMysql(cfg *ini.File) *config.MysqlConfig {
	mysqlConfig := &config.MysqlConfig{}
	mysqlConfig.User = cfg.Section("mysql").Key("user").MustString("root")
	mysqlConfig.Password = cfg.Section("mysql").Key("password").String()
	mysqlConfig.Host = cfg.Section("mysql").Key("host").String()
	mysqlConfig.Port = cfg.Section("mysql").Key("port").String()
	mysqlConfig.DbName = cfg.Section("mysql").Key("dbname").String()
	return mysqlConfig
}

// 初始化redis配置
func initRedis(cfg *ini.File) *config.RedisConfig {
	redisConfig := &config.RedisConfig{}
	redisConfig.Host = cfg.Section("redis").Key("host").String()
	redisConfig.Port = cfg.Section("redis").Key("port").String()
	return redisConfig
}
