package models

import (
    "errors"

    "gorm.io/gorm"
)

// Product represents the product entity in the database
type Product struct {
    ID          uint    `gorm:"primaryKey" json:"id"`
    Name        string  `gorm:"type:varchar(255);not null" json:"name"`
    Description string  `gorm:"type:text" json:"description"`
    Price       float64 `gorm:"type:decimal(10,2);not null" json:"price"`
    Quantity    int     `gorm:"not null" json:"quantity"`
}

// Create inserts a new product record in the database
func (p *Product) Create(db *gorm.DB) error {
    return db.Create(p).Error
}

// GetByID retrieves a product by its ID
func GetProductByID(db *gorm.DB, id uint) (*Product, error) {
    var product Product
    if err := db.First(&product, id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, nil
        }
        return nil, err
    }
    return &product, nil
}

// Update updates the product details in the database
func (p *Product) Update(db *gorm.DB) error {
    return db.Save(p).Error
}

// Delete removes a product record from the database
func (p *Product) Delete(db *gorm.DB) error {
    return db.Delete(p).Error
}

// GetAllProducts retrieves all products from the database
func GetAllProducts(db *gorm.DB) ([]Product, error) {
    var products []Product
    if err := db.Find(&products).Error; err != nil {
        return nil, err
    }
    return products, nil
}

