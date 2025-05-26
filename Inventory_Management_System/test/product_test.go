package tests

import (
	"gorm.io/driver/postgres"
	"inventory_management_system/migrations"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"inventory_management_system/models"
)

func TestCreateProduct(t *testing.T) {
	dsn := migrations.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	product := models.Product{
		Name:        "Test Product",
		Description: "Sample product description",
		Price:       99.99,
		Quantity:    10,
	}

	err = product.Create(db)
	assert.NoError(t, err)
	assert.NotZero(t, product.ID)
}

func TestGetProductByID(t *testing.T) {
	dsn := migrations.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	product := models.Product{
		Name:        "TestGet",
		Description: "Get by ID test",
		Price:       50,
		Quantity:    5,
	}
	err = product.Create(db)
	assert.NoError(t, err)

	fetched, err := models.GetProductByID(db, product.ID)
	assert.NoError(t, err)
	assert.NotNil(t, fetched)
	assert.Equal(t, product.Name, fetched.Name)
}

func TestUpdateProduct(t *testing.T) {
	dsn := migrations.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	product := models.Product{
		Name:        "Old Name",
		Description: "Old Desc",
		Price:       10,
		Quantity:    1,
	}
	err = product.Create(db)
	assert.NoError(t, err)

	product.Name = "New Name"
	product.Description = "New Desc"
	product.Price = 20
	product.Quantity = 5

	err = product.Update(db)
	assert.NoError(t, err)

	updated, err := models.GetProductByID(db, product.ID)
	assert.NoError(t, err)
	assert.Equal(t, "New Name", updated.Name)
	assert.Equal(t, "New Desc", updated.Description)
	assert.Equal(t, 20.0, updated.Price)
	assert.Equal(t, 5, updated.Quantity)
}

func TestDeleteProduct(t *testing.T) {
	dsn := migrations.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	product := models.Product{
		Name:        "Delete Test",
		Description: "",
		Price:       5,
		Quantity:    2,
	}
	err = product.Create(db)
	assert.NoError(t, err)

	err = product.Delete(db)
	assert.NoError(t, err)

	deleted, err := models.GetProductByID(db, product.ID)
	assert.NoError(t, err)
	assert.Nil(t, deleted)
}
