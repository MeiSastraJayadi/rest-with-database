package deliver

import (
	"database/sql"
	"net/http"

	"github.com/MeiSastraJayadi/rest-with-datatabase/usecase"
	"github.com/hashicorp/go-hclog"
)

type TransactionDeliver struct {
  transaction *usecase.TransactionUsecase
  logger hclog.Logger
}

func NewTransactionDeliver(db *sql.DB, l hclog.Logger) *TransactionDeliver {
  uc := usecase.NewTransactionUsecase(db, l)
  return &TransactionDeliver{
    transaction: uc,
    logger: l,
  }
}

func (tr *TransactionDeliver) Create(w http.ResponseWriter, r *http.Request) {
  tr.logger.Info("/transaction POST")
  err := tr.transaction.Create(r)
  if err != nil {
    http.Error(w, "Failed to fetch data from table transaction", http.StatusInternalServerError)
    tr.logger.Error("Failed to fetch data from table transaction", "error", err.Error())
    return
  }
}

func (tr *TransactionDeliver) FetchAll(w http.ResponseWriter, r *http.Request) {
  tr.logger.Info("/transaction GET")
  result, err := tr.transaction.FetchAll(r)
  if err != nil {
    http.Error(w, "Failed to fetch data from table transaction", http.StatusInternalServerError)
    tr.logger.Error("Failed to fetch data from table transaction", "error", err.Error())
    return
  }
  w.Header().Set("Content-Type", "application/json")
  err = usecase.ToJSON(w, result)
  if err != nil {
    http.Error(w, "Failed to encode data from table transaction", http.StatusInternalServerError)
    tr.logger.Error("Failed to encode data from table transaction", "error", err.Error())
    return
  }
}

