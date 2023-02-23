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
	result, resErr := db.FetchByID(ct, table)
	if !result.Next() {
		return errors.New(NotFoundError)
	}
	defer result.Close()
	if resErr != nil {
		return errors.New(NotFoundError)
	}
	nameId := ""
	switch table {
	case "owner":
		nameId = "owner_id"
	case "product":
		nameId = "product_id"
	case "consument":
		nameId = "consument_id"
  case "transaction_table" : 
    nameId = "transaction_id"
	default:
		nameId = "id"
	}
	ex := fmt.Sprintf("DELETE FROM %s WHERE %s = %s", table, nameId, id)
	_, err := db.db.ExecContext(ctx, ex)
	if err != nil {
		return errors.New(FailedDelete)
	}
	return nil
}
