package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//type User struct {
//	gorm.Model
//	Name string `gorm:"primary_key;column:user_name;type:varchar(100)"` //主键
//
//}
//
//func (u User) TableName() string {
//	return "ouyd"
//}

type Student struct {
	StudentName string
	gorm.Model
	ClassID  uint
	Card     Card
	Teachers []Teacher `gorm:"many2many:student_teacher;"`
}

type Teacher struct {
	gorm.Model
	TeacherName string
	Students    []Student `gorm:"many2many:student_teacher;"`
}

type Class struct {
	gorm.Model
	ClassName string
	Students  []Student
}

type Card struct {
	gorm.Model
	StudentID uint
	Num       int
}

/*
primary_key主键
column:user_name 自定义列名
type:'varchar(100)'  定义类型，长度
not null 非空
unique 唯一
-  忽略这个字段


*/

func main() {
	db, err := gorm.Open("mysql", "oyd:2w3e$R%T@tcp(62.234.122.162:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&Card{}, &Class{}, &Teacher{}, &Student{})
	//c := Card{
	//	Num: 987654,
	//}
	//s := Student{
	//	StudentName: "ouyd5",
	//	Card:        c,
	//}
	//t := Teacher{
	//	TeacherName: "ABC",
	//	Students:    []Student{s},
	//}
	//cl := Class{
	//	ClassName: "3",
	//	Students:  []Student{s},
	//}
	//_ = db.Create(&cl).Error
	//_ = db.Create(&t).Error
	r := gin.Default()
	r.POST("/student", func(c *gin.Context) {
		var student Student
		_ = c.BindJSON(&student)
		db.Create(&student)
	})
	r.GET("/student/:id", func(c *gin.Context) {
		id := c.Param("id")
		var student Student
		_ = c.BindJSON(&student)
		db.Preload("Teachers").Preload("Card").First(&student, "id = ?", id)
		c.JSON(200, gin.H{
			"s": student,
		})
	})

	r.GET("/class/:id", func(c *gin.Context) {
		id := c.Param("id")
		var class Class
		_ = c.BindJSON(&class)
		db.Preload("Students").Preload("Students.Teachers").First(&class, "id = ?", id)
		c.JSON(200, gin.H{
			"c": class,
		})
	})
	r.Run(":8888")

}
