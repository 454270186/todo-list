package dal

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func newDB() *gorm.DB {
	if DB != nil {
		return DB
	}

	DB = initDB()
	
	return DB
}

func initDB() *gorm.DB {
	dsn := "xiaofei:2021110003@tcp(127.0.0.1:3306)/todolist?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}

	return db
}