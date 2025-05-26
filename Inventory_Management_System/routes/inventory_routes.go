package controllers

import (
	"inventory_management_system/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterInventoryRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	controller := controllers.NewInventoryController(db)

	products := rg.Group("/inventory")
	{
		products.POST("", controller.CreateInventory)
		products.GET("", controller.GetAllInventorys)
		products.GET("/:id", controller.GetInventoryByID)
		products.PUT("/:id", controller.UpdateInventory)
		products.DELETE("/:id", controller.DeleteInventory)
	}
}
