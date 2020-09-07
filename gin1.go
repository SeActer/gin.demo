package main

import (
	"github.com/gin-gonic/gin/binding"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

// 绑定为json
type Login struct {
	User     string `json:"user" binding:"required,bigData"`
	Password string `json:"password" binding:"required,bigData"`
}

//type Login struct {
//	User     string `json:"user" uri:"user" form:"user"`
//	Password string `json:"password" uri:"password" form:"password"`
//}

func bigData(fl validator.FieldLevel) bool {
	if len(fl.Field().Interface().(string)) > 5 {
		return true
	}
	return false
}
func main() {
	router := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bigData", bigData)
	}
	// Example for binding JSON ({"user": "manu", "password": "123"})
	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		err := c.ShouldBindJSON(&json) //和boby里面的json转
		if err != nil {
			c.JSON(400, gin.H{
				"code": "error",
				"data": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"data": json,
				"code": "succeed",
			})
		}

	})

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}
