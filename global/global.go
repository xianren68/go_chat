// Package global 全局的各种配置与连接池
package global

import (
	"go_chat/config"
	"gorm.io/gorm"
)

var (
	// DB mysql数据库
	DB            *gorm.DB
	ServiceConfig *config.ServiceConfig
)
