// ========== order-service/main.go ==========
package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Order struct {
	ID     string `json:"id"`
	Amount int    `json:"amount"`
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	log.Println("Order created")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "order_created"})
}

func cancelOrder(w http.ResponseWriter, r *http.Request) {
	log.Println("Order cancelled")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "order_cancelled"})
}

func main() {
	http.HandleFunc("/create", createOrder)
	http.HandleFunc("/cancel", cancelOrder)
	log.Println("Order Service running on port 8001")
	http.ListenAndServe(":8001", nil)
}