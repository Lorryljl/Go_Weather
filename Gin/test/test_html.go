package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Static("/static", "../web/static")
	r.LoadHTMLGlob("../website/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message": "hello world",
		})
	})
	r.GET("/register.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"css": "/static/register.css",
			"js":  "/static/register.js",
		})
	})

	r.Run(":8080")
}
