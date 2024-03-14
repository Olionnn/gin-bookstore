package routes

import (
	"github.com/Olionnn/gin-bookstore/http/controllers"
	"github.com/gin-gonic/gin" // Import the "controllers" package
)

func InitRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id", controllers.GetUsersByID)	
	r.POST("/users", controllers.AddUsers)
	r.PUT("/users/:id", controllers.UpdateUsers)
	r.DELETE("/users/:id", controllers.DeleteUsers)

	r.GET("/login", controllers.Login)
	r.POST("/register", controllers.Register)
	r.GET("/logout", controllers.Logout)


	

	return r
}
