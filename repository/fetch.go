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
