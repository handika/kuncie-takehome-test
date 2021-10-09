package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/handika/kuncie-takehome-test/models"
	"github.com/handika/kuncie-takehome-test/transaction/mocks"
	ucase "github.com/handika/kuncie-takehome-test/transaction/usecase"
)

func TestGetByID(t *testing.T) {
	mockTransactionRepo := new(mocks.Repository)
	mockTransaction := models.Transaction{
		ID: 1,
	}

	t.Run("success", func(t *testing.T) {
		mockTransactionRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(&mockTransaction, nil).Once()
		u := ucase.NewTransactionUsecase(mockTransactionRepo, time.Second*2)

		a, err := u.GetByID(context.TODO(), mockTransaction.ID)

		assert.NoError(t, err)
		assert.NotNil(t, a)

		mockTransactionRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockTransactionRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(nil, errors.New("Unexpected")).Once()

		u := ucase.NewTransactionUsecase(mockTransactionRepo, time.Second*2)

		a, err := u.GetByID(context.TODO(), mockTransaction.ID)

		assert.Error(t, err)
		assert.Nil(t, a)

		mockTransactionRepo.AssertExpectations(t)
	})

}

func TestUpdate(t *testing.T) {
	mockTransactionRepo := new(mocks.Repository)
	mockTransaction := models.Transaction{
		UserId:     1,
		Date:       time.Now(),
		GrandTotal: 1000,
		ID:         23,
	}

	t.Run("success", func(t *testing.T) {
		mockTransactionRepo.On("Update", mock.Anything, &mockTransaction).Once().Return(nil)

		u := ucase.NewTransactionUsecase(mockTransactionRepo, time.Second*2)

		err := u.Update(context.TODO(), &mockTransaction)
		assert.NoError(t, err)
		mockTransactionRepo.AssertExpectations(t)
	})
}

// func TestStore(t *testing.T) {
// 	mockTransactionRepo := new(mocks.Repository)
// 	mockTransaction := models.Transaction{
// 		ID: 1,
// 	}

// 	t.Run("success", func(t *testing.T) {
// 		tempMockTransaction := mockTransaction
// 		tempMockTransaction.ID = 1
// 		mockTransactionRepo.On("GetByTitle", mock.Anything, mock.AnythingOfType("string")).Return(nil, models.ErrNotFound).Once()
// 		mockTransactionRepo.On("Store", mock.Anything, mock.AnythingOfType("*models.Transaction")).Return(nil).Once()

// 		u := ucase.NewTransactionUsecase(mockTransactionRepo, time.Second*2)

// 		err := u.Store(context.TODO(), &tempMockTransaction)

// 		assert.NoError(t, err)
// 		assert.Equal(t, mockTransaction.ID, tempMockTransaction.ID)
// 		mockTransactionRepo.AssertExpectations(t)
// 	})
// 	t.Run("existing-title", func(t *testing.T) {
// 		existingTransaction := mockTransaction
// 		mockTransactionRepo.On("GetByTitle", mock.Anything, mock.AnythingOfType("string")).Return(&existingTransaction, nil).Once()

// 		u := ucase.NewTransactionUsecase(mockTransactionRepo, time.Second*2)

// 		err := u.Store(context.TODO(), &mockTransaction)

// 		assert.Error(t, err)
// 		mockTransactionRepo.AssertExpectations(t)
// 	})

// }
