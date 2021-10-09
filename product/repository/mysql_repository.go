package repository

import (
	"context"
	"database/sql"

	"github.com/handika/kuncie-takehome-test/models"
	"github.com/handika/kuncie-takehome-test/product"
	"github.com/sirupsen/logrus"
)

const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00" // reduce precision from RFC3339Nano as date format
)

type mysqlProductRepository struct {
	Conn *sql.DB
}

// NewMysqlProductRepository will create an object that represent the product.Repository interface
func NewMysqlProductRepository(Conn *sql.DB) product.Repository {
	return &mysqlProductRepository{Conn}
}

func (m *mysqlProductRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Product, error) {
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

	result := make([]*models.Product, 0)
	for rows.Next() {
		t := new(models.Product)
		err = rows.Scan(
			&t.ID,
			&t.Sku,
			&t.Name,
			&t.Price,
			&t.Qty,
			&t.PromotionId,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlProductRepository) GetByID(ctx context.Context, id int64) (res *models.Product, err error) {
	query := `SELECT id, sku, name, price, qty, promotion_id FROM products WHERE id=?`

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

// func (m *mysqlProductRepository) Update(ctx context.Context, ar *models.Product) error {
// 	query := `UPDATE products set sku=?, name=?, price=?, qty=?, discount=? WHERE ID = ?`

// 	stmt, err := m.Conn.PrepareContext(ctx, query)
// 	if err != nil {
// 		return nil
// 	}

// 	res, err := stmt.ExecContext(ctx, ar.Sku, ar.Name, ar.Price, ar.Qty, ar.Discount)
// 	if err != nil {
// 		return err
// 	}
// 	affect, err := res.RowsAffected()
// 	if err != nil {
// 		return err
// 	}
// 	if affect != 1 {
// 		err = fmt.Errorf("Weird  Behaviour. Total Affected: %d", affect)

// 		return err
// 	}

// 	return nil
// }
