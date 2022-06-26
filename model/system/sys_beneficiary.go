package system

import (
	"gandi.icu/demo/global"
	"gorm.io/gorm"
)

type SysBeneficiary struct {
	global.CommonModel
	// basic
	BankName      string `json:"bankName" gorm:"comment:银行名称"`
	BranchName    string `json:"branchName" gorm:"comment:银行支行名称"`
	AccountName   string `json:"accountName" gorm:"comment:账户名称"`
	AccountNumber string `json:"accountNumber" gorm:"comment:账户号码"`

	TrustAccount   bool   `json:"trustAccount" gorm:"default:0;comment:是否信托账户"`
	CompanyName    string `json:"companyName" gorm:"comment:公司名称"`
	CompanyAddress string `json:"companyAddress" gorm:"comment:公司地址"`
	CompanyABN     string `json:"companyABN" gorm:"comment:公司ABN"`

	BSB_SwiftCode string `json:"bsb_swiftCode" gorm:"comment:BSB/SwiftCode"`
	USD_Address   string `json:"usd_address" gorm:"comment:USD地址"`
	IBAN          string `json:"iban" gorm:"comment:IBAN"`
	RoutingNumber string `json:"routingNumber" gorm:"comment:RoutingNumber"`

	// address
	Address  string `json:"address" gorm:"comment:地址"`
	Suburb   string `json:"suburb" gorm:"comment:地区"`
	State    string `json:"state" gorm:"comment:省份"`
	Postcode string `json:"postcode" gorm:"comment:邮编"`
	Country  string `json:"country" gorm:"comment:国家"`

	// contact
	Phone      string `json:"phone" gorm:"comment:电话"`
	Relation   string `json:"relation" gorm:"comment:关系"`
	Occupation string `json:"occupation" gorm:"comment:职业"`
	Purpose    string `json:"purpose" gorm:"comment:用途"`
	Reference  string `json:"reference" gorm:"comment:参考备注"`

	// document
	DocumentFront string `json:"documentFront" gorm:"comment:证件正面"`
	DocumentBack  string `json:"documentBack" gorm:"comment:证件反面"`

	RelatedDoc string `json:"relatedDoc" gorm:"comment:相关文件"`

	ClientID global.SnowflakeID `json:"clientId" gorm:"comment:客户ID"`
}

func (beneficiary *SysBeneficiary) BeforeCreate(tx *gorm.DB) (err error) {
	beneficiary.ID = global.SnowflakeID(global.AM_SNOWFLAKE.Generate().Int64())
	return
}
