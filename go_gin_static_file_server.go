package main

import (
        "github.com/gin-gonic/gin"
)

func main() {
        router := gin.New()
        router.GET("/hello", func(c *gin.Context) {
                c.String(200, "Hello, World!")
        })
        router.Static("/static", "/var/www/nadargobind.org/public")
        router.Run(":8080")
}
