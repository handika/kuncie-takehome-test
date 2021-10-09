package transaction

import (
	"context"

	"github.com/handika/kuncie-takehome-test/models"
)

// Repository represent the transaction's repository contract
type Repository interface {
	GetByID(ctx context.Context, id int64) (*models.Transaction, error)
	Store(ctx context.Context, a *models.Transaction) error
}
