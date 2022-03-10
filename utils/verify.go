package utils

var (
	AuthorityVerify   = Rules{"AuthorityName": {NotEmpty()}}
	AuthorityIdVerify = Rules{"ID": {NotEmpty()}}
	RegisterVerify    = Rules{"Email": {NotEmpty()}, "Nickname": {NotEmpty()}, "Password": {NotEmpty()}, "AuthorityIds": {NotEmpty()}}
	PageInfoVerify    = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
)
