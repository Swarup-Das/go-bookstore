package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)
var (
	db * gorm.DB
)

func Connect(){
	d, err := gorm.Open("mysql", "root:Wwervr@1234/temp1?charset=utf8&parseTime=True&loc=local")
	if err != nil{
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
}