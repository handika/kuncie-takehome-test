package product

import (
	"context"

	"github.com/handika/kuncie-takehome-test/models"
)

// Repository represent the products's repository contract
type Repository interface {
	GetByID(ctx context.Context, id int64) (*models.Product, error)
}
