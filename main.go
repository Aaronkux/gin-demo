package main

import (
	"gandi.icu/demo/core"
	"gandi.icu/demo/global"
	"gandi.icu/demo/initialize"
)

func main() {
	global.AM_VP = core.Viper()      // 初始化Viper
	global.AM_LOG = core.Zap()       // 初始化zap日志库
	global.AM_DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	initialize.DBList()
	initialize.SnowflakeGenerate()
	if global.AM_DB != nil {
		// initialize.RegisterTables(global.AM_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.AM_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
