package utils

import (
    "errors"
    "strings"
    "unicode/utf8"
)

// ValidateProductInput validates basic product fields
func ValidateProductInput(name, description string, price float64, quantity int) error {
    if utf8.RuneCountInString(strings.TrimSpace(name)) == 0 {
        return errors.New("name is required")
    }
    if price <= 0 {
        return errors.New("price must be greater than zero")
    }
    if quantity < 0 {
        return errors.New("quantity cannot be negative")
    }
    if utf8.RuneCountInString(description) > 1000 {
        // Arbitrary length limit for description
        return errors.New("description is too long")
    }
    return nil
}

// ValidateSalesOrderInput validates sales order input fields
func ValidateSalesOrderInput(productID uint, quantity int, totalPrice float64) error {
    if productID == 0 {
        return errors.New("product ID is required")
    }
    if quantity <= 0 {
        return errors.New("quantity must be greater than zero")
    }
    if totalPrice <= 0 {
        return errors.New("total price must be greater than zero")
    }
    return nil
}

// ValidatePurchaseOrderInput validates purchase order input fields
func ValidatePurchaseOrderInput(productID uint, quantity int) error {
    if productID == 0 {
        return errors.New("product ID is required")
    }
    if quantity <= 0 {
        return errors.New("quantity must be greater than zero")
    }
    return nil
}


