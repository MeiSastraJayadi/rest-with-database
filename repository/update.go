package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/MeiSastraJayadi/rest-with-datatabase/model"
)

func (db *UseDB) updateCategory(ctx context.Context) error {
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
	q := fmt.Sprintf("UPDATE category SET name = '%s' WHERE id = %s", data.Name, id)
	_, err := db.db.ExecContext(ctx, q)
	if err != nil {
		return err
	}
	return nil
}

func (db *UseDB) updateOwner(ctx context.Context) error {
	data := ctx.Value(ContextValue{}).(*model.Owner)
	result, resErr := db.FetchByID(ctx, "owner")
	if !result.Next() {
		return errors.New(NotFoundError)
	}
	defer result.Close()
	if resErr != nil {
		return resErr
	}
	id := ctx.Value("id")
	q := fmt.Sprintf("UPDATE owner SET ")
	newMap := make(map[string]string)
	if data.Name != "" {
		newMap["owner_name"] = data.Name
	}
	if data.Address != "" {
		newMap["owner_address"] = data.Address
	}
	if data.Phone != "" {
		newMap["phone_number"] = data.Phone
	}

	i := 0
	for key, value := range newMap {
		if len(newMap)-1 > i {
			q += fmt.Sprintf("%s = '%s', ", key, value)
			i++
		} else {
			q += fmt.Sprintf("%s = '%s' ", key, value)
		}
	}
	q += fmt.Sprintf("WHERE owner_id = %s", id)
	_, err := db.db.ExecContext(ctx, q)
	if err != nil {
		return err
	}
	return nil
}

func (db *UseDB) Update(ctx context.Context, table string) error {
	switch table {
	case "owner":
		return db.updateOwner(ctx)
	default:
		return db.updateCategory(ctx)
	}
}
