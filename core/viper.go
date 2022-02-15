package core

import (
	"fmt"
	"path/filepath"
	"time"

	"gandi.icu/demo/global"
	"gandi.icu/demo/utils"
	"github.com/fsnotify/fsnotify"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
)

func Viper() *viper.Viper {
	config := utils.ConfigFile

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")

	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.AM_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&global.AM_CONFIG); err != nil {
		fmt.Println(err)
	}

	global.AM_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(time.Second * time.Duration(global.AM_CONFIG.JWT.ExpiresTime)),
	)
	return v
}
