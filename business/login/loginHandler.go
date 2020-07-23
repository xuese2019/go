package login

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	user "github.com/xuese2019/demo/business/user"
)

// 表单形式
func Login(c *gin.Context) {
	acc := c.PostForm("acc")
	pwd := c.PostForm("pwd")

	u := user.FindByAccAndPwd(acc, pwd)
	if u.Uuid == "" {
		c.JSON(500, gin.H{"result": "帐号或密码错误"})
		return
	}

	// c.JSON(200, gin.H{"result": acc + pwd})

	// tk := token.CreateToken(acc)
	// c.Request.Header.Set("auth", tk)
	// c.JSON(200, gin.H{"status": "0", "result": acc + pwd})
	c.Redirect(http.StatusMovedPermanently, "/page/home")
}

// 表单形式
func Logout(c *gin.Context) {
	fmt.Println("注销")
}
