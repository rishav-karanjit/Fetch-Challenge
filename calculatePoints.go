package main

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func calculatePoints(receipt Receipt) int64 {
	return retailerPoints(receipt.Retailer) +
		totalAmountPoints(receipt.Total) +
		itemPoints(receipt.Items) +
		purchaseDatePoints(receipt.PurchaseDate) +
		purchaseTimePoints(receipt.PurchaseTime)
}

func retailerPoints(retailer string) int64 {
	points := int64(0)
	for _, char := range retailer {
		if regexp.MustCompile(`^[a-zA-Z0-9]$`).MatchString(string(char)) {
			points++
		}
	}
	return points
}

func totalAmountPoints(total string) int64 {
	points := int64(0)
	if strings.HasSuffix(total, ".00") {
		points += 50
	}
	if totalValue, err := strconv.ParseFloat(total, 64); err == nil {
		if math.Mod(totalValue, 0.25) == 0 {
			points += 25
		}
	}
	return points
}

func itemPoints(items []Item) int64 {
	points := int64(0)
	points += int64(len(items) / 2 * 5)
	for _, item := range items {
		descLen := len(strings.TrimSpace(item.ShortDescription))
		if descLen%3 == 0 {
			if itemPrice, err := strconv.ParseFloat(item.Price, 64); err == nil {
				points += int64(math.Ceil(itemPrice * 0.2))
			}
		}
	}
	return points
}

func purchaseDatePoints(purchaseDate string) int64 {
	points := int64(0)
	if parsedDate, err := time.Parse("2006-01-02", purchaseDate); err == nil {
		if parsedDate.Day()%2 != 0 {
			points += 6
		}
	}
	return points
}

func purchaseTimePoints(purchaseTime string) int64 {
	points := int64(0)
	if parsedTime, err := time.Parse("15:04", purchaseTime); err == nil {
		if parsedTime.Hour() >= 14 && parsedTime.Hour() <= 16 {
			points += 10
		}
	}
	return points
}
