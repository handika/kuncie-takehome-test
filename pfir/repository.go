package pfir

import (
	"context"

	"github.com/handika/kuncie-takehome-test/models"
)

// Repository represent the promo free item rule's repository contract
type Repository interface {
	GetByID(ctx context.Context, id int64) (*models.PromoFreeItemRule, error)
}
