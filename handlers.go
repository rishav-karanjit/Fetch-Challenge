package main

import (
	"encoding/json"
	"net/http"
	"github.com/google/uuid"
)

func processReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, "The receipt is invalid", http.StatusBadRequest)
		return
	}

	// Validate receipt
	if validateReceipt(receipt) == true {
		http.Error(w, "The receipt is invalid", http.StatusBadRequest)
		return
	}

	// Generate an ID and save to our in-memory DB (for demo purposes)
	id := uuid.New().String()
	receiptMutex.Lock()
	receiptDB[id] = receipt
	receiptMutex.Unlock()

	resp := ReceiptIDResponse{ID: id}
	json.NewEncoder(w).Encode(resp)
}

func getPoints(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/receipts/") : len(r.URL.Path)-len("/points")]

	receiptMutex.Lock()
	receipt, exists := receiptDB[id]
	receiptMutex.Unlock()

	if !exists {
		http.Error(w, "No receipt found for that id", http.StatusNotFound)
		return
	}

	// Check if points for this receipt are already calculated and stored in pointsDB.
	pointsMutex.Lock()
	points, found := pointsDB[id]
	pointsMutex.Unlock()

	if !found {
		// If not found in the map, calculate them.
		points = calculatePoints(receipt)

		// Store the calculated points in the pointsDB.
		pointsMutex.Lock()
		pointsDB[id] = points
		pointsMutex.Unlock()
	}
	resp := PointsResponse{Points: points}
	json.NewEncoder(w).Encode(resp)
}
