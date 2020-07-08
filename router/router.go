package router

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	login "github.com/xuese2019/demo/business/login"
	user "github.com/xuese2019/demo/business/user"
)

func init() {
	router := gin.Default()

	// 加载html模板
	router.LoadHTMLGlob("templates/**/*")

	// 加载静态资源 参数1：页面写的路径 参数2：文件物理路径 从工程根目录开始算 不是当前文件所在位置
	router.Static("/static", "./static")
	// router.StaticFS("/static", http.Dir("../statics"))
	// router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	// 创建记录日志的文件
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// 上传文件大小
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	// html模板分组
	HtmlR := router.Group("")
	{
		// 根目录重定向至登录页面
		HtmlR.GET("/", func(c *gin.Context) {
			// 路由内部重定向
			c.Request.URL.Path = "/page/login"
			router.HandleContext(c)
		})

		// 页面分组
		HtmlR2 := HtmlR.Group("/page")
		{

			// 404
			HtmlR2.GET("/404", func(c *gin.Context) {
				c.HTML(200, "404.html", gin.H{
					"title": "error",
				})
			})
			// 500
			HtmlR2.GET("/500", func(c *gin.Context) {
				c.HTML(200, "500.html", gin.H{
					"title": "error",
				})
			})

			// 登录页面
			HtmlR2.GET("/login", func(c *gin.Context) {
				c.HTML(200, "login.html", gin.H{
					"title": "欢迎来到demo系统",
				})
			})

			// 登录后主页面
			HtmlR2.GET("/home", func(c *gin.Context) {
				c.HTML(200, "home.html", gin.H{
					"title": "系统",
				})
			})
		}
	}

	//以下的接口，都使用Authorize()中间件身份验证
	// 全局中间件Use
	// router.Use(Authorize())

	loginR := router.Group("/login")
	{
		loginR.POST("/login", login.Login)
	}

	userR := router.Group("/user")
	{
		userR.POST("/add", user.Add)
		userR.DELETE("/remove/:uuid", user.Remove)
		userR.PUT("/edit/:uuid", user.Edit)
		userR.GET("/one/:uuid", user.One)
		userR.GET("/page/:pageNow/:pageSize", user.Page)
	}

	router.Run(":9000")
}

// 判断是否登录
func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("auth")
		if auth != "" {
			// 验证通过，会继续访问下一个中间件
			c.Next()
		} else {
			// 验证不通过，不再调用后续的函数处理
			c.Abort()
			// 返回值
			c.JSON(500, gin.H{"result": 501, "message": "访问未授权"})
		}
	}
}
