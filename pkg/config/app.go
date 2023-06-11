package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// This file puropse is to return a variable called db, which will help the ohther file to intreact with db
var (
	db * gorm.DB
)

func Connect(){
	d, err := gorm.Open("mysql", "user:password/table_name?charset=utf8&parseTime=True&loc=Local")

	if err !=nil{
		panic(err)
	}
	db=d
}

func GetDB() *gorm.DB{
	return db
}