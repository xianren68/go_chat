package main

import (
	"go_chat/global"
	"go_chat/initialize"
	"go_chat/router"
	"strings"
)

func main() {
	// 初始化日志
	initialize.InitLogger()
	// 初始化配置
	initialize.InitConfig()
	// 初始化数据库
	initialize.InitDB()
	// 初始化redis
	initialize.InitRedis()
	// 启动端口监听
	r := router.Router()
	err := r.Run(strings.Join([]string{":", global.ServiceConfig.Port}, ""))
	if err != nil {
		panic(err)
	}
}
