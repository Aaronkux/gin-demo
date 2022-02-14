package main

import (
	"fmt"
	"time"

	"github.com/bwmarrin/snowflake"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type CommonModel struct {
	ID        int64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	CommonModel
	Name      string
	Age       uint8
	CompanyId int64
	Company   Company
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = node.Generate().Int64()
	return
}

type Company struct {
	CommonModel
	Name string
}

func (u *Company) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = node.Generate().Int64()
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

	// user := User{Name: "aaaaa", Age: 19, Company: company}

	// err = db.Create(&user).Error

	// if err != nil {
	// 	fmt.Println(err)
	// }

	var company Company

	err = db.Unscoped().First(&company).Delete(&company).Error

	if err != nil {
		fmt.Println(err)
	}
}
