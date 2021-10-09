package models

// Promotion discount rule represent the promotion discount rule model
type PromoDiscountRule struct {
	PromotionId        int `json:"promotion_id"`
	RequirementMinQty  int `json:"requirement_min_qty"`
	PercentageDiscount int `json:"percentage_discount"`
}
