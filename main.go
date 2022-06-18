package main

import (
	"github.com/gin-gonic/gin"
)

func main () {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/css", "./public/css")
	router.Static("/img", "./public/img")
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H {
			"title": "Ahmad Lamaul Farid | Profile",
		})
	})
	router.Run(":8080")
}