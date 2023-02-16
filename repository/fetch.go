package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

var CancelError = "Cancel to do database operation"

func (db *UseDB) Fetch(ctx context.Context, table string) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT * FROM %s", table)
	row, err := db.db.QueryContext(ctx, query)
	select {
	case <-ctx.Done():
		err = errors.New(CancelError)
		return nil, err
	default:
		if err != nil {
			return nil, err
		}
	}

	return row, nil
}

func (db *UseDB) FetchByID(ctx context.Context, table string) (*sql.Rows, error) {
	id := ctx.Value("id")
	nameId := ""
	switch table {
	case "owner":
		nameId = "owner_id"
	default:
		nameId = "id"
	}
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = %s", table, nameId, id)
	row, err := db.db.QueryContext(ctx, query)
	select {
	case <-ctx.Done():
		err = errors.New(CancelError)
		return nil, err
	default:
		if err != nil {
			return nil, err
		}
	}

	return row, nil
}
