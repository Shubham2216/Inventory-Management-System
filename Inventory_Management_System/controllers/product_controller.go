package controllers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "inventory_management_system/models"
)

type ProductController struct {
    DB *gorm.DB
}

func NewProductController(db *gorm.DB) *ProductController {
    return &ProductController{DB: db}
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
    var input models.Product
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if input.Name == "" || input.Price <= 0 || input.Quantity < 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product data"})
        return
    }

    if err := input.Create(pc.DB); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"product": input})
}

func (pc *ProductController) GetProductByID(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil || id <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
        return
    }

    product, err := models.GetProductByID(pc.DB, uint(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve product"})
        return
    }
    if product == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"product": product})
}

func (pc *ProductController) GetAllProducts(c *gin.Context) {
    products, err := models.GetAllProducts(pc.DB)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve products"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"products": products})
}

func (pc *ProductController) UpdateProduct(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil || id <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
        return
    }

    existingProduct, err := models.GetProductByID(pc.DB, uint(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve product"})
        return
    }
    if existingProduct == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    var input models.Product
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if input.Name == "" || input.Price <= 0 || input.Quantity < 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product data"})
        return
    }

    existingProduct.Name = input.Name
    existingProduct.Description = input.Description
    existingProduct.Price = input.Price
    existingProduct.Quantity = input.Quantity

    if err := existingProduct.Update(pc.DB); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update product"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"product": existingProduct})
}

func (pc *ProductController) DeleteProduct(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil || id <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
        return
    }

    existingProduct, err := models.GetProductByID(pc.DB, uint(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve product"})
        return
    }
    if existingProduct == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    if err := existingProduct.Delete(pc.DB); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete product"})
        return
    }

    c.Status(http.StatusNoContent)
}
