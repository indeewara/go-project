package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-project/internal/models"
	"github.com/patrickmn/go-cache"
)

type App struct {
	Cache *cache.Cache
}

func (a *App) ConvertHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var request models.ConversionRequest
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Check if the data is already cached
	cachedData, found := a.Cache.Get("exchange_rates")
	if found {
		rates, ok := cachedData.(map[string]interface{})
		if !ok {
			http.Error(w, "Failed to fetch exchange rates from cache", http.StatusInternalServerError)
			return
		}
		processConversion(w, request, rates)
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

	a.Cache.Set("exchange_rates", rates, 3*time.Hour)

	processConversion(w, request, rates)
}

func processConversion(w http.ResponseWriter, request models.ConversionRequest, rates map[string]interface{}) {
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
