package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Hello struct {
	gorm.Model
	Name string
	Sex  bool
	Age  int
}

func main() {
	db, err := gorm.Open("mysql", "oyd:2w3e$R%T@tcp(62.234.122.162:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&Hello{})
	//db.Create(&Hello{
	//	Name: "bbb",
	//	Sex:  false,
	//	Age:  2,
	//})

	//查
	//var he Hello
	//db.First(&he, "name=?", "aaa")
	//var he []Hello
	//db.Where("age >? ", "10").Or("sex = ?", true).Find(&he)

	//改1个
	//db.Where("id = ?", "1").First(&Hello{}).Update("name", "zzz")

	//改多个
	//db.Where("id = ?", "1").First(&Hello{}).Updates(map[string]interface{}{
	//	"Name": "ccc",
	//	"Sex":  true,
	//	"Age":  30,
	//})
	//

	//db.Where("id in (?)", []int{1, 2}).Find(&[]Hello{}).Updates(map[string]interface{}{
	//	"Name": "ccc",
	//	"Sex":  true,
	//	"Age":  30,
	//})

	//删除数据 软删除
	//db.Where("id = ?", "1").Delete(&Hello{})

	//硬删除
	db.Where("id = ?", "2").Unscoped().Delete(&Hello{})

	defer db.Close()
}
