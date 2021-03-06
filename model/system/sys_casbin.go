package system

type CasbinModel struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Ptype       string `json:"ptype" gorm:"column:ptype"`
	AuthorityId string `json:"rolename" gorm:"column:v0"`
	Path        string `json:"path" gorm:"column:v1"`
	Method      string `json:"method" gorm:"column:v2"`
}

func (CasbinModel) TableName() string {
	return "casbin_rule"
}
