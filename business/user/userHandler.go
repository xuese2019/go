package user

import (
	"strconv"

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

	pageNow := c.Param("pageNow")
	pageSize := c.Param("pageSize")
	acc := c.PostForm("acc")

	ps, err := strconv.Atoi(pageSize)
	if err != nil {
		c.JSON(500, gin.H{"message": "每页条数必须是数字"})
	}
	pn, err2 := strconv.Atoi(pageNow)
	if err2 != nil {
		c.JSON(500, gin.H{"message": "当前页必须是数字"})
	}

	users := FindAllLimit(ps, pn, acc)
	c.JSON(200, gin.H{"data": users})
}
