package repository_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/handika/kuncie-takehome-test/models"
	productRepo "github.com/handika/kuncie-takehome-test/product/repository"
)

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"id", "sku", "name", "price", "qty", "promotion_id"}).
		AddRow(1, "Google Home", "120P90", 49.99, 10, 1)

	query := "SELECT id, sku, name, price, qty, promotion_id FROM products WHERE id = \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := productRepo.NewMysqlProductRepository(db)

	num := int64(5)
	anProduct, err := a.GetByID(context.TODO(), num)
	assert.NoError(t, err)
	assert.NotNil(t, anProduct)
}

func TestUpdate(t *testing.T) {
	ar := &models.Product{
		ID:          1,
		Sku:         "120P90",
		Name:        "Google Home",
		Price:       49.99,
		Qty:         39,
		PromotionId: 2,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "UPDATE products SET sku=\\?, name=\\?, price=\\?, qty=\\? WHERE id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(ar.Sku, ar.Name, ar.Price, ar.Qty, ar.ID).WillReturnResult(sqlmock.NewResult(12, 1))

	a := productRepo.NewMysqlProductRepository(db)

	err = a.Update(context.TODO(), ar)
	assert.NoError(t, err)
}
