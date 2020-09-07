package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping/:id", func(c *gin.Context) {
		id := c.Param("id")                 //uri传参数，使用Param
		user := c.Query("user")             //query传参，使用Query
		num := c.DefaultQuery("num", "oyd") //如果没有传，默认是这个
		c.JSON(200, gin.H{
			"message": id,
			"user":    user,
			"num":     num,
		})
	})
	r.POST("/ping", func(c *gin.Context) {
		user := c.DefaultPostForm("user", "ouyd") //从form表单里面取数据
		num := c.PostForm("num")
		c.JSON(200, gin.H{
			"message": "pong",
			"user":    user,
			"num":     num,
		})
	})
	r.DELETE("/ping/:id", func(c *gin.Context) {
		id := c.Param("id") //uri传参数，使用Param
		c.JSON(200, gin.H{
			"message": id,
		})
	})
	r.PUT("/ping", func(c *gin.Context) {
		user := c.DefaultPostForm("user", "ouyd") //从form表单里面取数据
		num := c.PostForm("num")
		c.JSON(200, gin.H{
			"message": "pong",
			"user":    user,
			"num":     num,
		})
	})
	r.Run(":1010") // listen and serve on 0.0.0.0:8080
}
