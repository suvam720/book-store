package database

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/suvam720/book-store/pkg/models"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
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

//var db *gorm.DB

// func Connect() {
// 	EnvKey, _ := godotenv.Read(".env")
// 	dsn := EnvKey["DB_STRING"]
// 	fmt.Printf("dsn: %v\n", dsn)
// 	d, err := gorm.Open(mysql.Open(dsn))
// 	if err != nil {
// 		fmt.Println(err)
// 		panic("failed to connect database")
// 	}
// 	fmt.Printf("d:%v", d)
// 	db = d
// }

// func GetDB() *gorm.DB {
// 	return db
// }

// var Db *gorm.DB

// func Connection() {
// 	Connect()
// 	Db = GetDB()
// 	Db.AutoMigrate(&models.Book{})
// }
