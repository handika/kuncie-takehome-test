package repository

import (
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"

	"github.com/handika/kuncie-takehome-test/models"
	"github.com/handika/kuncie-takehome-test/pfir"
)

const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00" // reduce precision from RFC3339Nano as date format
)

type mysqlPfirRepository struct {
	Conn *sql.DB
}

// NewMysqlPfirRepository will create an object that represent the pfir.Repository interface
func NewMysqlPfirRepository(Conn *sql.DB) pfir.Repository {
	return &mysqlPfirRepository{Conn}
}

func (m *mysqlPfirRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.PromoFreeItemRule, error) {
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

	result := make([]*models.PromoFreeItemRule, 0)
	for rows.Next() {
		t := new(models.PromoFreeItemRule)
		err = rows.Scan(
			&t.PromotionId,
			&t.FreeProductId,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlPfirRepository) GetByID(ctx context.Context, id int64) (res *models.PromoFreeItemRule, err error) {
	query := `SELECT promotion_id, free_product_id
  						FROM promo_free_item_rules WHERE transaction_id = ?`

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
