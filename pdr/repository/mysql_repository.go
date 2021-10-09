package repository

import (
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"

	"github.com/handika/kuncie-takehome-test/models"
	"github.com/handika/kuncie-takehome-test/pdr"
)

type mysqlPdrRepo struct {
	DB *sql.DB
}

// NewMysqlPdrRepository will create an implementation of pdr.Repository
func NewMysqlPdrRepository(db *sql.DB) pdr.Repository {
	return &mysqlPdrRepo{
		DB: db,
	}
}

func (m *mysqlPdrRepo) getOne(ctx context.Context, query string, args ...interface{}) (*models.PromoDiscountRule, error) {

	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	a := &models.PromoDiscountRule{}

	err = row.Scan(
		&a.PromotionId,
		&a.RequirementMinQty,
		&a.PercentageDiscount,
	)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return a, nil
}

func (m *mysqlPdrRepo) GetByID(ctx context.Context, id int64) (*models.PromoDiscountRule, error) {
	query := `SELECT promotion_id, requirement_min_qty, percentage_discount FROM promo_discount_rules WHERE promotion_id=?`
	return m.getOne(ctx, query, id)
}
