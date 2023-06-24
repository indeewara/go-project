package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-project/internal/models"
)

func ConvertHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var request models.ConversionRequest
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	client := http.Client{}
	resp, err := client.Get("https://api.exchangerate.host/latest")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var rates map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&rates)
	if err != nil {
		http.Error(w, "Failed to fetch exchange rates", http.StatusInternalServerError)
		return
	}

	FromCurrencyRate := rates["rates"].(map[string]interface{})[request.FromCurrency].(float64)
	ToCurrencyRate := rates["rates"].(map[string]interface{})[request.ToCurrency].(float64)

	convertedAmount := request.Amount * (ToCurrencyRate / FromCurrencyRate)

	response := models.ConversionResponse{
		Amount:   convertedAmount,
		Currency: request.ToCurrency,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
