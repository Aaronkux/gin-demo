package global

import (
	"gandi.icu/demo/config"
	"gandi.icu/demo/utils/timer"
	"github.com/bwmarrin/snowflake"
	"github.com/go-redis/redis/v8"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	AM_VP                  *viper.Viper
	AM_DB                  *gorm.DB
	AM_DBList              map[string]*gorm.DB
	AM_REDIS               *redis.Client
	AM_CONFIG              config.Server
	AM_LOG                 *zap.Logger
	AM_Timer               timer.Timer = timer.NewTimerTask()
	AM_SNOWFLAKE           *snowflake.Node
	AM_Concurrency_Control = &singleflight.Group{}

	BlackCache local_cache.Cache
)
