package repository

import (
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"

	"github.com/handika/kuncie-takehome-test/models"
	"github.com/handika/kuncie-takehome-test/pfir"
)

type mysqlPfirRepo struct {
	DB *sql.DB
}

// NewMysqlPfirRepository will create an implementation of pfir.Repository
func NewMysqlPfirRepository(db *sql.DB) pfir.Repository {
	return &mysqlPfirRepo{
		DB: db,
	}
}

func (m *mysqlPfirRepo) getOne(ctx context.Context, query string, args ...interface{}) (*models.PromoFreeItemRule, error) {

	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	a := &models.PromoFreeItemRule{}

	err = row.Scan(
		&a.PromotionId,
		&a.FreeProductId,
	)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return a, nil
}

func (m *mysqlPfirRepo) GetByID(ctx context.Context, id int64) (*models.PromoFreeItemRule, error) {
	query := `SELECT promotion_id, free_product_id FROM promo_free_item_rules WHERE promotion_id=?`
	return m.getOne(ctx, query, id)
}
