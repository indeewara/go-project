package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-project/internal/handlers"
	"github.com/go-project/internal/models"
	"github.com/patrickmn/go-cache"
)

func TestConvertHandler(t *testing.T) {
	c := cache.New(cache.NoExpiration, cache.NoExpiration)
	app := handlers.App{
		Cache: c,
	}

	requestBody := models.ConversionRequest{
		FromCurrency: "USD",
		ToCurrency:   "EUR",
		Amount:       100,
	}
	jsonData, _ := json.Marshal(requestBody)

	req, err := http.NewRequest("POST", "/convert", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.ConvertHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}
	var response models.ConversionResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("failed to unmarshal response body: %v", err)
	}
}
