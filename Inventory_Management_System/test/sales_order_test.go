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

func TestCreateSalesOrder(t *testing.T) {
	dsn := migrations.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	product := models.Product{Name: "TestProd", Price: 20, Quantity: 100}
	err = product.Create(db)
	assert.NoError(t, err)

	order := models.SalesOrder{
		ProductID:  product.ID,
		Quantity:   3,
		TotalPrice: 60,
		OrderDate:  time.Now(),
	}
	err = order.Create(db)
	assert.NoError(t, err)
	assert.NotZero(t, order.ID)
}

func TestGetSalesOrderByID(t *testing.T) {
	dsn := migrations.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	product := models.Product{Name: "Prod2", Price: 10, Quantity: 50}
	err = product.Create(db)
	assert.NoError(t, err)

	order := models.SalesOrder{
		ProductID:  product.ID,
		Quantity:   5,
		TotalPrice: 50,
		OrderDate:  time.Now(),
	}
	err = order.Create(db)
	assert.NoError(t, err)

	fetched, err := models.GetSalesOrderByID(db, order.ID)
	assert.NoError(t, err)
	assert.NotNil(t, fetched)
	assert.Equal(t, order.Quantity, fetched.Quantity)
}

func TestUpdateSalesOrder(t *testing.T) {
	dsn := migrations.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	product := models.Product{Name: "Prod3", Price: 30, Quantity: 70}
	err = product.Create(db)
	assert.NoError(t, err)

	order := models.SalesOrder{
		ProductID:  product.ID,
		Quantity:   2,
		TotalPrice: 60,
	}
	err = order.Create(db)
	assert.NoError(t, err)

	order.Quantity = 4
	order.TotalPrice = 120

	err = order.Update(db)
	assert.NoError(t, err)

	updated, err := models.GetSalesOrderByID(db, order.ID)
	assert.NoError(t, err)
	assert.Equal(t, 4, updated.Quantity)
	assert.Equal(t, 120.0, updated.TotalPrice)
}

func TestDeleteSalesOrder(t *testing.T) {
	dsn := migrations.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	product := models.Product{Name: "ProdDel", Price: 20, Quantity: 40}
	err = product.Create(db)
	assert.NoError(t, err)

	order := models.SalesOrder{
		ProductID:  product.ID,
		Quantity:   1,
		TotalPrice: 20,
	}
	err = order.Create(db)
	assert.NoError(t, err)

	err = order.Delete(db)
	assert.NoError(t, err)

	deleted, err := models.GetSalesOrderByID(db, order.ID)
	assert.NoError(t, err)
	assert.Nil(t, deleted)
}
