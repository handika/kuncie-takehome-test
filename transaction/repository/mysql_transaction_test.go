package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/handika/kuncie-takehome-test/models"
	transactionRepo "github.com/handika/kuncie-takehome-test/transaction/repository"
)

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"id", "user_id", "date", "grand_total"}).
		AddRow(1, 1, time.Now(), 100)

	query := "SELECT id, user_id, date, grand_total FROM transactions WHERE id = \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := transactionRepo.NewMysqlTransactionRepository(db)

	num := int64(5)
	anTransaction, err := a.GetByID(context.TODO(), num)
	assert.NoError(t, err)
	assert.NotNil(t, anTransaction)
}

func TestUpdate(t *testing.T) {
	now := time.Now()
	ar := &models.Transaction{
		ID:         1,
		UserId:     1,
		Date:       now,
		GrandTotal: 1000,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "UPDATE transactions set title=\\?, content=\\?, author_id=\\?, updated_at=\\? WHERE ID = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(ar.UserId, ar.Date, ar.GrandTotal, ar.ID).WillReturnResult(sqlmock.NewResult(12, 1))

	a := transactionRepo.NewMysqlTransactionRepository(db)

	err = a.Update(context.TODO(), ar)
	assert.NoError(t, err)
}

// func TestStore(t *testing.T) {
// 	now := time.Now()
// 	ar := &models.Transaction{
// 		Title:     "Judul",
// 		Content:   "Content",
// 		CreatedAt: now,
// 		UpdatedAt: now,
// 		Author: models.Author{
// 			ID:   1,
// 			Name: "Iman Tumorang",
// 		},
// 	}
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}

// 	query := "INSERT  transaction SET title=\\? , content=\\? , author_id=\\?, updated_at=\\? , created_at=\\?"
// 	prep := mock.ExpectPrepare(query)
// 	prep.ExpectExec().WithArgs(ar.Title, ar.Content, ar.Author.ID, ar.CreatedAt, ar.UpdatedAt).WillReturnResult(sqlmock.NewResult(12, 1))

// 	a := transactionRepo.NewMysqlTransactionRepository(db)

// 	err = a.Store(context.TODO(), ar)
// 	assert.NoError(t, err)
// 	assert.Equal(t, int64(12), ar.ID)
// }
