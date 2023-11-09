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
	Port        string // 服务器开启的端口
	EmailCode   string // 邮箱授权码
	Email       string // 邮箱账号
	Mysql       *MysqlConfig
	Redis       *RedisConfig
	QiNiuConfig *QiNiuConfig
}

// RedisConfig redis数据库相关配置
type RedisConfig struct {
	Host string
	Port string
}

// QiNiuConfig 七牛云配置
type QiNiuConfig struct {
	Zone        int
	AccessKey   string
	SecretKey   string
	Bucket      string
	QiniuServer string
}
