package repository

import (
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"

	"github.com/handika/kuncie-takehome-test/models"
	"github.com/handika/kuncie-takehome-test/ppr"
)

type mysqlPprRepo struct {
	DB *sql.DB
}

// NewMysqlPprRepository will create an implementation of ppr.Repository
func NewMysqlPprRepository(db *sql.DB) ppr.Repository {
	return &mysqlPprRepo{
		DB: db,
	}
}

func (m *mysqlPprRepo) getOne(ctx context.Context, query string, args ...interface{}) (*models.PromotionPaylessRule, error) {

	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	a := &models.PromotionPaylessRule{}

	err = row.Scan(
		&a.PromotionId,
		&a.RequirementQty,
		&a.PromoQty,
	)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return a, nil
}

func (m *mysqlPprRepo) GetByID(ctx context.Context, id int64) (*models.PromotionPaylessRule, error) {
	query := `SELECT promotion_id, requirement_qty, promo_qty FROM promo_payless_rules WHERE promotion_id=?`
	return m.getOne(ctx, query, id)
}
