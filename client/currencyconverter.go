package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-project/internal/models"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "client/index.html")
	})

	http.HandleFunc("/convert", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			decoder := json.NewDecoder(r.Body)
			var request models.ConversionRequest
			err := decoder.Decode(&request)
			if err != nil {
				http.Error(w, "Invalid request", http.StatusBadRequest)
				return
			}
			convertedRequest := models.ConversionRequest{
				FromCurrency: request.FromCurrency,
				Amount:       request.Amount,
				ToCurrency:   request.ToCurrency,
			}

			payload, err := json.Marshal(convertedRequest)
			if err != nil {
				http.Error(w, "Failed to encode request", http.StatusInternalServerError)
				return
			}
			resp, err := http.Post("http://localhost:8080/convert", "application/json", bytes.NewBuffer(payload))
			if err != nil {
				http.Error(w, "Failed to send request to server", http.StatusInternalServerError)
				return
			}

			defer resp.Body.Close()
			var response models.ConversionResponse
			err = json.NewDecoder(resp.Body).Decode(&response)
			if err != nil {
				http.Error(w, "Failed to decode response", http.StatusInternalServerError)
				return
			}
			jsonResponse, err := json.Marshal(response)
			if err != nil {
				http.Error(w, "Failed to encode response", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResponse)
		}
	})

	fmt.Println("Server listening on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
