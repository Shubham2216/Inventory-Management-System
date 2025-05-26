package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"inventory_management_system/models"
)

type InventoryController struct {
	DB *gorm.DB
}

func NewInventoryController(db *gorm.DB) *InventoryController {
	return &InventoryController{DB: db}
}

func (pc *InventoryController) CreateInventory(c *gin.Context) {
	var input models.Inventory
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.ProductID == 0 || input.Quantity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid purchase order data"})
		return
	}

	if err := input.Create(pc.DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create purchase order"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"purchase_order": input})
}

func (pc *InventoryController) GetInventoryByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid purchase order ID"})
		return
	}

	order, err := models.GetInventoryByProductID(pc.DB, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve purchase order"})
		return
	}

	if order == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Purchase order not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"purchase_order": order})
}

func (pc *InventoryController) GetAllInventorys(c *gin.Context) {
	orders, err := models.GetAllInventory(pc.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve purchase orders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"purchase_orders": orders})
}

func (pc *InventoryController) UpdateInventory(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid purchase order ID"})
		return
	}

	existingOrder, err := models.GetInventoryByProductID(pc.DB, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve purchase order"})
		return
	}
	if existingOrder == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Purchase order not found"})
		return
	}

	var input models.Inventory
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.ProductID == 0 || input.Quantity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid purchase order data"})
		return
	}

	if err := existingOrder.Update(pc.DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update purchase order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"purchase_order": existingOrder})
}

func (pc *InventoryController) DeleteInventory(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid purchase order ID"})
		return
	}

	existingOrder, err := models.DeleteInventoryByID(pc.DB, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve purchase order"})
		return
	}
	if existingOrder == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Purchase order not found"})
		return
	}

	if err := existingOrder.Delete(pc.DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete purchase order"})
		return
	}

	c.Status(http.StatusNoContent)
}
