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

type TransactionUsecase struct {
	db *repository.UseDB
  logger hclog.Logger
}

func NewTransactionUsecase(db *sql.DB, logger hclog.Logger) *TransactionUsecase {
  usedb := repository.NewUseDB(logger, db)
  return &TransactionUsecase{
    db : usedb, 
    logger: logger,
  }
}

func (tu *TransactionUsecase) Create(r *http.Request) error {
  data := &model.Transaction{}
  err := FromJSON(r.Body, data)
  if err != nil {
    return err
  }
  ctx := r.Context()
  ctx = context.WithValue(ctx, repository.ContextValue{}, data)
  err = tu.db.CreateTransaction(ctx, "transaction_table")
  if err != nil {
    return err
  }
  return nil
}

func (tr *TransactionUsecase) Delete(r *http.Request) error {
  vr := mux.Vars(r)
  ctx := r.Context()
  id := vr["id"]
  err := tr.db.Delete(ctx, id, "transaction_table")
  if err != nil {
    return err
  }
  return nil
}

func (tu *TransactionUsecase) FetchAll(r *http.Request) (*model.Transactions, error){
  ctx := r.Context()
  result, err := tu.db.Fetch(ctx, "transaction_table")
  if err != nil {
    return nil, err
  }
  
  var data model.Transactions
  for result.Next() {
    var (
      id int
      buyer int
      product int
      pcs int
    )
    err = result.Scan(&id, &buyer, &product, &pcs) 
    if err != nil {
      return nil, err
    }
    transaction := model.NewTransaction(id, buyer, product, pcs) 
    data = append(data, transaction)
  }
  if data == nil {
    data = make(model.Transactions, 0)
    
  }
  return &data, nil
}


