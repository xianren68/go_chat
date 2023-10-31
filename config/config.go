// Package config 配置信息
package config

// MysqlConfig mysql数据库配置
type MysqlConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DbName   string
}

// ServiceConfig 服务器相关配置
type ServiceConfig struct {
	Port  string // 服务器开启的端口
	Mysql *MysqlConfig
	Redis *RedisConfig
}

// RedisConfig redis数据库相关配置
type RedisConfig struct {
	Host string
	Port string
}
