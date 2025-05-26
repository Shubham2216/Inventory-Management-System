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

func TestCreateInventory(t *testing.T) {
	dsn := migrations.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	product := models.Product{Name: "ProdInv", Price: 45, Quantity: 20}
	err = product.Create(db)
	assert.NoError(t, err)

	inventory := models.Inventory{
		ProductID: product.ID,
		Quantity:  20,
	}
	err = inventory.Create(db)
	assert.NoError(t, err)
	assert.NotZero(t, inventory.ID)
}

func TestGetInventoryByProductID(t *testing.T) {
	dsn := migrations.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	product := models.Product{Name: "ProdInv2", Price: 60, Quantity: 25}
	err = product.Create(db)
	assert.NoError(t, err)

	inventory := models.Inventory{
		ProductID: product.ID,
		Quantity:  25,
	}
	err = inventory.Create(db)
	assert.NoError(t, err)

	fetched, err := models.GetInventoryByProductID(db, product.ID)
	assert.NoError(t, err)
	assert.NotNil(t, fetched)
	assert.Equal(t, 25, fetched.Quantity)
}

func TestUpdateInventory(t *testing.T) {
	dsn := migrations.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	product := models.Product{Name: "ProdInv3", Price: 55, Quantity: 50}
	err = product.Create(db)
	assert.NoError(t, err)

	inventory := models.Inventory{
		ProductID: product.ID,
		Quantity:  10,
	}
	err = inventory.Create(db)
	assert.NoError(t, err)

	inventory.Quantity = 40
	err = inventory.Update(db)
	assert.NoError(t, err)

	updated, err := models.GetInventoryByProductID(db, product.ID)
	assert.NoError(t, err)
	assert.Equal(t, 40, updated.Quantity)
}

func TestDeleteInventory(t *testing.T) {
	dsn := migrations.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	product := models.Product{Name: "ProdInvDel", Price: 35, Quantity: 60}
	err = product.Create(db)
	assert.NoError(t, err)

	inventory := models.Inventory{
		ProductID: product.ID,
		Quantity:  30,
	}
	err = inventory.Create(db)
	assert.NoError(t, err)

	err = inventory.Delete(db)
	assert.NoError(t, err)

	deleted, err := models.GetInventoryByProductID(db, product.ID)
	assert.NoError(t, err)
	assert.Nil(t, deleted)
}
