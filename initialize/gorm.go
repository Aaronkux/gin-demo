package initialize

import (
	"os"

	"gandi.icu/demo/config"
	"gandi.icu/demo/global"
	"gandi.icu/demo/initialize/internal"
	"gandi.icu/demo/model/system"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	m := global.AM_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config()); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

// GormMysqlByConfig 初始化Mysql数据库用过传入配置
func GormMysqlByConfig(m config.DB) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config()); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

// User 拥有并属于多种 language，`user_languages` 是连接表
type User struct {
	gorm.Model
	Languages []*Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name  string
	Users []*User `gorm:"many2many:user_languages;"`
}

// RegisterTables 注册数据库表专用
// Author SliverHorn
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		system.JwtBlacklist{},
		system.SysUser{},
		User{},
		Language{},
	)
	if err != nil {
		global.AM_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.AM_LOG.Info("register table success")
}
