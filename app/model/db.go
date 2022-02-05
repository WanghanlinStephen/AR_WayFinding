package model

import (
	"database/sql"
	"fmt"
	"pro/config"
)

var Db *sql.DB

//gorm链接数据库
func Run() error {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local",
			config.Mysql.UserName,
			config.Mysql.Password,
			config.Mysql.Host,
			config.Mysql.Port,
			config.Mysql.Database,
		),
	)
	if err != nil {
		return err
	}
	/*连接池信息*/
	//db.DB().SetMaxIdleConns(config.Mysql.MaxIdleConns) //设置最大空闲数
	//db.DB().SetMaxOpenConns(config.Mysql.MaxOpenConns) //设置最大连接数
	//db.SingularTable(true)
	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return "ww_" + defaultTableName
	//}
	Db = db
	return nil
}
