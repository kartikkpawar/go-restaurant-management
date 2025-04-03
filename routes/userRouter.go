package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/kartikkpawar/go-restaurant-management/controllers"
)

func UserRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/users", controller.GetUsers())
	incommingRoutes.GET("/users/:userId", controller.GetUser())

	incommingRoutes.POST("/users/signup", controller.SignUp())
	incommingRoutes.POST("/users/login", controller.Login())
}
