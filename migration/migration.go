package migration

import (
	"golang_todo_app_ashrtech/model"

	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Todo{})
	db.AutoMigrate(&model.User{})
}
