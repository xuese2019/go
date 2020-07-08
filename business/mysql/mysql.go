package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/godemo?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}

	//尝试链接数据库
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
}

// 新增
func Add(sql string) (*sql.Stmt, error) {
	return db.Prepare(sql)
}

// 删除
func Remove(sql string) (*sql.Stmt, error) {
	return db.Prepare(sql)
}

// 修改
func Ebit(sql string) (*sql.Stmt, error) {
	return db.Prepare(sql)
}

// 查询单个
func One(sql string) *sql.Row {
	return db.QueryRow(sql)
}

// 查询所有
func All(sql string) (*sql.Rows, error) {
	return db.Query(sql)
}
