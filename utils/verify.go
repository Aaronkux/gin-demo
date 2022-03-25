package utils

var (
	IdVerify              = Rules{"ID": {NotEmpty()}}
	ApiVerify             = Rules{"Path": {NotEmpty()}, "Description": {NotEmpty()}, "ApiGroup": {NotEmpty()}, "Method": {NotEmpty()}}
	AuthorityCreateVerify = Rules{"AuthorityName": {NotEmpty()}}
	AuthorityUpdateVerify = Rules{"ID": {NotEmpty()}, "AuthorityName": {NotEmpty()}}
	RegisterVerify        = Rules{"Email": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty()}, "AuthorityIds": {NotEmpty()}}
	UpdateSelfVerify      = Rules{"NickName": {NotEmpty()}}
	LoginVerify           = Rules{"CaptchaId": {NotEmpty()}, "Captcha": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}}
	PageInfoVerify        = Rules{"Page": {NotEmpty(), Gt("0")}, "PageSize": {NotEmpty(), Lt("101"), Gt("0")}}
	MenuCreateVerify      = Rules{"MenuName": {NotEmpty()}, "Path": {NotEmpty()}, "Hidden": {NotEmpty()}, "ParentId": {NotEmpty()}}
	MenuUpdateVerify      = Rules{"ID": {NotEmpty()}, "MenuName": {NotEmpty()}, "Path": {NotEmpty()}, "Hidden": {NotEmpty()}}
	AuthorityMenuVerify   = Rules{"AuthorityId": {NotEmpty()}, "MenuIds": {NotEmpty()}}
	AuthorityIdVerify     = Rules{"AuthorityId": {NotEmpty()}}
	BranchCreateVerify    = Rules{"Name": {NotEmpty()}}
	BranchUpdateVerify    = Rules{"ID": {NotEmpty()}, "Name": {NotEmpty()}}
	SaleCreateVerify      = Rules{"Name": {NotEmpty()}, "Email": {NotEmpty()}, "BranchId": {NotEmpty()}}
)
