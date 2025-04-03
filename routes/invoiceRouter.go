package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/kartikkpawar/go-restaurant-management/controllers"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/invoices", controller.GetInvoices())
	incomingRoutes.GET("/invoices/:invoiceId", controller.GetInvoice())
	incomingRoutes.POST("/invoices", controller.CreateInvoice())
	incomingRoutes.PATCH("/invoices/:invoiceId", controller.UpdateInvoice())

}
