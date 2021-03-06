package request

import (
	"gandi.icu/demo/model/common/request"
)

type SearchClientParams struct {
	request.PageInfo
	ClientType     int    `json:"clientType"`
	Name           string `json:"name"`
	RegistrationId uint8  `json:"registrationId"`
}
