package utils

var (
	IdVerify              = Rules{"ID": {NotEmpty()}}
	AuthorityCreateVerify = Rules{"AuthorityName": {NotEmpty()}}
	AuthorityUpdateVerify = Rules{"ID": {NotEmpty()}, "AuthorityName": {NotEmpty()}}
	RegisterVerify        = Rules{"Email": {NotEmpty()}, "Nickname": {NotEmpty()}, "Password": {NotEmpty()}, "AuthorityIds": {NotEmpty()}}
	PageInfoVerify        = Rules{"Page": {NotEmpty(), Gt("0")}, "PageSize": {NotEmpty(), Lt("50"), Gt("0")}}
	MenuCreateVerify      = Rules{"MenuName": {NotEmpty()}, "Path": {NotEmpty()}, "Hidden": {NotEmpty()}, "Icon": {NotEmpty()}, "ParentId": {NotEmpty()}, "Order": {NotEmpty()}}
	MenuUpdateVerify      = Rules{"ID": {NotEmpty()}, "MenuName": {NotEmpty()}, "Path": {NotEmpty()}, "Hidden": {NotEmpty()}, "Icon": {NotEmpty()}, "Order": {NotEmpty()}}
	AuthorityMenuVerify   = Rules{"AuthorityId": {NotEmpty()}, "MenuIds": {NotEmpty()}}
	AuthorityIdVerify     = Rules{"AuthorityId": {NotEmpty()}}
)
