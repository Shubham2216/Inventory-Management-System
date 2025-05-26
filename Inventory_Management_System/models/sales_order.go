package models

import (
    "errors"
    "time"

    "gorm.io/gorm"
)

// SalesOrder represents a sales order entity
type SalesOrder struct {
    ID         uint      `gorm:"primaryKey" json:"id"`
    ProductID  uint      `gorm:"not null" json:"product_id"`
    Product    Product   `gorm:"foreignKey:ProductID" json:"-"`
    Quantity   int       `gorm:"not null" json:"quantity"`
    TotalPrice float64   `gorm:"type:decimal(10,2);not null" json:"total_price"`
    OrderDate  time.Time `gorm:"autoCreateTime" json:"order_date"`
}

// Create inserts a new sales order
func (s *SalesOrder) Create(db *gorm.DB) error {
    return db.Create(s).Error
}

// GetSalesOrderByID retrieves a sales order by ID
func GetSalesOrderByID(db *gorm.DB, id uint) (*SalesOrder, error) {
    var order SalesOrder
    if err := db.Preload("Product").First(&order, id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, nil
        }
        return nil, err
    }
    return &order, nil
}

// Update updates a sales order
func (s *SalesOrder) Update(db *gorm.DB) error {
    return db.Save(s).Error
}

// Delete deletes a sales order
func (s *SalesOrder) Delete(db *gorm.DB) error {
    return db.Delete(s).Error
}

// GetAllSalesOrders retrieves all sales orders
func GetAllSalesOrders(db *gorm.DB) ([]SalesOrder, error) {
    var orders []SalesOrder
    if err := db.Preload("Product").Find(&orders).Error; err != nil {
        return nil, err
    }
    return orders, nil
}
