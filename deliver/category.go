package deliver

import (
	"database/sql"
	"net/http"

	"github.com/MeiSastraJayadi/rest-with-datatabase/usecase"
	"github.com/hashicorp/go-hclog"
)

type CategoryDeliver struct {
	usecase *usecase.CategoryUsecase
	logger  hclog.Logger
}

func NewCategoryDeliver(db *sql.DB, logger hclog.Logger) *CategoryDeliver {
	cu := usecase.NewCategoryUsecase(db, logger)
	return &CategoryDeliver{
		usecase: cu,
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
