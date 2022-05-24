package config

import "github.com/jinzhu/gorm"

var DB *gorm.DB

func Init() *gorm.DB {
	db, err := gorm.Open("postgres", "host=postgres port=5432 user=root dbname=todoashrdb password=secret  sslmode=disable")

	if err != nil {
		panic(err.Error())
	}

	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
