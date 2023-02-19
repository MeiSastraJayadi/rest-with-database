package repository

import (
	"context"
	"fmt"
	"strconv"

	"github.com/MeiSastraJayadi/rest-with-datatabase/model"
)

func (db *UseDB) CreateTransaction(ctx context.Context, table string) error {
	data := ctx.Value(ContextValue{}).(*model.Transaction)
	ex1 := fmt.Sprintf("INSERT INTO %s (buyer_id, product_id, total_pcs)", table)
  buyer := strconv.FormatInt(int64(data.Buyer), 10)
  product := strconv.FormatInt(int64(data.Product), 10)
  total := strconv.FormatInt(int64(data.Total), 10)
	ex2 := fmt.Sprintf("VALUES (%s, %s, %s)", buyer, product, total)
	q := ex1 + ex2
	_, err := db.db.ExecContext(ctx, q)
	if err != nil {
		return err
	}
  return nil
}
