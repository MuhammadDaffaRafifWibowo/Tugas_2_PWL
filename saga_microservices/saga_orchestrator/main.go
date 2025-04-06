// ========== orchestrator/main.go ==========
package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func callService(url string) bool {
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		log.Printf("Error calling %s: %v", url, err)
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}

func orchestrateOrder(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting Order Saga")
	if !callService("http://localhost:8001/create") {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !callService("http://localhost:8002/pay") {
		callService("http://localhost:8001/cancel")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !callService("http://localhost:8003/ship") {
		callService("http://localhost:8003/cancel")
		callService("http://localhost:8002/refund")
		callService("http://localhost:8001/cancel")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println("Order Saga completed successfully")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "order_completed"})
}

func main() {
	http.HandleFunc("/orchestrate", orchestrateOrder)
	log.Println("Orchestrator running on port 8000")
	http.ListenAndServe(":8000", nil)
}