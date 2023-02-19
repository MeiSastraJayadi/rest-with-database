package deliver

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/MeiSastraJayadi/rest-with-datatabase/usecase"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

type ConsumentDeliver struct {
	consument *usecase.ConsumentUsecase
	logger    hclog.Logger
}

func NewConsumentDeliver(db *sql.DB, logger hclog.Logger) *ConsumentDeliver {
	uc := usecase.NewConsumentUsecase(db, logger)
	return &ConsumentDeliver{
		consument: uc,
		logger:    logger,
	}
}

func (cd *ConsumentDeliver) Create(w http.ResponseWriter, r *http.Request) {
	cd.logger.Info("/consument POST")
	err := cd.consument.Create(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		cd.logger.Error("Failed to create data", "error", err.Error())
		return
	}
}

func (cd *ConsumentDeliver) Delete(w http.ResponseWriter, r *http.Request) {
	vr := mux.Vars(r)
	id := vr["id"]
	path := fmt.Sprintf("/consument/%s DELETE", id)
	cd.logger.Info(path)
	err := cd.consument.Delete(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		cd.logger.Error("Failed to delete data", "error", err.Error())
		return
	}
}

func (cd *ConsumentDeliver) Update(w http.ResponseWriter, r *http.Request) {
	vr := mux.Vars(r)
	id := vr["id"]
	path := fmt.Sprintf("/consument/%s PUT", id)
	cd.logger.Info(path)
	err := cd.consument.Update(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		cd.logger.Error("Failed to update data", "error", err.Error())
		return
	}
}

func (cd *ConsumentDeliver) Fetch(w http.ResponseWriter, r *http.Request) {
	cd.logger.Info("/consument GET")
	data, err := cd.consument.FetchAll(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		cd.logger.Error("Failed to fetch data", "error", err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = usecase.ToJSON(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		cd.logger.Error("Failed to fetch data", "error", err.Error())
		return
	}
}
