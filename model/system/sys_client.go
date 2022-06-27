package system

import (
	"gandi.icu/demo/global"
	"github.com/shopspring/decimal"
)

type SysClient struct {
	global.CommonModel
	// **common part**
	// basic
	ClientType     string             `json:"clientType" gorm:"comment:'客户类型'"`
	Name           string             `json:"name" gorm:"comment:名称"`
	Email          string             `json:"email" gorm:"comment:邮箱"`
	Phone          string             `json:"phone" gorm:"comment:联系方式"`
	Purpose        string             `json:"purpose" gorm:"comment:汇款目的"`
	RegistrationId uint8              `json:"registrationId" gorm:"comment:注册ID"`
	Unsubscribe    bool               `json:"unsubscribe" gorm:"default:1;comment:是否取消订阅"`
	SaleID         global.SnowflakeID `json:"saleId" gorm:"comment:销售ID"`
	Sale           SysSale            `json:"sale" gorm:"foreignkey:SaleID"`
	// documents related
	Document1Front  string `json:"document1Front" gorm:"comment:证件1正面"`
	Document1Back   string `json:"document1Back" gorm:"comment:证件1反面"`
	Document1Expiry string `json:"document1Expiry" gorm:"comment:证件1有效期"`
	Document2Front  string `json:"document2Front" gorm:"comment:证件2正面"`
	Document2Back   string `json:"document2Back" gorm:"comment:证件2反面"`
	Document2Expiry string `json:"document2Expiry" gorm:"comment:证件2有效期"`
	Signature       string `json:"signature" gorm:"comment:签名"`

	// **individual part**
	// client
	Gender         string          `json:"gender" gorm:"comment:性别"`
	DOB            string          `json:"dob" gorm:"comment:生日"`
	Occupation     string          `json:"occupation" gorm:"comment:职业"`
	EmployerName   string          `json:"employerName" gorm:"comment:雇主名字"`
	AnnualIncome   decimal.Decimal `json:"annualIncome" gorm:"comment:年收入"`
	SourceOfIncome string          `json:"sourceOfIncome" gorm:"comment:收入来源"`
	// living address
	Address  string `json:"address" gorm:"comment:地址"`
	Suburb   string `json:"suburb" gorm:"comment:区/市"`
	State    string `json:"state" gorm:"comment:州"`
	Country  string `json:"country" gorm:"comment:国家"`
	Postcode string `json:"postcode" gorm:"comment:邮政编码"`
	// documents related
	FacePic string `json:"facePic" gorm:"comment:人脸照片"`

	// **company part**
	// client
	EntityType        string `json:"entityType" gorm:"comment:实体类型"`
	RegisteredAddress string `json:"registeredAddress" gorm:"comment:注册地址"`
	PrincipleAddress  string `json:"principleAddress" gorm:"comment:主要地址"`
	ABN_ACN_ARBN      string `json:"abn_acn_arbn" gorm:"comment:ABN/ACN/ARBN"`

	// primary account holder
	AccountHolderName     string `json:"accountHolderName" gorm:"comment:主要账户名"`
	AccountHolderDOB      string `json:"accountHolderDOB" gorm:"comment:主要账户生日"`
	AccountHolderPosition string `json:"accountHolderPosition" gorm:"comment:主要账户职位"`
	AccountHolderPhone    string `json:"accountHolderPhone" gorm:"comment:主要账户联系方式"`
	AccountHolderEmail    string `json:"accountHolderEmail" gorm:"comment:主要账户邮箱"`
	AccountHolderAddress  string `json:"accountHolderAddress" gorm:"comment:主要账户地址"`
	// documents related
	CompanyExtract string `json:"companyExtract" gorm:"comment:公司摘录"`

	Beneficiaries []SysBeneficiary `json:"beneficiaries" gorm:"foreignkey:ClientID"`
}
