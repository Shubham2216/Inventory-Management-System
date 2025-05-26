package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"inventory_management_system/controllers"
)

func RegisterSalesOrderRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	controller := controllers.NewSalesOrderController(db)

	products := rg.Group("/SalesOrder")
	{
		products.POST("", controller.CreateSalesOrder)
		products.GET("", controller.GetAllSalesOrders)
		products.GET("/:id", controller.GetSalesOrderByID)
		products.PUT("/:id", controller.UpdateSalesOrder)
		products.DELETE("/:id", controller.DeleteSalesOrder)
	}
}
