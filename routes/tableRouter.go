package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/kartikkpawar/go-restaurant-management/controllers"
)

func TableRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/tables", controller.GetTables())
	incomingRoutes.GET("/tables/:tableId", controller.GetTable())
	incomingRoutes.POST("/tables", controller.CreateTable())
	incomingRoutes.PATCH("/tables/:tableId", controller.UpdateTable())
}
