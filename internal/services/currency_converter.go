package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ConversionRequest struct {
	FromCurrency string  `json:"fromCurrency"`
	Amount       float64 `json:"amount"`
	ToCurrency   string  `json:"toCurrency"`
}

type ConversionResponse struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

func main() {
	// Create the request payload
	request := ConversionRequest{
		FromCurrency: "USD",
		Amount:       234.0,
		ToCurrency:   "LKR",
	}

	// Convert the request payload to JSON
	payload, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}

	// Make a POST request to the /convert endpoint
	resp, err := http.Post("http://localhost:8080/convert", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Decode the response
	var response ConversionResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Fatal(err)
	}

	// Print the converted amount
	fmt.Println("Converted Amount:", response.Amount, response.Currency)
}
