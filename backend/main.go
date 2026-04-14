package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	resp := Response{
		Message: "Hello from Go backend 🚀",
	}

	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/api/hello", helloHandler)

	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}