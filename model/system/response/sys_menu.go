package response

import (
	"gandi.icu/demo/model/system"
)

type SysMenuResponse struct {
	Menu system.SysMenu `json:"menu"`
}
