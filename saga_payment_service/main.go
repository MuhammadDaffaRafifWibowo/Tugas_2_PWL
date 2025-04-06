// ========== payment-service/main.go ==========
package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func processPayment(w http.ResponseWriter, r *http.Request) {
	log.Println("Payment processed")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "payment_success"})
}

func refundPayment(w http.ResponseWriter, r *http.Request) {
	log.Println("Payment refunded")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "payment_refunded"})
}

func main() {
	http.HandleFunc("/pay", processPayment)
	http.HandleFunc("/refund", refundPayment)
	log.Println("Payment Service running on port 8002")
	http.ListenAndServe(":8002", nil)
}