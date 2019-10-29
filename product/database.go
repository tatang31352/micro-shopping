package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func CreateConnection()(*gorm.DB,error)  {
	host := "127.0.0.1"
	user := "root"
	dbName := "shopping"
	password := "test"
	return gorm.Open("mysql",fmt.Sprintf(
		"%s:%s@tcp(%s:3307)/%s?charset=utf8&parseTime=True&loc=Local",
		user,password,host,dbName,
	),
	)
}
