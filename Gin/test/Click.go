package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Static("/static", "../web/static")
	r.LoadHTMLGlob("../website/*")

	//登录跳转天气查询界面
	r.POST("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "weather.html", nil)
	})

	//发送给邮箱界面
	r.GET("/E-mail.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "E-mail.html", gin.H{})
	})
	//注册界面
	r.GET("/register.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", gin.H{})
	})
	//主界面
	r.GET("/", func(c *gin.Context) {
		c.File("../website/index.html")
	})

	r.Run(":8080")
}
