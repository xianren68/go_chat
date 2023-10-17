package main

import "go_chat/initialize"

func main() {
	// 初始化日志
	initialize.InitLogger()
	// 初始化配置
	initialize.InitConfig()
	// 初始化数据库
	initialize.InitDB()
}
