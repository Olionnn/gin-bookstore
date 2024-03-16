package routes

import (
	"github.com/Olionnn/gin-bookstore/http/controllers"
	"github.com/gin-gonic/gin" // Import the "controllers" package
)

func InitRoutes() *gin.Engine {
	r := gin.Default()

	public := r.Group("/api")
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	// r.GET("/users", controllers.GetUsers)
	// r.GET("/users/:id", controllers.GetUsersByID)
	// r.POST("/users", controllers.AddUsers)
	// r.PUT("/users/:id", controllers.UpdateUsers)
	// r.DELETE("/users/:id", controllers.DeleteUsers)

	return r
}
