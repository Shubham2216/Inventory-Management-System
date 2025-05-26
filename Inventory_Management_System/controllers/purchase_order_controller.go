package controllers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "inventory_management_system/models"
)

type PurchaseOrderController struct {
    DB *gorm.DB
}

func NewPurchaseOrderController(db *gorm.DB) *PurchaseOrderController {
    return &PurchaseOrderController{DB: db}
}

func (pc *PurchaseOrderController) CreatePurchaseOrder(c *gin.Context) {
    var input models.PurchaseOrder
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

func (pc *PurchaseOrderController) GetPurchaseOrderByID(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil || id <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid purchase order ID"})
        return
    }

    order, err := models.GetPurchaseOrderByID(pc.DB, uint(id))
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

func (pc *PurchaseOrderController) GetAllPurchaseOrders(c *gin.Context) {
    orders, err := models.GetAllPurchaseOrders(pc.DB)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve purchase orders"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"purchase_orders": orders})
}

func (pc *PurchaseOrderController) UpdatePurchaseOrder(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil || id <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid purchase order ID"})
        return
    }

    existingOrder, err := models.GetPurchaseOrderByID(pc.DB, uint(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve purchase order"})
        return
    }
    if existingOrder == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Purchase order not found"})
        return
    }

    var input models.PurchaseOrder
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if input.ProductID == 0 || input.Quantity <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid purchase order data"})
        return
    }

    existingOrder.ProductID = input.ProductID
    existingOrder.Quantity = input.Quantity
    existingOrder.SupplierDetails = input.SupplierDetails

    if err := existingOrder.Update(pc.DB); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update purchase order"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"purchase_order": existingOrder})
}

func (pc *PurchaseOrderController) DeletePurchaseOrder(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil || id <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid purchase order ID"})
        return
    }

    existingOrder, err := models.GetPurchaseOrderByID(pc.DB, uint(id))
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
