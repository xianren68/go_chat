// Package initialize 项目初始化配置
package initialize

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"go_chat/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// InitDB 初始化mysql连接
func InitDB() {
	// 设置数据库连接信息
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.ServiceConfig.Mysql.User,
		global.ServiceConfig.Mysql.Password,
		global.ServiceConfig.Mysql.Host,
		global.ServiceConfig.Mysql.Port,
		global.ServiceConfig.Mysql.DbName)

	// sql日志配置
	newLogger := logger.New(
		// 日志打印的位置
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		})
	var err error
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic(err)
	}
}

// InitRedis 初始化redis连接
func InitRedis() {
	opt := &redis.Options{
		Addr:     global.ServiceConfig.Redis.Host + ":" + global.ServiceConfig.Redis.Port,
		Password: "",
	}
	global.RedisDB = redis.NewClient(opt)
}
