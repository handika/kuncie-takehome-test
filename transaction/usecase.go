package transaction

import (
	"context"

	"github.com/handika/kuncie-takehome-test/models"
)

// Usecase represent the transaction's usecases
type Usecase interface {
	GetByID(ctx context.Context, id int64) (*models.Transaction, error)
	Store(context.Context, *models.Transaction) error
	Update(ctx context.Context, ar *models.Transaction) error
}
