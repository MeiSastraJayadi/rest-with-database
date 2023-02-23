package deliver

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/MeiSastraJayadi/rest-with-datatabase/usecase"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

type OwnerDeliver struct {
	owner  *usecase.OwnerUsecase
	logger hclog.Logger
}

func NewOwnerDeliver(db *sql.DB, logger hclog.Logger) *OwnerDeliver {
	uc := usecase.NewOwnerUsecase(logger, db)
	return &OwnerDeliver{
		owner:  uc,
		logger: logger,
	}
}

func (del *OwnerDeliver) Fetch(w http.ResponseWriter, r *http.Request) {
	del.logger.Info("/owner GET")
	data, err := del.owner.Fetch(r)
	if err != nil {
		del.logger.Error("Error when fetching data from owner", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = usecase.ToJSON(w, data)
	if err != nil {
		del.logger.Error("Error when encode data to json", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (del *OwnerDeliver) Delete(w http.ResponseWriter, r *http.Request) {
	vr := mux.Vars(r)
	id := vr["id"]
	path := fmt.Sprintf("/owner/%s DELETE", id)
	del.logger.Info(path, "id", id)
	err := del.owner.Delete(r)
	if err != nil {
		http.Error(w, "Failed to delete data", http.StatusBadRequest)
		del.logger.Error("Failed to delete data", "error", err)
		return
	}
}

func (del *OwnerDeliver) Update(w http.ResponseWriter, r *http.Request) {
	vr := mux.Vars(r)
	id := vr["id"]
	path := fmt.Sprintf("/owner/%s PUT", id)
	del.logger.Info(path, "id", id)
	err := del.owner.Update(r)
	if err != nil {
		http.Error(w, "Failed to update data", http.StatusBadRequest)
		del.logger.Error("Failed to update data", "error", err)
		return
	}
}

func (del *OwnerDeliver) Create(w http.ResponseWriter, r *http.Request) {
	del.logger.Info("/owner POST")
	err := del.owner.Create(r)
	if err != nil {
		del.logger.Error("Error to insert item", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
