package response

import "gandi.icu/demo/model/system"

type SysViewResponse struct {
	View system.SysView `json:"view"`
}
