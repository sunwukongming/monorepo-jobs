package db

import (
	"app/internal/db/mysql"

	"gorm.io/gorm"
)

// Default 返回 GORM 会话。
//
// 每次调用都基于底层连接派生一个全新的 *gorm.DB Session，
// 以对齐旧 xorm Engine「每次链式调用产生独立会话、条件互不污染」的语义，
// 避免不同查询之间的 WHERE / SELECT 等子句相互串联。
func Default() *gorm.DB {
	return mysql.Gorm().Session(&gorm.Session{})
}

// Get 兼容旧 xorm `Session.Get` 的语义：
// 取满足条件的第一条记录写入 dest，返回记录是否存在。
// 「未找到」不作为 error 返回（与 xorm 一致，区别于 GORM 的 First/ErrRecordNotFound）。
//
// 用法：ok, err := db.Get(db.Default().Where("id = ?", id), &account)
func Get(tx *gorm.DB, dest interface{}) (bool, error) {
	res := tx.Limit(1).Find(dest)
	if res.Error != nil {
		return false, res.Error
	}
	return res.RowsAffected > 0, nil
}

// TableName 解析 model 对应的数据库表名（替代旧 xorm Engine.TableName）。
func TableName(model interface{}) string {
	stmt := &gorm.Statement{DB: mysql.Gorm()}
	if err := stmt.Parse(model); err != nil {
		return ""
	}
	return stmt.Schema.Table
}
