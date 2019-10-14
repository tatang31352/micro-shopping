package repository

import "github.com/jinzhu/gorm"

type User struct {
	Db *gorm.DB
}
