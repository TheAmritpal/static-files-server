package main

import (
        "github.com/gin-gonic/gin"
)

func main() {
        router := gin.New()
	router.Use(func(c *gin.Context) {
        	c.Header("Access-Control-Allow-Origin", "*")
        	c.Next()
	})
        router.GET("/hello", func(c *gin.Context) {
                c.String(200, "Hello, World!")
        })
        router.Static("/static", "/var/www/nadargobind.org/public")
        router.Run(":8080")
}
