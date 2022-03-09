package response

import "gandi.icu/demo/model/system"

type SysAuthorityResponse struct {
	Authority system.SysAuthority `json:"authority"`
}
