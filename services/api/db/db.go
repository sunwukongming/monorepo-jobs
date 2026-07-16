package db

import (
	"app/config"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"

	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var xormConnection *xorm.Engine

func Reload() error {
	dbConfig := config.Get().DBS.Default
	preDB := xormConnection
	xormConnection = xormConnect(dbConfig, "")
	if preDB != nil {
		go func(engine *xorm.Engine) {
			time.Sleep(time.Second * 10)
			engine.Close()
		}(preDB)
	}
	return nil
}

func Default() *xorm.Engine {
	return xormConnection
}

func xormConnect(dbConfig config.DB, prefix string) *xorm.Engine {
	connectionString := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.DBName)
	engine, err := xorm.NewEngine("mysql", connectionString)
	if err != nil {
		logrus.Fatal(err)
		return nil
	}
	tbMapper := names.NewPrefixMapper(names.SnakeMapper{}, prefix)
	engine.SetTableMapper(tbMapper)
	engine.SetConnMaxLifetime(time.Second * 5)
	engine.SetMaxIdleConns(dbConfig.PoolSize / 10)
	engine.SetMaxOpenConns(dbConfig.PoolSize)
	local, err := time.LoadLocation(config.Get().TimeZone)
	if err != nil {
		logrus.Fatal("时区设置错误")
		return nil
	}
	engine.SetTZDatabase(local)
	engine.SetTZLocation(local)
	engine.ShowSQL(true)
	err = engine.Ping()
	if err != nil {
		logrus.Fatal(err)
		return nil
	}
	return engine
}
