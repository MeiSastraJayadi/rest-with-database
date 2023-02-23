package repository

import (
	"context"
	"fmt"

	"github.com/MeiSastraJayadi/rest-with-datatabase/model"
)

func (db *UseDB) CreateConsument(ctx context.Context, table string) error {
	data := ctx.Value(ContextValue{}).(*model.Consument)
	ex1 := fmt.Sprintf("INSERT INTO %s (consument_name, email, phone_number)", table)
	ex2 := fmt.Sprintf("VALUES ('%s', '%s', '%s')", data.Name, data.Email, data.Phone)
	q := ex1 + ex2
	_, err := db.db.ExecContext(ctx, q)
	if err != nil {
		return err
	}
	return nil
}
