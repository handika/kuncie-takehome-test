package repository

import (
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"

	"github.com/handika/kuncie-takehome-test/models"
	"github.com/handika/kuncie-takehome-test/ppr"
)

const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00" // reduce precision from RFC3339Nano as date format
)

type mysqlPprRepository struct {
	Conn *sql.DB
}

// NewMysqlPprRepository will create an object that represent the ppr.Repository interface
func NewMysqlPprRepository(Conn *sql.DB) ppr.Repository {
	return &mysqlPprRepository{Conn}
}

func (m *mysqlPprRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.PromoPaylessRule, error) {
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

	result := make([]*models.PromoPaylessRule, 0)
	for rows.Next() {
		t := new(models.PromoPaylessRule)
		err = rows.Scan(
			&t.PromotionId,
			&t.RequirementQty,
			&t.PromoQty,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlPprRepository) GetByID(ctx context.Context, id int64) (res *models.PromoPaylessRule, err error) {
	query := `SELECT promotion_id, requirement_qty, promo_qty
  						FROM promo_payless_rules WHERE transaction_id = ?`

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
