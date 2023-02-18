package repository

import (
	"context"
	"fmt"
	"strconv"

	"github.com/MeiSastraJayadi/rest-with-datatabase/model"
)

func (db *UseDB) CreateProduct(ctx context.Context, table string) error {
	data := ctx.Value(ContextValue{}).(*model.Product)
	category := strconv.FormatInt(int64(*data.Category), 10)
	owner := strconv.FormatInt(int64(data.Owner), 10)
	price := strconv.FormatInt(int64(data.Price), 10)
	stock := strconv.FormatInt(int64(data.Stock), 10)
	query := fmt.Sprintf("INSERT INTO %s (product_name, category, product_owner, product_price, stock) VALUES ", table)
	query += fmt.Sprintf("('%s', %s, %s, %s, %s)", data.Name, category, owner, price, stock)
	_, err := db.db.ExecContext(ctx, query)
	if err != nil {
		return err
	}
	return nil
}
