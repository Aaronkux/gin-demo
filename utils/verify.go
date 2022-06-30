package utils

var (
	IdVerify  = Rules{"ID": {NotEmpty()}}
	ApiVerify = Rules{"Path": {NotEmpty()}, "Description": {NotEmpty()}, "ApiGroup": {NotEmpty()}, "Method": {NotEmpty()}}

	CreateAuthorityVerify = Rules{"AuthorityName": {NotEmpty()}}
	UpdateAuthorityVerify = Rules{"ID": {NotEmpty()}, "AuthorityName": {NotEmpty()}}

	RegisterVerify         = Rules{"Email": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty(), Gt("8"), RegexpMatch("[a-z]+"), RegexpMatch("[A-Z]+"), RegexpMatch("[0-9]")}, "AuthorityIds": {NotEmpty()}, "Avatar": {NotEmpty()}}
	UpdateSelfVerify       = Rules{"NickName": {NotEmpty()}}
	GetUserListVerify      = Rules{"Page": {NotEmpty(), Gt("0")}, "PageSize": {NotEmpty(), Lt("101"), Gt("0")}, "IsActive": {NotEmpty()}}
	UpdateUserVerify       = Rules{"ID": {NotEmpty()}, "Email": {NotEmpty()}, "NickName": {NotEmpty()}, "Avatar": {NotEmpty()}, "AuthorityIds": {NotEmpty()}, "IsActive": {NotEmpty()}}
	UpdateUserAvatarVerify = Rules{"ID": {NotEmpty()}, "File": {NotEmpty()}}
	LoginVerify            = Rules{"CaptchaId": {NotEmpty()}, "Captcha": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}}

	PageInfoVerify = Rules{"Page": {NotEmpty(), Gt("0")}, "PageSize": {NotEmpty(), Lt("101"), Gt("0")}}

	CreateMenuVerify = Rules{"MenuName": {NotEmpty()}, "Path": {NotEmpty()}, "Hidden": {NotEmpty()}, "ParentId": {NotEmpty()}, "MenuKey": {NotEmpty()}}
	UpdateMenuVerify = Rules{"ID": {NotEmpty()}, "MenuName": {NotEmpty()}, "Path": {NotEmpty()}, "Hidden": {NotEmpty()}, "MenuKey": {NotEmpty()}}

	SetAuthorityVerify = Rules{"AuthorityId": {NotEmpty()}, "MenuIds": {NotEmpty()}}
	AuthorityIdVerify  = Rules{"AuthorityId": {NotEmpty()}}

	CreateBranchVerify = Rules{"Name": {NotEmpty()}}
	UpdateBranchVerify = Rules{"ID": {NotEmpty()}, "Name": {NotEmpty()}}

	CreateSaleVerify       = Rules{"Name": {NotEmpty()}, "Avatar": {NotEmpty()}, "Email": {NotEmpty()}, "BranchId": {NotEmpty()}}
	GetSaleListVerify      = Rules{"Page": {NotEmpty(), Gt("0")}, "PageSize": {NotEmpty(), Lt("101"), Gt("0")}, "IsActive": {NotEmpty()}}
	UpdateSalesVerify      = Rules{"Name": {NotEmpty()}, "Email": {NotEmpty()}, "BranchId": {NotEmpty()}, "IsActive": {NotEmpty()}}
	UpdateSaleAvatarVerify = Rules{"ID": {NotEmpty()}, "File": {NotEmpty()}}

	CreateReferralVerify = Rules{"Name": {NotEmpty()}}
	UpdateReferralVerify = Rules{"Name": {NotEmpty()}}

	CreateClientVerify  = Rules{"ClientType": {NotEmpty()}, "Name": {NotEmpty()}, "Gender": {NotEmpty()}}
	GetClientListVerify = Rules{"Page": {NotEmpty(), Gt("0")}, "PageSize": {NotEmpty(), Lt("101"), Gt("0")}, "ClientType": {NotEmpty()}}
	UpdateClientVerify  = Rules{"ID": {NotEmpty()}, "ClientType": {NotEmpty()}, "Name": {NotEmpty()}, "Email": {NotEmpty()}}

	CreateViewVerify             = Rules{"UserID": {NotEmpty()}, "Type": {NotEmpty()}, "Value": {NotEmpty()}}
	GetViewByUserIdAndTypeVerify = Rules{"UserID": {NotEmpty()}, "Type": {NotEmpty()}}
	UpdateViewVerify             = Rules{"ID": {NotEmpty()}, "Value": {NotEmpty()}}
)
