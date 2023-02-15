package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/MeiSastraJayadi/rest-with-datatabase/model"
)

func (db *UseDB) Update(ctx context.Context, table string) error {
	data := ctx.Value(ContextValue{}).(*model.Category)
	result, resErr := db.FetchByID(ctx, "category")
	if !result.Next() {
		return errors.New(NotFoundError)
	}
	defer result.Close()
	if resErr != nil {
		return resErr
	}
	id := ctx.Value("id")
	q := fmt.Sprintf("UPDATE %s SET name = '%s' WHERE id = %s", table, data.Name, id)
	_, err := db.db.ExecContext(ctx, q)
	if err != nil {
		return err
	}
	return nil
}
