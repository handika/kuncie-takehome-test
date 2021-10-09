package models

// Promotion payless rule represent the promotion payless rule model
type ProductPaylessRule struct {
	PromotionId    int `json:"promotion_id"`
	RequirementQty int `json:"requirement_qty"`
	PromoQty       int `json:"promo_qty"`
}
