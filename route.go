package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func middel() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("前")
		c.Next()
		fmt.Println("后")
	}
}

func middel1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("a")
		c.Next()
		fmt.Println("b")
	}
}

func middel2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("aa")
		c.Next()
		fmt.Println("bb")
	}
}

func main() {
	r := gin.Default()
	gin.DisableConsoleColor()
	v1 := r.Group("v1").Use(middel(), middel2()).Use(middel1())
	v1.GET("test", func(c *gin.Context) {
		fmt.Println("中")
		c.JSON(200, gin.H{
			"message": "succeed",
		})
	})
	r.Run(":8080")
}
