package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	go1inch "github.com/jon4hz/go-1inch"
)

func main() {
	client := go1inch.NewClient()
	tokenResponse, _, err := client.Tokens(context.Background(), "eth")
	if err != nil {
		log.Fatal("Error while getting an approve calldata ", err)
	}

	http.HandleFunc("/Token", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(tokenResponse)
	})
	http.ListenAndServe(":8080", nil)

}
