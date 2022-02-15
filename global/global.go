package global

import (
	"gandi.icu/demo/config"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
)

var (
	AM_VP     *viper.Viper
	AM_CONFIG config.Server

	BlackCache local_cache.Cache
)
