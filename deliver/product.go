package deliver

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/MeiSastraJayadi/rest-with-datatabase/usecase"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

type ProductDeliver struct {
	product *usecase.ProductUsecase
	logger  hclog.Logger
}

func NewProductDeliver(db *sql.DB, logger hclog.Logger) *ProductDeliver {
	uc := usecase.NewProductUsecase(db, logger)
	return &ProductDeliver{
		product: uc,
		logger:  logger,
	}
}

func (pd *ProductDeliver) Create(w http.ResponseWriter, r *http.Request) {
	pd.logger.Info("/product POST")
	err := pd.product.Create(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		pd.logger.Error("Failed to create product", "error", err.Error())
		return
	}
}

func (pd *ProductDeliver) FetchAll(w http.ResponseWriter, r *http.Request) {
	pd.logger.Info("/product GET")
	data, err := pd.product.GetAll(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		pd.logger.Error("Failed to retrieve data", "error", err.Error())
		return
	}
	w.Header().Add("Content-Type", "application/json")
	err = usecase.ToJSON(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		pd.logger.Error("Failed to retrieve data", "error", err.Error())
		return
	}
}

func (pd *ProductDeliver) Delete(w http.ResponseWriter, r *http.Request) {
	vr := mux.Vars(r)
	id := vr["id"]
	path := fmt.Sprintf("/product/%s DELETE", id)
	pd.logger.Info(path)
	err := pd.product.Delete(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		pd.logger.Error("Failed to delete data", "error", err.Error())
		return
	}
}
