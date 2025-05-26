package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"inventory_management_system/migrations"
	controllers "inventory_management_system/routes"
	"log"

	"inventory_management_system/models"
)

func main() {
	dsn := migrations.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate models
	err = db.AutoMigrate(&models.Product{}, &models.SalesOrder{}, &models.PurchaseOrder{}, &models.Inventory{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	router := gin.Default()

	apiGroup := router.Group("/api")
	{
		controllers.RegisterProductRoutes(apiGroup, db)
		controllers.RegisterInventoryRoutes(apiGroup, db)
		controllers.RegisterSalesOrderRoutes(apiGroup, db)
		controllers.RegisterPurchaseOrderRoutes(apiGroup, db)
	}

	port := "8080"
	fmt.Printf("Starting server on port %s...\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
