package controllers

import (
	"bookweb/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Create book schema for validate
type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

// GET /books
// GET all books
func FindBooks(ctx *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	ctx.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

// POST /books
// Create new book
func CreateBook(ctx *gin.Context) {
	// valid input
	var input CreateBookInput
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// Create book
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)
	ctx.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

// GET /books/:id
// Find a book
func FindBook(ctx *gin.Context) {
	var book models.Book
	id := ctx.Param("id")
	err := models.DB.Where("id = ?", id).First(&book).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Content not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": book})

}

// Update a book
func UpdateBook(ctx *gin.Context) {
	var book models.Book
	id := ctx.Param("id")
	err := models.DB.Where("id = ?", id).First(&book).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Content not found",
		})
	}
	// Validate input
	var input UpdateBookInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	models.DB.Model(&book).Updates(input)
	ctx.JSON(http.StatusOK, gin.H{
		"data": book,
	})

}

// DELETE /books/:id
// Delete a book
func DeleteBook(ctx *gin.Context) {
	// Get model if exist
	var book models.Book
	id := ctx.Param("id")
	if err := models.DB.Where("id = ?", id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
	}
	models.DB.Delete(&book)
	ctx.JSON(http.StatusOK,gin.H{
		"data":true,
	})
}
