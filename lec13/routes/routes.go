package routes

import (
	"lec13/controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter ...
func SetupRouter() *gin.Engine {
	r := gin.Default()
	gr1 := r.Group("/api")
	{
		gr1.GET("user", controllers.GetUsers)    //api/user +GET
		gr1.POST("user", controllers.CreateUser) //api/user + POST (.json)
		gr1.GET("user/:id", controllers.GetUserByID)
		gr1.PUT("user/:id", controllers.UpdateUser)
		gr1.DELETE("user/:id", controllers.DeleteUser)
	}

	return r
}
