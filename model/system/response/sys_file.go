package response

import "gandi.icu/demo/model/system"

type SysFileResponse struct {
	File system.SysFile `json:"file"`
}
