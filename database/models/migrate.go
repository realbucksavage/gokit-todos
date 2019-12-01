package models

import "github.com/jinzhu/gorm"

func Automigrate(db *gorm.DB) {
	db.AutoMigrate(&Todo{})
}
