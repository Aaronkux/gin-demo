package main

import (
	"encoding/json"
	"fmt"

	"gandi.icu/demo/global"
	"github.com/bwmarrin/snowflake"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type User struct {
	global.CommonModel
	Name      string  `json:"name"`
	Age       uint8   `json:"age"`
	CompanyId int64   `json:"company_id"`
	Company   Company `json:"company"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = global.SnowflakeID(node.Generate().Int64())
	return
}

type Company struct {
	global.CommonModel
	Name string
}

func (u *Company) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = global.SnowflakeID(node.Generate().Int64())
	return
}

var node *snowflake.Node
var snowflakeErr error

func main() {
	node, snowflakeErr = snowflake.NewNode(1)
	if snowflakeErr != nil {
		fmt.Println(snowflakeErr)
	}
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		// CreateBatchSize: 1000,
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
			TablePrefix:   "t_",
		},
		// 禁用物理外键, 使用逻辑外键
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	// db := db.Session(&gorm.Session{CreateBatchSize: 1000})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

	// company := Company{Name: "award"}

	// user := User{model.CommonModel: model.CommonModel{ID: 123456789}, Name: "aaaaa", Age: 19, Company: Company{model.CommonModel: model.CommonModel{ID: 1232414214214124}, Name: "award"}}
	// s, _ := json.Marshal(user)
	// blob := string(s)
	// fmt.Println(blob)
	// var user2 User
	// err = json.Unmarshal([]byte(blob), &user2)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(user2)

	cc := global.CommonModel{ID: 123456789}
	s, _ := json.Marshal(cc)
	blob := string(s)
	fmt.Println(blob)
	var cc2 global.CommonModel
	err = json.Unmarshal([]byte(blob), &cc2)
	if err != nil {
		panic(err)
	}
	fmt.Println(cc2)

	// dd := SnowflakeID(123456789)
	// s, _ := json.Marshal(dd)
	// blob := string(s)
	// fmt.Println(blob)
	// var dd2 SnowflakeID
	// err = json.Unmarshal([]byte(blob), &dd2)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(dd2)

	// err = db.Create(&user).Error

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// var company Company

	// err = db.Unscoped().First(&company).Delete(&company).Error

	// if err != nil {
	// 	fmt.Println(err)
	// }
}
