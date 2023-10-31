// Package global 全局的各种配置与连接池
package global

import (
	"github.com/go-redis/redis/v8"
	"go_chat/config"
	"gorm.io/gorm"
)

var (
	// DB mysql数据库
	DB *gorm.DB

	// RedisDB redis连接
	RedisDB       *redis.Client
	ServiceConfig *config.ServiceConfig
)
