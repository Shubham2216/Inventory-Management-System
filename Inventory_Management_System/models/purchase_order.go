package models

import (
    "errors"
    "time"

    "gorm.io/gorm"
)

// PurchaseOrder represents a purchase order entity to restock inventory
type PurchaseOrder struct {
    ID              uint      `gorm:"primaryKey" json:"id"`
    ProductID       uint      `gorm:"not null" json:"product_id"`
    Product         Product   `gorm:"foreignKey:ProductID" json:"-"`
    Quantity        int       `gorm:"not null" json:"quantity"`
    SupplierDetails string    `gorm:"type:varchar(255)" json:"supplier_details"`
    OrderDate       time.Time `gorm:"autoCreateTime" json:"order_date"`
}

// Create inserts a new purchase order
func (p *PurchaseOrder) Create(db *gorm.DB) error {
    return db.Create(p).Error
}

// GetPurchaseOrderByID retrieves a purchase order by ID
func GetPurchaseOrderByID(db *gorm.DB, id uint) (*PurchaseOrder, error) {
    var order PurchaseOrder
    if err := db.Preload("Product").First(&order, id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, nil
        }
        return nil, err
    }
    return &order, nil
}

// Update updates a purchase order
func (p *PurchaseOrder) Update(db *gorm.DB) error {
    return db.Save(p).Error
}

// Delete deletes a purchase order
func (p *PurchaseOrder) Delete(db *gorm.DB) error {
    return db.Delete(p).Error
}

// GetAllPurchaseOrders retrieves all purchase orders
func GetAllPurchaseOrders(db *gorm.DB) ([]PurchaseOrder, error) {
    var orders []PurchaseOrder
    if err := db.Preload("Product").Find(&orders).Error; err != nil {
        return nil, err
    }
    return orders, nil
}
