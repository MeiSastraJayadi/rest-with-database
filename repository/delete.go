package repository

import (
	"context"
	"errors"
	"fmt"
)

var NotFoundError = "Item not found"
var FailedDelete = "Failed to delete item in table"

func (db *UseDB) Delete(ctx context.Context, id string, table string) error {
	ct := context.WithValue(ctx, "id", id)
	result, resErr := db.FetchByID(ct, "category")
	if !result.Next() {
		return errors.New(NotFoundError)
	}
	defer result.Close()
	if resErr != nil {
		return errors.New(NotFoundError)
	}
	ex := fmt.Sprintf("DELETE FROM %s WHERE id = %s", table, id)
	_, err := db.db.ExecContext(ctx, ex)
	if err != nil {
		return errors.New(FailedDelete)
	}
	return nil
}
