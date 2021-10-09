package repository_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	pprRepo "github.com/handika/kuncie-takehome-test/ppr/repository"
)

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"promotion_id", "requirement_qty", "promo_qty"}).
		AddRow(1, 1, 1)

	query := "SELECT promotion_id, requirement_qty, promo_qty FROM promo_payless_rules WHERE transaction_id = \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := pprRepo.NewMysqlPprRepository(db)

	num := int64(5)
	anPpr, err := a.GetByID(context.TODO(), num)
	assert.NoError(t, err)
	assert.NotNil(t, anPpr)
}
