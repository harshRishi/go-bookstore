package models

import (
	"github.com/harshRishi/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name         string `gorm:"column:name" json:"name"`
	Author       string `json:"author"`
	Publications string `json:"publications"`
}

func init() {
	db = config.GetDb()
	db.AutoMigrate(&Book{})
}
