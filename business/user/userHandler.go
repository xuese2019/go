package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/xuese2019/demo/business/mysql"
)

// 新增
func Add(c *gin.Context) {
	c.JSON(200, gin.H{"result": 200, "message": "成功"})
}

// 删除
func Remove(c *gin.Context) {
	c.JSON(200, gin.H{"result": 200, "message": "成功"})
}

// 修改
func Edit(c *gin.Context) {
	c.JSON(200, gin.H{"result": 200, "message": "成功"})
}

// 查询单体
func One(c *gin.Context) {
	c.JSON(200, gin.H{"result": 200, "message": "成功"})
}

// 分页查询
func Page(c *gin.Context) {
	d, err := db.All("select uuid,account,password from user_table")
	if err != nil {
		fmt.Println(err)
	}
	var u UserStruct
	d.Scan(u.Uuid, u.Account, u.Password)
	fmt.Println(u)
	c.JSON(200, gin.H{"result": 200, "message": "成功"})
}
