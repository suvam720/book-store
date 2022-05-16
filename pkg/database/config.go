package database

import (
	//_ "github.com/jinzu/gorm/dislects/mysql"
	"github.com/joho/godotenv"
	"github.com/suvam720/book-store/pkg/models"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	EnvKey, _ := godotenv.Read(".env")
	dsn := EnvKey["DB_STRING"]
	d, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("failed to connect database")
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}

var Db *gorm.DB

func init() {
	Connect()
	Db = GetDB()
	Db.AutoMigrate(&models.Book{})
}
