package controllers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "inventory_management_system/models"
)

type SalesOrderController struct {
    DB *gorm.DB
}

func NewSalesOrderController(db *gorm.DB) *SalesOrderController {
    return &SalesOrderController{DB: db}
}

func (sc *SalesOrderController) CreateSalesOrder(c *gin.Context) {
    var input models.SalesOrder
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if input.ProductID == 0 || input.Quantity <= 0 || input.TotalPrice <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sales order data"})
        return
    }

    // Optionally, you can check product existence here

    if err := input.Create(sc.DB); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create sales order"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"sales_order": input})
}

func (sc *SalesOrderController) GetSalesOrderByID(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil || id <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sales order ID"})
        return
    }

    order, err := models.GetSalesOrderByID(sc.DB, uint(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve sales order"})
        return
    }
    if order == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Sales order not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"sales_order": order})
}

func (sc *SalesOrderController) GetAllSalesOrders(c *gin.Context) {
    orders, err := models.GetAllSalesOrders(sc.DB)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve sales orders"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"sales_orders": orders})
}

func (sc *SalesOrderController) UpdateSalesOrder(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil || id <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sales order ID"})
        return
    }

    existingOrder, err := models.GetSalesOrderByID(sc.DB, uint(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve sales order"})
        return
    }
    if existingOrder == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Sales order not found"})
        return
    }

    var input models.SalesOrder
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if input.ProductID == 0 || input.Quantity <= 0 || input.TotalPrice <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sales order data"})
        return
    }

    existingOrder.ProductID = input.ProductID
    existingOrder.Quantity = input.Quantity
    existingOrder.TotalPrice = input.TotalPrice
    // Optional: Update order date if needed

    if err := existingOrder.Update(sc.DB); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update sales order"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"sales_order": existingOrder})
}

func (sc *SalesOrderController) DeleteSalesOrder(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil || id <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sales order ID"})
        return
    }

    existingOrder, err := models.GetSalesOrderByID(sc.DB, uint(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve sales order"})
        return
    }
    if existingOrder == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Sales order not found"})
        return
    }

    if err := existingOrder.Delete(sc.DB); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete sales order"})
        return
    }

    c.Status(http.StatusNoContent)
}
