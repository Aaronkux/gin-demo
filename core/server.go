package core

import (
	"fmt"
	"time"

	"gandi.icu/demo/global"
	"gandi.icu/demo/initialize"
	"gandi.icu/demo/service/system"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.AM_CONFIG.System.UseMultipoint {
		// 初始化redis服务
		initialize.Redis()
	}

	// 从db加载jwt数据
	if global.AM_DB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()

	// Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.AM_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.AM_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	server is running
`, address)
	global.AM_LOG.Error(s.ListenAndServe().Error())
}
