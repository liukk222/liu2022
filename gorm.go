package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

// 创建表
func create(db *gorm.DB) {
	db.AutoMigrate(&Product{})

}

// 查找
func find(db *gorm.DB) {
	p := Product{}
	db.Find(&p, 1)
	fmt.Printf("p: %v\n", p)
	db.Find(&p, "Code=?", "1001")
	fmt.Printf("p: %v\n", p)
}

// 添加
func insert(db *gorm.DB) {
	p := Product{
		Code:  "1001",
		Price: 100,
	}
	db.Create(&p)
}

// 更新数据
func updat(db *gorm.DB) {
	p := Product{}
	db.Find(&p, 1)
	db.Model(&p).Update("Price", 1000)
	db.Model(&p).Updates(Product{Price: 1001, Code: "1002"})
	//db.Model(&p).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
}

// 删除
func del(db *gorm.DB) {
	p := Product{}
	db.Find(&p, 1)
	db.Delete(&p, 1)
}
func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("err")
	}
	del(db)
	//updat(db)
}
