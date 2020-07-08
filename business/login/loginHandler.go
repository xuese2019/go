package login

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 表单形式
func Login(c *gin.Context) {
	acc := c.PostForm("acc")
	pwd := c.PostForm("pwd")
	if strings.EqualFold(acc, pwd) {
		// c.JSON(200, gin.H{"status": "0", "result": acc + pwd})
		c.Redirect(http.StatusMovedPermanently, "/page/home")
	} else {
		// c.JSON(500, gin.H{"status": "-1", "result": "帐号密码错误"})
	}
}

// 表单形式
func Logout(c *gin.Context) {
	fmt.Println("注销")
}

// 表单形式
// func Upload(c *gin.Context) {
// 	// single file
// 	file, _ := c.FormFile("file")

// 	// Upload the file to specific dst.
// 	//c.SaveUploadedFile(file, dst)

// 	c.String(200, fmt.Sprintf("'%s' uploaded!", file.Filename))
// }
