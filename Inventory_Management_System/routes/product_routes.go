package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"inventory_management_system/controllers"
)

// RegisterProductRoutes registers product-related routes
func RegisterProductRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	controller := controllers.NewProductController(db)

	products := rg.Group("/products")
	{
		products.POST("", controller.CreateProduct)
		products.GET("", controller.GetAllProducts)
		products.GET("/:id", controller.GetProductByID)
		products.PUT("/:id", controller.UpdateProduct)
		products.DELETE("/:id", controller.DeleteProduct)
	}
}
