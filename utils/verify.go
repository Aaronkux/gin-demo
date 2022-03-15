package utils

var (
	IdVerify              = Rules{"ID": {NotEmpty()}}
	ApiVerify             = Rules{"Path": {NotEmpty()}, "Description": {NotEmpty()}, "ApiGroup": {NotEmpty()}, "Method": {NotEmpty()}}
	AuthorityCreateVerify = Rules{"AuthorityName": {NotEmpty()}}
	AuthorityUpdateVerify = Rules{"ID": {NotEmpty()}, "AuthorityName": {NotEmpty()}}
	RegisterVerify        = Rules{"Email": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty()}, "AuthorityIds": {NotEmpty()}}
	LoginVerify           = Rules{"CaptchaId": {NotEmpty()}, "Captcha": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}}
	PageInfoVerify        = Rules{"Page": {NotEmpty(), Gt("0")}, "PageSize": {NotEmpty(), Lt("50"), Gt("0")}}
	MenuCreateVerify      = Rules{"MenuName": {NotEmpty()}, "Path": {NotEmpty()}, "Hidden": {NotEmpty()}, "Icon": {NotEmpty()}, "ParentId": {NotEmpty()}, "Order": {NotEmpty()}}
	MenuUpdateVerify      = Rules{"ID": {NotEmpty()}, "MenuName": {NotEmpty()}, "Path": {NotEmpty()}, "Hidden": {NotEmpty()}, "Icon": {NotEmpty()}, "Order": {NotEmpty()}}
	AuthorityMenuVerify   = Rules{"AuthorityId": {NotEmpty()}, "MenuIds": {NotEmpty()}}
	AuthorityIdVerify     = Rules{"AuthorityId": {NotEmpty()}}
)
