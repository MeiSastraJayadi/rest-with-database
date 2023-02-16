package repository

import (
	"context"
	"fmt"

	"github.com/MeiSastraJayadi/rest-with-datatabase/model"
)

func (db *UseDB) CreateOwner(ctx context.Context, table string) error {
	data := ctx.Value(ContextValue{}).(*model.Owner)
	ex1 := fmt.Sprintf("INSERT INTO %s (owner_name, owner_address, phone_number)", table)
	ex2 := fmt.Sprintf("VALUES ('%s', '%s', '%s')", data.Name, data.Address, data.Phone)
	q := ex1 + ex2
	_, err := db.db.ExecContext(ctx, q)
	if err != nil {
		return err
	}
	return nil
}
