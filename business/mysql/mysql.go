package mysql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	var err error
	dsn := "root:root@tcp(127.0.0.1:3306)/godemo?charset=utf8"
	Db, err = sql.Open("mysql", dsn)

	// 校验dsn参数格式是否正确
	if err != nil {
		log.Printf("数据库参数校验失败{%s}，err:%v\n", dsn, err)
	}

	//尝试连接数据库
	err = Db.Ping()
	if err != nil {
		log.Printf("连接数据库失败{%s}，err:%v\n", dsn, err)
	} else {

		// 设置最大链接数
		Db.SetMaxOpenConns(15)
		// 最大空闲连接数
		Db.SetMaxIdleConns(5)

		log.Printf("数据库连接成功")
	}
}

// // 新增
// func Add(sql string) (*sql.Stmt, error) {
// 	return db.Prepare(sql)
// }

// // 删除
// func Remove(sql string) (*sql.Stmt, error) {
// 	return db.Prepare(sql)
// }

// // 修改
// func Ebit(sql string) (*sql.Stmt, error) {
// 	return db.Prepare(sql)
// }

// // 查询单个
// func One(sql string) *sql.Row {
// 	return db.QueryRow(sql)
// }

// // 查询所有
// func All(sql string) (*sql.Rows, error) {
// 	return db.Query(sql)
// }
