package routes

import (
	"github.com/Olionnn/gin-bookstore/http/controllers"
	"github.com/Olionnn/gin-bookstore/http/middleware"
	"github.com/gin-gonic/gin" // Import the "controllers" package
)

func InitRoutes() *gin.Engine {
	r := gin.Default()

	public := r.Group("/api")
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	
	private := r.Group("/api")
	private.Use(middleware.AuthMiddleware)
	
	private.GET("/users", controllers.FindAllUsers)
	
	private.GET("/books", controllers.FindAllBooks)
	private.POST("/books", controllers.AddBook)

	return r
}
