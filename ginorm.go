package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func main() {
	db, err := gorm.Open("mysql", "oyd:2w3e$R%T@tcp(62.234.122.162:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		panic(err.Error())
	}
	r := gin.Default()
	r.POST("/student", func(c *gin.Context) {

	})
}
