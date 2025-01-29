package routes

import (
	"go-restapi/controllers"

	"github.com/gin-gonic/gin"
)

func BookRoutes(r *gin.Engine) {
	r.POST("/books", controllers.CreateBook)
	r.GET("/books", controllers.GetBooks)
	r.GET("/books/:id", controllers.GetBookByID)
	r.PUT("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
}
