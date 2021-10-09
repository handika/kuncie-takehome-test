package repository

import (
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/handika/kuncie-takehome-test/models"
	"github.com/handika/kuncie-takehome-test/transaction"
)

const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00" // reduce precision from RFC3339Nano as date format
)

type mysqlTransactionRepository struct {
	Conn *sql.DB
}

// NewMysqlTransactionRepository will create an object that represent the transaction.Repository interface
func NewMysqlTransactionRepository(Conn *sql.DB) transaction.Repository {
	return &mysqlTransactionRepository{Conn}
}

func (m *mysqlTransactionRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Transaction, error) {
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

	result := make([]*models.Transaction, 0)
	for rows.Next() {
		t := new(models.Transaction)
		err = rows.Scan(
			&t.ID,
			&t.UserId,
			&t.Date,
			&t.GrandTotal,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlTransactionRepository) GetByID(ctx context.Context, id int64) (res *models.Transaction, err error) {
	query := `SELECT id, user_id, date, grand_total FROM transactions WHERE id = ?`

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

func (m *mysqlTransactionRepository) Store(ctx context.Context, a *models.Transaction) error {
	// insert into transaction
	query := `INSERT transactions SET user_id=?, grand_total=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, a.UserId, 0)
	if err != nil {
		return err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	a.ID = lastID

	var grandTotal float64 = 0
	var totalDiscount float64 = 0

	freeItems := map[int]int{}
	mainProductId := 0
	freeProductId := 0
	pfir := false

	for _, detail := range a.Details {
		// get product by id
		product := models.Product{}
		row := m.Conn.QueryRowContext(ctx, "SELECT id, price, qty, promotion_id FROM products WHERE id = ?", detail.ProductId)
		err := row.Scan(&product.ID, &product.Price, &product.Qty, &product.PromotionId)
		if err != nil {
			return err
		}

		// count sub total item
		subTotal := product.Price * float64(detail.Qty)

		// check product promotion
		freeItems[int(product.ID)] = detail.Qty
		var discount float64 = 0
		if product.PromotionId == 1 {
			// check promo free item rule
			pfir = true
			pfir := models.ProductFreeItemRule{}
			row := m.Conn.QueryRowContext(ctx, "SELECT * FROM promo_free_item_rules pfir where pfir.promotion_id = ?", product.PromotionId)
			err := row.Scan(&pfir.PromotionId, &pfir.FreeProductId)
			if err != nil {
				return err
			}

			mainProductId = int(product.ID)
			freeProductId = pfir.FreeProductId
		} else if product.PromotionId == 2 {
			// check promo payless rule
			ppr := models.ProductPaylessRule{}
			row := m.Conn.QueryRowContext(ctx, "SELECT * FROM promo_payless_rules ppr where ppr.promotion_id = ?", product.PromotionId)
			err := row.Scan(&ppr.PromotionId, &ppr.RequirementQty, &ppr.PromoQty)
			if err != nil {
				return err
			}

			if detail.Qty >= ppr.RequirementQty {
				divQty := detail.Qty / ppr.RequirementQty
				promoPrice := float64(divQty) * float64(ppr.PromoQty) * product.Price
				modQty := detail.Qty % ppr.RequirementQty
				regularPrice := float64(modQty) * product.Price
				subTotalPrice := float64(detail.Qty) * product.Price
				discount = subTotalPrice - (promoPrice + regularPrice)
			}
		} else if product.PromotionId == 3 {
			// check promo discount rule
			pdr := models.ProductDiscountRule{}
			row := m.Conn.QueryRowContext(ctx, "SELECT * FROM promo_discount_rules pdr where pdr.promotion_id = ?", product.PromotionId)
			err := row.Scan(&pdr.PromotionId, &pdr.RequirementMinQty, &pdr.PercentageDiscount)
			if err != nil {
				return err
			}

			if detail.Qty >= pdr.RequirementMinQty {
				discount = (float64(detail.Qty) * float64(product.Price)) * float64(pdr.PercentageDiscount) / 100
			}
		}

		// insert into transaction details
		query = `INSERT INTO transaction_details(transaction_id, product_id, price, qty, sub_total, discount) VALUES(?,?,?,?,?,?)`
		stmt, err = m.Conn.PrepareContext(ctx, query)
		if err != nil {
			return err
		}

		res, err = stmt.ExecContext(ctx, a.ID, detail.ProductId, product.Price, detail.Qty, subTotal, discount)
		if err != nil {
			return err
		}

		// update product qty
		var qty = product.Qty - detail.Qty
		query = `UPDATE products SET qty=? WHERE id=?`

		stmt, err = m.Conn.PrepareContext(ctx, query)
		if err != nil {
			return nil
		}

		res, err = stmt.ExecContext(ctx, qty, product.ID)
		if err != nil {
			return err
		}
		affect, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if affect != 1 {
			err = fmt.Errorf("Weird  Behaviour. Total Affected: %d", affect)

			return err
		}

		grandTotal = grandTotal + subTotal
		totalDiscount = totalDiscount + discount
	}

	// check if product with promo free item exist
	if pfir {
		var limitBuy, limitGet int
		for key, item := range freeItems {
			if key == mainProductId {
				limitBuy = item
			}

			if key == freeProductId {
				limitGet = item
			}
		}

		// get free product price
		product := models.Product{}
		row := m.Conn.QueryRowContext(ctx, "SELECT id, price FROM products WHERE id = ?", freeProductId)
		err := row.Scan(&product.ID, &product.Price)
		if err != nil {
			return err
		}

		// set calculate for free item
		var discount float64
		if limitBuy >= limitGet {
			discount = float64(limitGet) * product.Price
		} else {
			discount = float64(limitBuy) * product.Price
		}

		// set discount for free item
		query = `UPDATE transaction_details SET discount=? WHERE transaction_id=? and product_id=?`

		stmt, err = m.Conn.PrepareContext(ctx, query)
		if err != nil {
			return nil
		}

		res, err = stmt.ExecContext(ctx, discount, a.ID, freeProductId)
		if err != nil {
			return err
		}
		affect, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if affect != 1 {
			err = fmt.Errorf("Weird  Behaviour. Total Affected: %d", affect)

			return err
		}

		// count total discount
		totalDiscount = totalDiscount + discount
	}

	// update transaction grand total
	grandTotal = grandTotal - totalDiscount
	query = `UPDATE transactions set grand_total=? WHERE id = ?`

	stmt, err = m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil
	}

	res, err = stmt.ExecContext(ctx, grandTotal, a.ID)
	if err != nil {
		return err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affect != 1 {
		err = fmt.Errorf("Weird  Behaviour. Total Affected: %d", affect)

		return err
	}

	return nil
}

// DecodeCursor will decode cursor from user for mysql
func DecodeCursor(encodedTime string) (time.Time, error) {
	byt, err := base64.StdEncoding.DecodeString(encodedTime)
	if err != nil {
		return time.Time{}, err
	}

	timeString := string(byt)
	t, err := time.Parse(timeFormat, timeString)

	return t, err
}

// EncodeCursor will encode cursor from mysql to user
func EncodeCursor(t time.Time) string {
	timeString := t.Format(timeFormat)

	return base64.StdEncoding.EncodeToString([]byte(timeString))
}
