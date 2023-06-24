package models

type ConversionRequest struct {
	FromCurrency string  `json:"fromCurrency"`
	Amount       float64 `json:"amount"`
	ToCurrency   string  `json:"toCurrency"`
}

type ConversionResponse struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}
