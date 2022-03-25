package response

import "gandi.icu/demo/model/system"

type SysSaleResponse struct {
	Sale system.SysSale `json:"sale"`
}
