package usecase

import (
	"context"
	"time"

	"github.com/handika/kuncie-takehome-test/models"
	"github.com/handika/kuncie-takehome-test/transaction"
)

type transactionUsecase struct {
	transactionRepo transaction.Repository
	contextTimeout  time.Duration
}

// NewTransactionUsecase will create new an transactionUsecase object representation of transaction.Usecase interface
func NewTransactionUsecase(a transaction.Repository, timeout time.Duration) transaction.Usecase {
	return &transactionUsecase{
		transactionRepo: a,
		contextTimeout:  timeout,
	}
}

/*
* In this function below, I'm using errgroup with the pipeline pattern
* Look how this works in this package explanation
* in godoc: https://godoc.org/golang.org/x/sync/errgroup#ex-Group--Pipeline
 */
func (a *transactionUsecase) GetByID(c context.Context, id int64) (*models.Transaction, error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err := a.transactionRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *transactionUsecase) Store(c context.Context, m *models.Transaction) error {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	err := a.transactionRepo.Store(ctx, m)
	if err != nil {
		return err
	}
	return nil
}

func (a *transactionUsecase) Update(c context.Context, ar *models.Transaction) error {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	return a.transactionRepo.Update(ctx, ar)
}
