// ========== shipping-service/main.go ==========
package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

func shipOrder(w http.ResponseWriter, r *http.Request) {
	log.Println("Shipping process started")
	if rand.Intn(2) == 0 {
		log.Println("Shipping failed")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"status": "shipping_failed"})
		return
	}
	log.Println("Shipping successful")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "shipping_success"})
}

func cancelShipping(w http.ResponseWriter, r *http.Request) {
	log.Println("Shipping cancelled")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "shipping_cancelled"})
}

func main() {
	http.HandleFunc("/ship", shipOrder)
	http.HandleFunc("/cancel", cancelShipping)
	log.Println("Shipping Service running on port 8003")
	http.ListenAndServe(":8003", nil)
}