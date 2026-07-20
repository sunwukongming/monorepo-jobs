package mysql

import (
	"app/config"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var gromConnection *gorm.DB

func Reload() error {
	dbConfig := config.Get().DBS.Default
	preDB := gromConnection
	gromConnection = gormConnect(dbConfig)
	if preDB != nil {
		go func(db *gorm.DB) {
			time.Sleep(time.Second * 10)
			sql, err := db.DB()
			if err != nil {
				return
			}
			sql.Close()
		}(preDB)
	}
	return nil
}

func Gorm() *gorm.DB {
	return gromConnection
}

// SetGorm 直接设置底层 GORM 连接。主要用于测试时注入 mock/内存数据库。
func SetGorm(db *gorm.DB) {
	gromConnection = db
}

func gormConnect(config config.DB) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local", config.Username, config.Password, config.Host, config.Port, config.DBName)
	connection, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.Prefix,
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatal("数据库连接失败", err)
	}
	sqlDB, err := connection.DB()
	if err != nil {
		log.Fatal("数据库连接创建失败", err)
	}
	sqlDB.SetConnMaxIdleTime(time.Second * 10)
	sqlDB.SetConnMaxLifetime(time.Second * 10)
	sqlDB.SetMaxIdleConns(config.PoolSize / 10)
	sqlDB.SetMaxOpenConns(config.PoolSize)
	return connection
}
