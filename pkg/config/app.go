package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	var err error
	db, err = gorm.Open("mysql", "<username>:<password>@tcp(127.0.0.1:3306)/<tableName>?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
