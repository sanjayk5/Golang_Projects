package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var Books []Book

func main() {
	Books = []Book{
		{ID: "1", Title: "Python Tricks", Author: "Dan Bader"},
		{ID: "2", Title: "Coding Interview Patterns", Author: "Alex Xu"},
		{ID: "3", Title: "Go in Action", Author: "Dmitry Nesterov"},
		{ID: "4", Title: "Clean Code", Author: "Robert C. Martin"},
		{ID: "5", Title: "Effective Java", Author: "Joshua Bloch"},
	}

	router := gin.Default() // Initializing a Gin router
	// Assigning handler's functions to different endpoint routes
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByID)
	router.POST("/books", createBook)
	router.DELETE("/books/:id", deleteBook)

	router.Run("localhost:8080") // Attaching the router to server and starting the server
}

func createBook(c *gin.Context) {
	var newBook Book
	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Failed to create a new book"})
		return
	}

	Books = append(Books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Books)
}

func getBookByID(c *gin.Context) {
	for _, book := range Books {
		if book.ID == c.Param("id") {
			c.IndentedJSON(http.StatusOK, book)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Requested book is not found"})
}

func deleteBook(c *gin.Context) {
	for index, book := range Books {
		if book.ID == c.Param("id") {
			Books = append(Books[:index], Books[index+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Book is deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book is not found"})
}
