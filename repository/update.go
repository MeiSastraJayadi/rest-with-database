package repository

import (
	"context"
	"errors"
	"fmt"
	"strconv"

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

func (db *UseDB) updateConsument(ctx context.Context) error {
	data := ctx.Value(ContextValue{}).(*model.Consument)
	result, resErr := db.FetchByID(ctx, "consument")
	if !result.Next() {
		return errors.New(NotFoundError)
	}
	defer result.Close()
	if resErr != nil {
		return resErr
	}
	q := "UPDATE consument SET "
	id := ctx.Value("id")
	newMap := make(map[string]string)
	if data.Name != "" {
		newMap["consument_name"] = data.Name
	}
	if data.Email != "" {
		newMap["email"] = data.Email
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
	q += fmt.Sprintf("WHERE consument_id = %s", id)
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

func (db *UseDB) updateProduct(ctx context.Context) error {
	data := ctx.Value(ContextValue{}).(*model.Product)
	result, resErr := db.FetchByID(ctx, "product")
	if !result.Next() {
		return errors.New(NotFoundError)
	}
	defer result.Close()
	if resErr != nil {
		return resErr
	}
	id := ctx.Value("id")
	q := fmt.Sprintf("UPDATE product SET ")
	newMap := make(map[string]interface{})
	if data.Name != "" {
		newMap["product_name"] = data.Name
	}
	if data.Category != nil {
		newMap["category"] = *data.Category
	}
	if data.Owner != 0 {
		newMap["product_owner"] = data.Owner
	}
	if data.Price != 0 {
		newMap["product_price"] = data.Price
	}

	if data.Stock != nil {
		newMap["stock"] = *data.Stock
	}

	i := 0
	for key, value := range newMap {
		if len(newMap)-1 > i {
			if key == "product_name" {
				q += fmt.Sprintf("%s = '%s', ", key, value.(string))
				i++
			} else {
				q += fmt.Sprintf("%s = '%s', ", key, strconv.FormatInt(int64(value.(int)), 10))
				i++
			}
		} else {
			if key == "product_name" {
				q += fmt.Sprintf("%s = '%s' ", key, value.(string))
			} else {
				q += fmt.Sprintf("%s = '%s' ", key, strconv.FormatInt(int64(value.(int)), 10))
			}
		}
	}
	q += fmt.Sprintf("WHERE product_id = %s", id)
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
	case "product":
		return db.updateProduct(ctx)
	case "consument":
		return db.updateConsument(ctx)
	default:
		return db.updateCategory(ctx)
	}
}
