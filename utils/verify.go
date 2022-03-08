package utils

var (
	RegisterVerify = Rules{"Email": {NotEmpty()}, "Nickname": {NotEmpty()}, "Password": {NotEmpty()}, "AuthorityIds": {NotEmpty()}}
)
