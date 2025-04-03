package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/kartikkpawar/go-restaurant-management/controllers"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controller.GetOrderItems())
	incomingRoutes.GET("/orderItems/:orderItemId", controller.GetOrderItem())
	incomingRoutes.POST("/orderItems", controller.CreateOrder())
	incomingRoutes.PATCH("/orderItems/:orderItemId", controller.UpdateOrderItem())

	incomingRoutes.GET("/orderItems-order/:orderItemId", controller.GetOrderItemsbyOrder())

}
