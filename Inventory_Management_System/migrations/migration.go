package migrations

import (
	"log"

	"gorm.io/gorm"

	"inventory_management_system/models"
)

// Migrate runs the database migrations for the inventory management system
func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Product{},
		&models.SalesOrder{},
		&models.PurchaseOrder{},
		&models.Inventory{},
	)
	if err != nil {
		log.Printf("Migration failed: %v", err)
		return err
	}
	log.Println("Migration completed successfully")
	return nil
}

func GetDSN() string {
	return "postgres://postgres:Shubham@123@localhost:5432/postgres?sslmode=disable"
}
