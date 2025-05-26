package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"inventory_management_system/controllers"
)

func RegisterPurchaseOrderRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	controller := controllers.NewPurchaseOrderController(db)

	products := rg.Group("/PurchaseOrder")
	{
		products.POST("", controller.CreatePurchaseOrder)
		products.GET("", controller.GetAllPurchaseOrders)
		products.GET("/:id", controller.GetPurchaseOrderByID)
		products.PUT("/:id", controller.UpdatePurchaseOrder)
		products.DELETE("/:id", controller.DeletePurchaseOrder)
	}
}
