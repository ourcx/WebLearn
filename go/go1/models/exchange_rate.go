package models

import "time"

type ExchangeRate struct {
	ID           uint      `gorm:"primary_key" json:"_id"`
	FromCurrency string    `json:"from_currency"`
	ToCurrency   string    `json:"to_currency"`
	Rate         float64   `json:"rate"`
	Data         time.Time `json:"data"`
}
