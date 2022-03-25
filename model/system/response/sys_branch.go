package response

import "gandi.icu/demo/model/system"

type SysBranchResponse struct {
	Branch system.SysBranch `json:"branch"`
}
