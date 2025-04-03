package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/kartikkpawar/go-restaurant-management/controllers"
)

func MenuRoutes(incomeingRoutes *gin.Engine) {

	incomeingRoutes.GET("/menus", controller.GetMenus())
	incomeingRoutes.GET("/menus/:menuId", controller.GetMenu())
	incomeingRoutes.POST("/menus", controller.CreateMenu())
	incomeingRoutes.PATCH("/menus/:menuId", controller.UpdateMenu())

}
