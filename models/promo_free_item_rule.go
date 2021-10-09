package models

// Promotion free item rule represent the promotion free item rule model
type PromoFreeItemRule struct {
	PromotionId   int `json:"promotion_id"`
	FreeProductId int `json:"free_product_id"`
}
