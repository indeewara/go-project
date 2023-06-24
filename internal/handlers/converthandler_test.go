package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-project/internal/models"
)

func TestConvertHandler(t *testing.T) {
	request := models.ConversionRequest{
		FromCurrency: "USD",
		Amount:       100,
		ToCurrency:   "EUR",
	}
	payload, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/convert", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(ConvertHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("unexpected status code: got %v, want %v", rr.Code, http.StatusOK)
	}

	var response models.ConversionResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	expectedAmount := 91.59
	expectedCurrency := "EUR"
	epsilon := 0.01
	if diff := response.Amount - expectedAmount; diff > epsilon || diff < -epsilon {
		t.Errorf("unexpected amount: got %.2f, want %.2f", response.Amount, expectedAmount)
	}
	if response.Currency != expectedCurrency {
		t.Errorf("unexpected currency: got %s, want %s", response.Currency, expectedCurrency)
	}
}
