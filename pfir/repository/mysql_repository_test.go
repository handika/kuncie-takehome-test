package repository_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	pdrRepo "github.com/handika/kuncie-takehome-test/pfir/repository"
)

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"promotion_id", "free_product_id"}).
		AddRow(1, 1)

	query := "SELECT promotion_id, free_product_id FROM promo_free_item_rules WHERE transaction_id = \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := pdrRepo.NewMysqlPfirRepository(db)

	num := int64(5)
	anPfir, err := a.GetByID(context.TODO(), num)
	assert.NoError(t, err)
	assert.NotNil(t, anPfir)
}
