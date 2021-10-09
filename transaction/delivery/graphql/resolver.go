package graphql

import (
	"context"
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/handika/kuncie-takehome-test/models"
	"github.com/handika/kuncie-takehome-test/transaction"
	"github.com/mitchellh/mapstructure"
)

// TransactionEdge holds information of transaction edge.
type TransactionEdge struct {
	Node   models.Transaction
	Cursor string
}

// TransactionResult holds information of transaction result.
type TransactionResult struct {
	Edges    []TransactionEdge
	PageInfo PageInfo
}

// PageInfo holds information of page info.
type PageInfo struct {
	EndCursor   string
	HasNextPage bool
}

type Resolver interface {
	GetTransactionByID(params graphql.ResolveParams) (interface{}, error)
	StoreTransaction(params graphql.ResolveParams) (interface{}, error)
}

type resolver struct {
	transactionService transaction.Usecase
}

func (r resolver) GetTransactionByID(params graphql.ResolveParams) (interface{}, error) {
	var (
		id int
		ok bool
	)

	ctx := context.Background()
	if id, ok = params.Args["id"].(int); !ok || id == 0 {
		return nil, fmt.Errorf("id is not integer or zero")
	}

	result, err := r.transactionService.GetByID(ctx, int64(id))
	if err != nil {
		return nil, err
	}
	return *result, nil
}

func (r resolver) StoreTransaction(params graphql.ResolveParams) (interface{}, error) {
	var transaction models.TransactionInput

	mapstructure.Decode(params.Args["input"], &transaction)

	ctx := context.Background()

	storedTransaction := &models.Transaction{
		UserId:  transaction.UserId,
		Details: mapItemsFromInput(transaction.Items),
	}

	if err := r.transactionService.Store(ctx, storedTransaction); err != nil {
		return nil, err
	}

	return *storedTransaction, nil
}

func mapItemsFromInput(transactionDetailInput []models.TransactionDetailInput) []models.TransactionDetail {
	var items []models.TransactionDetail
	for _, itemInput := range transactionDetailInput {
		items = append(items, models.TransactionDetail{
			ProductId: int(itemInput.ProductID),
			Qty:       int(itemInput.Qty),
		})
	}
	return items
}

func NewResolver(transactionService transaction.Usecase) Resolver {
	return &resolver{
		transactionService: transactionService,
	}
}
