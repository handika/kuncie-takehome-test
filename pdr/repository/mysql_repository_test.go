package repository_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/handika/kuncie-takehome-test/pdr/repository"
)

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"promotion_id", "requirement_min_qty", "percentage_discount"}).
		AddRow(2, 3, 10)

	query := "SELECT promotion_id, requirement_min_qty, percentage_discount FROM promo_discount_rules WHERE promotion_id=\\?"

	prep := mock.ExpectPrepare(query)
	userID := int64(1)
	prep.ExpectQuery().WithArgs(userID).WillReturnRows(rows)

	a := repository.NewMysqlPdrRepository(db)

	anArticle, err := a.GetByID(context.TODO(), userID)
	assert.NoError(t, err)
	assert.NotNil(t, anArticle)
}
