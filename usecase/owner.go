package usecase

import (
	"context"
	"database/sql"
	"errors"
	"io"
	"net/http"

	"github.com/MeiSastraJayadi/rest-with-datatabase/model"
	"github.com/MeiSastraJayadi/rest-with-datatabase/repository"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

type OwnerUsecase struct {
	db     *repository.UseDB
	logger hclog.Logger
}

func NewOwnerUsecase(logger hclog.Logger, db *sql.DB) *OwnerUsecase {
	usedb := repository.NewUseDB(logger, db)
	return &OwnerUsecase{
		db:     usedb,
		logger: logger,
	}
}

func (ou *OwnerUsecase) Fetch(w io.Writer, r *http.Request) (*model.Owners, error) {
	ctx := r.Context()
	row, err := ou.db.Fetch(ctx, "owner")
	defer row.Close()
	if err != nil {
		return nil, err
	}
	var owners model.Owners
	for row.Next() {
		var (
			id      int
			name    string
			address string
			phone   string
		)
		err = row.Scan(&id, &name, &address, &phone)
		if err != nil {
			return nil, err
		}
		owner := model.NewOwner(id, name, address, phone)
		owners = append(owners, owner)
	}

	if owners == nil {
		owners = make(model.Owners, 0)
	}
	return &owners, nil
}

func (ou *OwnerUsecase) Delete(r *http.Request) error {
	vr := mux.Vars(r)
	id := vr["id"]
	ctx := r.Context()
	err := ou.db.Delete(ctx, id, "owner")
	if err != nil {
		return err
	}
	return nil
}

func (ou *OwnerUsecase) Update(r *http.Request) error {
	data := &model.Owner{}
	err := FromJSON(r.Body, data)
	if err != nil {
		return err
	}
	vr := mux.Vars(r)
	id := vr["id"]
	ctx := r.Context()
	ctx = context.WithValue(ctx, repository.ContextValue{}, data)
	ctx = context.WithValue(ctx, "id", id)
	err = ou.db.Update(ctx, "owner")
	if err != nil {
		return err
	}
	return nil
}

func (ou *OwnerUsecase) Create(r *http.Request) error {
	data := &model.Owner{}
	jsonError := FromJSON(r.Body, data)
	if jsonError != nil {
		errors.New("Failed to decode data to json")
	}
	ctx := r.Context()
	ctx = context.WithValue(ctx, repository.ContextValue{}, data)
	err := ou.db.CreateOwner(ctx, "owner")
	if err != nil {
		return errors.New("Failed insert data to table owner")
	}
	return nil
}
