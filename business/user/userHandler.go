package user

import (
	"github.com/gin-gonic/gin"
)

// 新增
func AddUser(c *gin.Context) {
	c.JSON(200, gin.H{"result": 200, "message": "成功"})
}

// 删除
func RemoveUser(c *gin.Context) {
	c.JSON(200, gin.H{"result": 200, "message": "成功"})
}

// 修改
func UpdateUser(c *gin.Context) {
	c.JSON(200, gin.H{"result": 200, "message": "成功"})
}

// 查询单体
func OneUser(c *gin.Context) {
	c.JSON(200, gin.H{"result": 200, "message": "成功"})
}

// 分页查询
func PageUser(c *gin.Context) {
	c.JSON(200, gin.H{"result": 200, "message": "成功"})
}
