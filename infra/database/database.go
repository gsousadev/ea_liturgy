package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
var Err error

func InitDB() {
	dsn := "root:root@tcp(ea_liturgy-db-1)/easy_liturgy?charset=utf8mb4"
	Db, Err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if Err != nil {
		panic(Err)
	}
}
