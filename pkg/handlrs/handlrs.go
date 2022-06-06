package handlrs

import (
	"github.com/suvam720/book-store/pkg/database"
	"github.com/suvam720/book-store/pkg/models"
	"gorm.io/gorm"
)

func CreateBook(b *models.Book) *models.Book {
	database.Db.Create(&b)

	return b
}

func GetAllBooks() []models.Book {
	var Books []models.Book
	database.Db.Find(&Books)

	return Books
}

func GetBookById(Id string) (*models.Book, *gorm.DB) {
	var getBook models.Book
	db := database.Db.Where("ID=?", Id).Find(&getBook)

	return &getBook, db
}

func DeleteBook(Id string) models.Book {
	var Book models.Book
	database.Db.Where("ID=?", Id).Delete(Book)

	return Book
}

func UpdateBook(Id string, b *models.Book) models.Book {
	bookDetails, db := GetBookById(Id)

	bookDetails.Author = b.Author
	bookDetails.Title = b.Title
	bookDetails.Publisher = b.Publisher
	bookDetails.Rating = b.Rating
	bookDetails.PublishedDate = b.PublishedDate
	db.Save(&bookDetails)

	return *bookDetails
}
