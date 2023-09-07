package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func factorialHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Input int `json:"input"`
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		log.Println("invalid JSON request:", err)
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		return
	}

	result := factorial(request.Input)

	response := struct {
		Result int `json:"result"`
	}{
		Result: result,
	}

	log.Println("called /factorial, input:", request.Input, "output:", result)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/factorial", factorialHandler)

	port := 9000
	fmt.Printf("Server is listening on port %d...\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
