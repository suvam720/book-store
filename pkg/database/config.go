package database

import (
	"github.com/suvam720/book-store/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect(dsn string) error {

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}
	db.AutoMigrate(&models.Book{})
	Db = db

	return nil
}
