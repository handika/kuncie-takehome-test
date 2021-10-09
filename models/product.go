package models

// Product represent the product model
type Product struct {
	ID          int64   `json:"id"`
	Sku         string  `json:"sku"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Qty         int     `json:"qty"`
	PromotionId int     `json:"promotion_id"`
}
