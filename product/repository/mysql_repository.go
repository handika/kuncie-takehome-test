package repository

import (
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"

	"github.com/handika/kuncie-takehome-test/models"
	"github.com/handika/kuncie-takehome-test/product"
)

type mysqlProductRepo struct {
	DB *sql.DB
}

// NewMysqlProductRepository will create an implementation of product.Repository
func NewMysqlProductRepository(db *sql.DB) product.Repository {
	return &mysqlProductRepo{
		DB: db,
	}
}

func (m *mysqlProductRepo) getOne(ctx context.Context, query string, args ...interface{}) (*models.Product, error) {

	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	a := &models.Product{}

	err = row.Scan(
		&a.ID,
		&a.Sku,
		&a.Name,
		&a.Price,
		&a.Qty,
		&a.PromotionId,
	)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return a, nil
}

func (m *mysqlProductRepo) GetByID(ctx context.Context, id int64) (*models.Product, error) {
	query := `SELECT id, sku, name, price, qty, promotion_id FROM products WHERE id=?`
	return m.getOne(ctx, query, id)
}
