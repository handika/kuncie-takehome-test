package repository

import (
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"

	"github.com/handika/kuncie-takehome-test/models"
	"github.com/handika/kuncie-takehome-test/pdr"
)

const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00" // reduce precision from RFC3339Nano as date format
)

type mysqlPdrRepository struct {
	Conn *sql.DB
}

// NewMysqlPdrRepository will create an object that represent the pdr.Repository interface
func NewMysqlPdrRepository(Conn *sql.DB) pdr.Repository {
	return &mysqlPdrRepository{Conn}
}

func (m *mysqlPdrRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.PromoDiscountRule, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	result := make([]*models.PromoDiscountRule, 0)
	for rows.Next() {
		t := new(models.PromoDiscountRule)
		err = rows.Scan(
			&t.PromotionId,
			&t.RequirementMinQty,
			&t.PercentageDiscount,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlPdrRepository) GetByID(ctx context.Context, id int64) (res *models.PromoDiscountRule, err error) {
	query := `SELECT promotion_id, requirement_min_qty, percentage_discount
  						FROM promo_discount_rules WHERE transaction_id = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return nil, models.ErrNotFound
	}

	return
}
