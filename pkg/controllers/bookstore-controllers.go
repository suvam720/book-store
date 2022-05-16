package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suvam720/book-store/pkg/handlrs"
	"github.com/suvam720/book-store/pkg/models"
	"github.com/suvam720/book-store/pkg/utils"
)

var NewBook models.Book

func GetBooks(c *gin.Context) {
	books := handlrs.GetAllBooks()
	c.JSON(http.StatusOK, gin.H{"data": books})
}

func GetBookById(c *gin.Context) {
	id := c.Param("id")
	books, err := handlrs.GetBookById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
	}
	c.JSON(http.StatusOK, gin.H{"book": books})
}

func CreateBook(c *gin.Context) {
	CreateBook := &models.Book{}
	utils.ParseBody(c, CreateBook)

	book := handlrs.CreateBook(CreateBook)

	c.JSON(http.StatusOK, gin.H{"created": book})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	bookDetails := handlrs.DeleteBook(id)

	c.JSON(http.StatusOK, gin.H{"data": bookDetails.DeletedAt.Time})
}

func UpdateBook(c *gin.Context) {
	var updatedBook = &models.Book{}
	utils.ParseBody(c, updatedBook)
	id := c.Param("id")

	bookDetails := handlrs.UpdateBook(id, updatedBook)

	c.JSON(http.StatusOK, gin.H{"data": bookDetails})
}
