package models

import "time"

// Transaction represent the transaction model
type Transaction struct {
	ID         int64               `json:"id"`
	UserId     int                 `json:"user_id"`
	Date       time.Time           `json:"date"`
	GrandTotal float64             `json:"grand_total"`
	Details    []TransactionDetail `json:"transaction_detail"`
}

// Transaction detail represent the transaction detail model
type TransactionDetail struct {
	ID            int64   `json:"id"`
	TransactionId int     `json:"transaction_id"`
	ProductId     int     `json:"product_id"`
	Price         float64 `json:"price"`
	Qty           int     `json:"qty"`
	SubTotal      float64 `json:"sub_total"`
	Discount      float64 `json:"discount"`
}

type TransactionInput struct {
	UserId int                      `json:"user_id" mapstructure:"user_id"`
	Items  []TransactionDetailInput `json:"items" mapstructure:"items"`
}

type TransactionDetailInput struct {
	ProductID int `json:"product_id" mapstructure:"product_id"`
	Qty       int `json:"qty" mapstructure:"qty"`
}
