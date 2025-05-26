package tests

import (
	"gorm.io/driver/postgres"
	"inventory_management_system/migrations"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"inventory_management_system/models"
)

func TestCreatePurchaseOrder(t *testing.T) {
	dsn := migrations.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	product := models.Product{Name: "ProdPO", Price: 25, Quantity: 50}
	err = product.Create(db)
	assert.NoError(t, err)

	order := models.PurchaseOrder{
		ProductID:       product.ID,
		Quantity:        10,
		SupplierDetails: "Supplier X",
		OrderDate:       time.Now(),
	}
	err = order.Create(db)
	assert.NoError(t, err)
	assert.NotZero(t, order.ID)
}

func TestGetPurchaseOrderByID(t *testing.T) {
	dsn := migrations.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	product := models.Product{Name: "ProdPO2", Price: 40, Quantity: 60}
	err = product.Create(db)
	assert.NoError(t, err)

	order := models.PurchaseOrder{
		ProductID:       product.ID,
		Quantity:        7,
		SupplierDetails: "Supplier Y",
		OrderDate:       time.Now(),
	}
	err = order.Create(db)
	assert.NoError(t, err)

	fetched, err := models.GetPurchaseOrderByID(db, order.ID)
	assert.NoError(t, err)
	assert.NotNil(t, fetched)
	assert.Equal(t, order.Quantity, fetched.Quantity)
}

func TestUpdatePurchaseOrder(t *testing.T) {
	dsn := migrations.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	product := models.Product{Name: "ProdPO3", Price: 35, Quantity: 40}
	err = product.Create(db)
	assert.NoError(t, err)

	order := models.PurchaseOrder{
		ProductID:       product.ID,
		Quantity:        3,
		SupplierDetails: "Supplier Z",
	}
	err = order.Create(db)
	assert.NoError(t, err)

	order.Quantity = 8
	order.SupplierDetails = "Supplier ZZ"

	err = order.Update(db)
	assert.NoError(t, err)

	updated, err := models.GetPurchaseOrderByID(db, order.ID)
	assert.NoError(t, err)
	assert.Equal(t, 8, updated.Quantity)
	assert.Equal(t, "Supplier ZZ", updated.SupplierDetails)
}

func TestDeletePurchaseOrder(t *testing.T) {
	dsn := migrations.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	product := models.Product{Name: "ProdDelPO", Price: 50, Quantity: 30}
	err = product.Create(db)
	assert.NoError(t, err)

	order := models.PurchaseOrder{
		ProductID:       product.ID,
		Quantity:        5,
		SupplierDetails: "Supplier Del",
	}
	err = order.Create(db)
	assert.NoError(t, err)

	err = order.Delete(db)
	assert.NoError(t, err)

	deleted, err := models.GetPurchaseOrderByID(db, order.ID)
	assert.NoError(t, err)
	assert.Nil(t, deleted)
}
