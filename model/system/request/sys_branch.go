package request

import "gandi.icu/demo/model/common/request"

type CreateBranch struct {
	Name string `json:"name"`
}

type SearchBranchParams struct {
	request.PageInfo
	Name string `json:"name"`
}
