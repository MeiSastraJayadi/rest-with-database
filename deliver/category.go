package deliver

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/MeiSastraJayadi/rest-with-datatabase/usecase"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

type CategoryDeliver struct {
	usecase *usecase.CategoryUsecase
	reponse *Reponse
	logger  hclog.Logger
}

func NewCategoryDeliver(db *sql.DB, logger hclog.Logger) *CategoryDeliver {
	cu := usecase.NewCategoryUsecase(db, logger)
	return &CategoryDeliver{
		usecase: cu,
		reponse: &Reponse{},
		logger:  logger,
	}
}

func (cd *CategoryDeliver) GetAll(w http.ResponseWriter, r *http.Request) {
	cd.logger.Info("/category GET")
	ctg, err := cd.usecase.Fetch(w, r)
	if err != nil {
		cd.logger.Error("Error when fetching data", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = usecase.ToJSON(w, ctg)
	if err != nil {
		http.Error(w, "Failed to encode data", http.StatusInternalServerError)
		cd.logger.Error("Failed to encode data", "error", err)
	}
}

func (cd *CategoryDeliver) Create(w http.ResponseWriter, r *http.Request) {
	cd.logger.Info("/category POST")
	err := cd.usecase.Create(r.Body, r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		cd.logger.Error("Failed to create new category", "error", err)
		return
	}
}

func (cd *CategoryDeliver) Delete(w http.ResponseWriter, r *http.Request) {
	vr := mux.Vars(r)
	id := vr["id"]
	path := fmt.Sprintf("/category/%s DELETE", id)
	cd.logger.Info(path, "id", id)
	err := cd.usecase.Delete(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		cd.logger.Error("Failed to delete item on table category", "error", err)
		return
	}
}

func (cd *CategoryDeliver) Update(w http.ResponseWriter, r *http.Request) {
	vr := mux.Vars(r)
	id := vr["id"]
	path := fmt.Sprintf("/category/%s PUT", id)
	cd.logger.Info(path, "id", id)
	ctx := r.Context()
	ctx = context.WithValue(ctx, "id", id)
	r = r.WithContext(ctx)
	err := cd.usecase.Update(r)
	if err != nil {
		http.Error(w, "Fail to update data", http.StatusInternalServerError)
		cd.logger.Error("Fail to update data", "error", err)
		return
	}
}
