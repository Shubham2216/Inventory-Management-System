package models

import (
	"errors"

	"gorm.io/gorm"
)

// Inventory represents the current inventory level of a product
type Inventory struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	ProductID uint    `gorm:"unique;not null" json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID" json:"-"`
	Quantity  int     `gorm:"not null" json:"quantity"`
}

// Create inserts a new inventory record
func (i *Inventory) Create(db *gorm.DB) error {
	return db.Create(i).Error
}

// GetInventoryByProductID retrieves inventory by product ID
func GetInventoryByProductID(db *gorm.DB, productID uint) (*Inventory, error) {
	var inventory Inventory
	if err := db.Where("product_id = ?", productID).First(&inventory).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &inventory, nil
}

// GetInventoryByProductID retrieves inventory by product ID
func DeleteInventoryByID(db *gorm.DB, productID uint) (*Inventory, error) {
	var inventory Inventory
	if err := db.Where("product_id = ?", productID).First(&inventory).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	db.Delete(&inventory)
	return &inventory, nil
}

// Update updates the inventory record
func (i *Inventory) Update(db *gorm.DB) error {
	return db.Save(i).Error
}

// Delete deletes the inventory record
func (i *Inventory) Delete(db *gorm.DB) error {
	return db.Delete(i).Error
}

// GetAllInventory retrieves all inventory records
func GetAllInventory(db *gorm.DB) ([]Inventory, error) {
	var inventory []Inventory
	if err := db.Preload("Product").Find(&inventory).Error; err != nil {
		return nil, err
	}
	return inventory, nil
}
