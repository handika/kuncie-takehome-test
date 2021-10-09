package ppr

import (
	"context"

	"github.com/handika/kuncie-takehome-test/models"
)

// Repository represent the promo discount rule's repository contract
type Repository interface {
	GetByID(ctx context.Context, id int64) (*models.PromotionPaylessRule, error)
}
