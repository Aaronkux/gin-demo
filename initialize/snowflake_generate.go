package initialize

import (
	"gandi.icu/demo/global"
	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
)

func SnowflakeGenerate() {
	// 初始化Snowflake
	node, err := snowflake.NewNode(1)
	if err != nil {
		global.AM_LOG.Error("SnowflakeGenerate", zap.Error(err))
		panic(err)
	}
	global.AM_SNOWFLAKE = node
	global.AM_LOG.Info("snowflake init success")
}
