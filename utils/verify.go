package utils

var (
	IdVerify              = Rules{"ID": {NotEmpty()}}
	ApiVerify             = Rules{"Path": {NotEmpty()}, "Description": {NotEmpty()}, "ApiGroup": {NotEmpty()}, "Method": {NotEmpty()}}
	AuthorityCreateVerify = Rules{"AuthorityName": {NotEmpty()}}
	AuthorityUpdateVerify = Rules{"ID": {NotEmpty()}, "AuthorityName": {NotEmpty()}}
	RegisterVerify        = Rules{"Email": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty(), Gt("8"), RegexpMatch("[a-z]+"), RegexpMatch("[A-Z]+"), RegexpMatch("[0-9]")}, "AuthorityIds": {NotEmpty()}}
	UpdateSelfVerify      = Rules{"NickName": {NotEmpty()}}
	GetUserListVerify     = Rules{"Page": {NotEmpty(), Gt("0")}, "PageSize": {NotEmpty(), Lt("101"), Gt("0")}, "IsActive": {NotEmpty()}}
	UserUpdateVerify      = Rules{"ID": {NotEmpty()}, "NickName": {NotEmpty()}, "Avatar": {NotEmpty()}, "AuthorityIds": {NotEmpty()}, "IsActive": {NotEmpty()}}
	LoginVerify           = Rules{"CaptchaId": {NotEmpty()}, "Captcha": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}}
	PageInfoVerify        = Rules{"Page": {NotEmpty(), Gt("0")}, "PageSize": {NotEmpty(), Lt("101"), Gt("0")}}
	MenuCreateVerify      = Rules{"MenuName": {NotEmpty()}, "Path": {NotEmpty()}, "Hidden": {NotEmpty()}, "ParentId": {NotEmpty()}, "MenuKey": {NotEmpty()}}
	MenuUpdateVerify      = Rules{"ID": {NotEmpty()}, "MenuName": {NotEmpty()}, "Path": {NotEmpty()}, "Hidden": {NotEmpty()}, "MenuKey": {NotEmpty()}}
	AuthorityMenuVerify   = Rules{"AuthorityId": {NotEmpty()}, "MenuIds": {NotEmpty()}}
	AuthorityIdVerify     = Rules{"AuthorityId": {NotEmpty()}}
	BranchCreateVerify    = Rules{"Name": {NotEmpty()}}
	BranchUpdateVerify    = Rules{"ID": {NotEmpty()}, "Name": {NotEmpty()}}
	SaleCreateVerify      = Rules{"Name": {NotEmpty()}, "Avatar": {NotEmpty()}, "Email": {NotEmpty()}, "BranchId": {NotEmpty()}}
	GetSaleListVerify     = Rules{"Page": {NotEmpty(), Gt("0")}, "PageSize": {NotEmpty(), Lt("101"), Gt("0")}, "IsActive": {NotEmpty()}}
	SaleUpdateVerify      = Rules{"Name": {NotEmpty()}, "Avatar": {NotEmpty()}, "Email": {NotEmpty()}, "BranchId": {NotEmpty()}, "IsActive": {NotEmpty()}}
	ReferralCreateVerify  = Rules{"Name": {NotEmpty()}}
	GetReferralListVerify = Rules{"Page": {NotEmpty(), Gt("0")}, "PageSize": {NotEmpty(), Lt("101"), Gt("0")}}
	ReferralUpdateVerify  = Rules{"Name": {NotEmpty()}}
)
