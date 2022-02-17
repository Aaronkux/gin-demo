package initialize

import (
	"fmt"

	"gandi.icu/demo/config"
	"gandi.icu/demo/global"
	"gandi.icu/demo/utils"
)

func Timer() {
	if global.AM_CONFIG.Timer.Start {
		for i := range global.AM_CONFIG.Timer.Detail {
			go func(detail config.Detail) {
				global.AM_Timer.AddTaskByFunc("ClearDB", global.AM_CONFIG.Timer.Spec, func() {
					err := utils.ClearTable(global.AM_DB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				})
			}(global.AM_CONFIG.Timer.Detail[i])
		}
	}
}
