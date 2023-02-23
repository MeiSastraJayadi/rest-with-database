package usecase

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/MeiSastraJayadi/rest-with-datatabase/model"
	"github.com/MeiSastraJayadi/rest-with-datatabase/repository"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

type ConsumentUsecase struct {
	db     *repository.UseDB
	logger hclog.Logger
}

func NewConsumentUsecase(db *sql.DB, logger hclog.Logger) *ConsumentUsecase {
	usedb := repository.NewUseDB(logger, db)
	return &ConsumentUsecase{
		db:     usedb,
		logger: logger,
	}
}

func (cd *ConsumentUsecase) Delete(r *http.Request) error {
	vr := mux.Vars(r)
	ctx := r.Context()
	id := vr["id"]
	err := cd.db.Delete(ctx, id, "consument")
	if err != nil {
		return err
	}
	return nil
}

func (cu *ConsumentUsecase) Create(r *http.Request) error {
	ctx := r.Context()
	data := &model.Consument{}
	err := FromJSON(r.Body, data)
	if err != nil {
		return err
	}
	ctx = context.WithValue(ctx, repository.ContextValue{}, data)
	err = cu.db.CreateConsument(ctx, "consument")
	if err != nil {
		return err
	}
	return nil
}

func (cu *ConsumentUsecase) Update(r *http.Request) error {
	ctx := r.Context()
	data := &model.Consument{}
	err := FromJSON(r.Body, data)
	if err != nil {
		return err
	}
	vr := mux.Vars(r)
	id := vr["id"]
	ctx = context.WithValue(ctx, "id", id)
	ctx = context.WithValue(ctx, repository.ContextValue{}, data)
	err = cu.db.Update(ctx, "consument")
	if err != nil {
		return err
	}
	return nil
}

func (cu *ConsumentUsecase) FetchAll(r *http.Request) (*model.Consuments, error) {
	ctx := r.Context()
	result, err := cu.db.Fetch(ctx, "consument")
	if err != nil {
		return nil, err
	}
	var consumentData model.Consuments
	for result.Next() {
		var (
			id    int
			name  string
			email string
			phone string
		)

		err = result.Scan(&id, &name, &email, &phone)
		if err != nil {
			return nil, err
		}
		consument := model.NewConsument(id, name, email, phone)
		consumentData = append(consumentData, consument)
	}
	if consumentData == nil {
		consumentData = make(model.Consuments, 0)
	}
	return &consumentData, nil
}
