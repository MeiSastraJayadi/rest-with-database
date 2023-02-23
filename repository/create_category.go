package repository

import (
	"context"
	"fmt"

	appmodel "github.com/MeiSastraJayadi/rest-with-datatabase/model"
)

type ContextValue struct{}

func (db *UseDB) CreateCategory(ctx context.Context, table string) error {
	data := ctx.Value(ContextValue{}).(*appmodel.Category)
	command := fmt.Sprintf("INSERT INTO %s (name) VALUES ('%s')", table, data.Name)
	_, err := db.db.ExecContext(ctx, command)
	if err != nil {
		return err
	}
	return nil
}
