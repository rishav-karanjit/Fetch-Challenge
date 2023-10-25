package main

import (
	"regexp"
	"time"
	"fmt"
)

func validateReceipt(receipt Receipt) bool {
	// Retailer pattern ("^\\S+$") is not validated because example on
	// readme ("retailer": "M&M Corner Market") does not match its validation

	if receipt.Retailer == "" {
		return true
	}

	// Date Validation
	_, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err != nil {
		fmt.Println("0")
		return true
	}

	// Time validation
	_, err = time.Parse("15:04", receipt.PurchaseTime)
	if err != nil {
		fmt.Println("1")
		return true
	}

	// Ensure there is at least one item
	if len(receipt.Items) < 1 {
		fmt.Println("2")
		return true
	}

	// Total amount format validation
	if !regexp.MustCompile(`^\d+\.\d{2}$`).MatchString(receipt.Total) {
		fmt.Println("3")
		return true
	}

	// Validate each item's shortDescription and price format
    for _, item := range receipt.Items {
        if !regexp.MustCompile(`^[\w\s\-]+$`).MatchString(item.ShortDescription) {
            return true
        }
        
        if !regexp.MustCompile(`^\d+\.\d{2}$`).MatchString(item.Price) {
            return true
        }
    }

	return false
}