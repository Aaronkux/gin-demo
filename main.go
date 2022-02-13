package main

import (
	"fmt"
	"time"

	"github.com/bwmarrin/snowflake"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type CommonModel struct {
	ID        int64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	CommonModel
	Name string
	Age  uint8
}

var node *snowflake.Node
var snowflakeErr error

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	id := node.Generate()
	fmt.Println(id.Time())
	u.ID = id.Int64()
	return
}

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
	})

	// db := db.Session(&gorm.Session{CreateBatchSize: 1000})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

	var users = []User{{Name: "jinzhu1", Age: 19}, {Name: "jinzhu2", Age: 19}, {Name: "jinzhu3", Age: 19}}
	err = db.Create(&users).Error

	if err != nil {
		fmt.Println(err)
	}
	// db.Create(&User{Name: "aaron", Age: 26})

	// var user User

	// db.First(&user)
	// fmt.Println(user)
	// db.First(&user, 2)
	// fmt.Println(user)
}
