package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func CreateConnection(dbconf map[string]interface{})(*gorm.DB,error)  {
	host := dbconf["host"]
	port := dbconf["port"]
	user := dbconf["user"]
	dbName := dbconf["database"]
	password := dbconf["password"]
	return gorm.Open("mysql",fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,password,host,port,dbName,
	),
	)
}
