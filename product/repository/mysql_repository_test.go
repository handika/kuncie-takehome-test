package repository_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	productRepo "github.com/handika/kuncie-takehome-test/product/repository"
)

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"id", "sku", "name", "price", "qty", "promotion_id"}).
		AddRow(1, "Google Home", "120P90", 49.99, 10, 1)

	query := "SELECT id, sku, name, price, qty, promotion_id FROM products WHERE id =\\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := productRepo.NewMysqlProductRepository(db)

	num := int64(1)
	anProduct, err := a.GetByID(context.TODO(), num)
	assert.NoError(t, err)
	assert.NotNil(t, anProduct)
}
