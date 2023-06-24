package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestConversionRequest(t *testing.T) {
	request := ConversionRequest{
		FromCurrency: "USD",
		Amount:       100.50,
		ToCurrency:   "EUR",
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		t.Fatalf("failed to marshal ConversionRequest: %v", err)
	}

	var decodedRequest ConversionRequest
	err = json.Unmarshal(jsonData, &decodedRequest)
	if err != nil {
		t.Fatalf("failed to unmarshal ConversionRequest: %v", err)
	}

	if !reflect.DeepEqual(request, decodedRequest) {
		t.Errorf("ConversionRequest mismatch, got %+v, want %+v", decodedRequest, request)
	}
}

func TestConversionResponse(t *testing.T) {
	response := ConversionResponse{
		Amount:   91.59,
		Currency: "EUR",
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		t.Fatalf("failed to marshal ConversionResponse: %v", err)
	}

	var decodedResponse ConversionResponse
	err = json.Unmarshal(jsonData, &decodedResponse)
	if err != nil {
		t.Fatalf("failed to unmarshal ConversionResponse: %v", err)
	}

	if !reflect.DeepEqual(response, decodedResponse) {
		t.Errorf("ConversionResponse mismatch, got %+v, want %+v", decodedResponse, response)
	}
}
