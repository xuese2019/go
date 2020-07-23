package user

import (
	"log"

	"github.com/xuese2019/demo/business/mysql"
)

// 新增
func Add(u UserStruct) *error {

	// 开启事务
	tx, err := mysql.Db.Begin()
	if err != nil {
		return &err
	}

	sqlStr := `insert into user_table (uuid,account,password) values (?,?,?)`
	// // 非事务绑定 start
	// ret, err := mysql.Db.Exec(sqlStr, u.Uuid, u.Account, u.Password)
	// if err != nil {
	// 	log.Panicln(err)
	// 	return _, &err
	// }
	// // 影响的行数
	// n, err := ret.RowsAffected()
	// if err != nil {
	// 	log.Panicln(err)
	// 	return _, &err
	// }
	// return n, _
	// // 非事务绑定 end

	// 事务绑定
	_, err = tx.Exec(sqlStr, u.Uuid, u.Account, u.Password)
	if err != nil {
		// 回滚事务
		tx.Rollback()
		log.Panicln(err)
		return &err
	}
	// 提交事务
	err = tx.Commit()
	if err != nil {
		// 回滚事务
		tx.Rollback()
		log.Panicln(err)
		return &err
	}
	return nil
}

// 修改
func Update(u UserStruct) *error {
	sqlStr := `update user_table set password = ? where uuid = ?`
	ret, err := mysql.Db.Exec(sqlStr, u.Password, u.Uuid)
	if err != nil {
		log.Panicln(err)
		return &err
	}
	// 影响的行数
	_, err = ret.RowsAffected()
	if err != nil {
		log.Panicln(err)
		return &err
	}
	return nil
}

// 修改
func Delete(uuid string) *error {
	sqlStr := `delete from user_table where uuid = ?`
	ret, err := mysql.Db.Exec(sqlStr, uuid)
	if err != nil {
		log.Panicln(err)
		return &err
	}
	// 影响的行数
	_, err = ret.RowsAffected()
	if err != nil {
		log.Panicln(err)
		return &err
	}
	return nil
}

// 根据帐号和密码查询
func FindByAccAndPwd(acc string, pwd string) *UserStruct {

	sqlStr := `select uuid,account,password from godemo.user_table where account = ? and password = ?`
	row := mysql.Db.QueryRow(sqlStr, acc, pwd)

	var u UserStruct
	// scan中带有close功能
	row.Scan(&u.Uuid, &u.Account, &u.Password)
	return &u
}

// 根据id查询
func FindById(uuid string) *UserStruct {

	sqlStr := `select uuid,account,password from godemo.user_table where uuid = ?`
	row := mysql.Db.QueryRow(sqlStr, uuid)

	var u UserStruct
	// scan中带有close功能
	row.Scan(&u.Uuid, &u.Account, &u.Password)
	return &u
}

// 查询所有
func FindAllLimit(acc string, pageSize int, pageNum int) *[]UserStruct {

	// 切片
	users := make([]UserStruct, 0)

	sqlStr := `select uuid,account,password from godemo.user_table limit ?,?`
	rows, err := mysql.Db.Query(sqlStr, (pageNum-1)*pageSize, pageSize)
	if err != nil {
		log.Println(err)
		return nil
	}

	defer rows.Close()

	// 动态查询条件
	// if acc != "" {
	// 	mysql.Db.Where("account like %?%", acc)
	// }
	// if pageSize > 0 && pageNum > 0 {
	// 	mysql.Db.Limit(pageSize).Offset((pageNum - 1) * pageSize)
	// }

	// 循环取值
	for rows.Next() {
		var u UserStruct
		err = rows.Scan(&u.Uuid, &u.Account, &u.Password)
		if err != nil {
			log.Panicln(err)
		}
		log.Println(u)
	}
	return &users
}
