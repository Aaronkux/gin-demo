package response

import "gandi.icu/demo/model/system"

type SysClientResponse struct {
	Client system.SysClient `json:"client"`
}
